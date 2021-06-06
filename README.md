# cryptography-golang
This repository contains the implementation of various cryptography concepts in golang

## Hash
```bash
# Get help
cryptoctl hash -h

# Calculate SHA 256, SHA 512 Hash
cryptoctl hash --algo="sha2" --input="hello"

# Calculate SHA3-256, SHA3-384, SHA3-512 Hash
cryptoctl hash --algo="sha3" --input="hello"
```