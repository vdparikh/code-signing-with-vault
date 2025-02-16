#!/bin/bash

# Base64-encode the data to be signed
data=$(echo -n "your-code-or-data" | base64)

# Sign the data using the Transit Engine
signature=$(vault write -field=signature transit/sign/code-signing-key \
    input="$data" \
    hash_algorithm="sha2-256")

# Save the signature to a file
echo "$signature" > signature.bin

echo "Signature saved to signature.bin"