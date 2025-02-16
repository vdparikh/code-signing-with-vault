#!/bin/bash

# Extract the public key from the certificate
openssl x509 -in certificate.pem -pubkey -noout > public-key.pem

# Save the base64-encoded data to a temporary file
data=$(echo -n "your-code-or-data" | base64)
echo -n "$data" | base64 --decode > data-file.txt

# Remove the "vault:v1:" prefix and decode the signature
signature=$(cat signature.bin | sed 's/^vault:v1://' | base64 --decode)

# Save the decoded signature to a temporary file
echo -n "$signature" > decoded-signature.bin

# Verify the signature using the public key
openssl dgst -sha256 -verify public-key.pem -signature decoded-signature.bin data-file.txt

# Clean up
rm public-key.pem data-file.txt decoded-signature.bin