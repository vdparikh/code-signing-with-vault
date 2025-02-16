#!/bin/bash

# Base64-encode the data to be verified
data=$(echo -n "your-code-or-data" | base64)

# Verify the signature using the Transit Engine
vault write transit/verify/code-signing-key \
    input="$data" \
    signature="$(cat signature.bin)" \
    hash_algorithm="sha2-256"