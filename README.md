# Code Signing with HashiCorp Vault

This repository demonstrates how to use **HashiCorp Vault's PKI and Transit engines** to securely manage code signing certificates and perform signing operations without exposing private keys.

## **Features**
- Issue code-signing certificates using Vault's PKI engine.
- Securely store private keys in Vault's Transit engine.
- Sign and verify data using Vault's Transit engine.
- Verify signatures using OpenSSL and the public key.

---

## **Prerequisites**
1. **HashiCorp Vault**: Installed and running. [Install Vault](https://learn.hashicorp.com/tutorials/vault/getting-started-install).
2. **OpenSSL**: Installed on your system.
3. **Go**: If you plan to run the Go program. [Install Go](https://golang.org/doc/install).

---

## **Setup**

### **1. Start Vault in Development Mode**
Run Vault in development mode for testing:

```bash
vault server -dev
```

Set the environment variables:

```bash
export VAULT_ADDR='http://127.0.0.1:8200'
export VAULT_TOKEN='<root-token>'
```

### **2. Enable and Configure PKI and Transit Engines**
Run the setup script to enable and configure the PKI and Transit engines:

```bash
chmod +x ./scripts/vault-setup.sh
./scripts/vault-setup.sh
```

---

## **Usage**

### **1. Issue a Code-Signing Certificate**
Run the Go program to issue a code-signing certificate and store the private key in the Transit engine:

```bash
go run cmd/api/api.go
```

This will:
- Issue a certificate and save it to `certificate.pem`.
- Save the issuing CA certificate (public key) to `issuing_ca.pem`.

### **2. Sign Data Using the Transit Engine**
Use the `sign-data.sh` script to sign data:

```bash
chmod +x ./scripts/sign-data.sh
./scripts/sign-data.sh
```

This will:
- Sign the data using the Transit engine.
- Save the signature to `signature.bin`.

### **3. Verify the Signature**
We can either use Vault to verify of use the public key

#### Using Vault
Use the `verify-signature.sh` script to verify the signature:

```bash
chmod +x ./scripts/verify-signature.sh
./scripts/verify-signature.sh
```

#### Using Public Key
Use the `verify-with-public-key.sh` script to verify the signature:

```bash
chmod +x ./scripts/verify-with-public-key.sh
./scripts/verify-with-public-key.sh
```

This will:
- Extract the public key from the certificate.
- Verify the signature using the public key.




## **License**
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
