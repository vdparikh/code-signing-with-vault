package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
)

func main() {

	vaultAddr := os.Getenv("VAULT_ADDR")
	token := os.Getenv("VAULT_TOKEN")
	// Vault client configuration
	config := api.DefaultConfig()
	config.Address = vaultAddr

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Vault client: %v", err)
	}

	client.SetToken(token)

	// Issue a code-signing certificate
	secret, err := client.Logical().Write("pki/issue/code-signing-role", map[string]interface{}{
		"common_name": "code-signer.vishalparikh.me",
		"ttl":         "720h",
	})
	if err != nil {
		log.Fatalf("Failed to issue certificate: %v", err)
	}

	// Extract certificate and private key
	certificate := secret.Data["certificate"].(string)
	privateKey := secret.Data["private_key"].(string)

	// Store the private key in the Transit Engine
	encodedPrivateKey := base64.StdEncoding.EncodeToString([]byte(privateKey))
	_, err = client.Logical().Write("transit/keys/code-signing-key", map[string]interface{}{
		"type": "rsa-2048",
		"key":  encodedPrivateKey,
	})
	if err != nil {
		log.Fatalf("Failed to store private key in Transit Engine: %v", err)
	}

	// Return the public key
	fmt.Println("Certificate:\n", certificate)
	fmt.Println("Public Key (for verification):\n", secret.Data["issuing_ca"].(string))
}
