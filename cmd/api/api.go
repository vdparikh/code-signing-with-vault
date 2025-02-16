package main

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
)

func main() {
	// Set Vault address and token from environment variables
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

	// Save the certificate to a file
	certFile, err := os.Create("certificate.pem")
	if err != nil {
		log.Fatalf("Failed to create certificate file: %v", err)
	}
	defer certFile.Close()

	_, err = certFile.WriteString(certificate)
	if err != nil {
		log.Fatalf("Failed to write certificate to file: %v", err)
	}

	// Save the issuing CA certificate (public key) to a file
	issuingCA := secret.Data["issuing_ca"].(string)
	caFile, err := os.Create("issuing_ca.pem")
	if err != nil {
		log.Fatalf("Failed to create issuing CA file: %v", err)
	}
	defer caFile.Close()

	_, err = caFile.WriteString(issuingCA)
	if err != nil {
		log.Fatalf("Failed to write issuing CA to file: %v", err)
	}

	// Store the private key in the Transit Engine
	encodedPrivateKey := base64.StdEncoding.EncodeToString([]byte(privateKey))
	_, err = client.Logical().Write("transit/keys/code-signing-key", map[string]interface{}{
		"type": "rsa-2048",
		"key":  encodedPrivateKey,
	})
	if err != nil {
		log.Fatalf("Failed to store private key in Transit Engine: %v", err)
	}

	log.Println("Certificate saved to certificate.pem")
	log.Println("Issuing CA (public key) saved to issuing_ca.pem")
}
