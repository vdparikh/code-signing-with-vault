#!/bin/bash

# Enable PKI Secrets Engine
vault secrets enable pki

# Configure PKI Engine with Root CA
vault write pki/root/generate/internal \
    common_name="My Code Signing CA" \
    ttl=87600h

# Create Role for Code-Signing Certificates
vault write pki/roles/code-signing-role \
    allowed_domains="vishalparikh.me" \
    allow_subdomains=true \
    max_ttl=720h \
    key_type="rsa" \
    key_bits=2048 \
    ext_key_usage="code_signing"

# Enable Transit Secrets Engine
vault secrets enable transit