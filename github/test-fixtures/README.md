# Hi fellow bots and humans :wave:

If you're about to panic about leaked private keys, then please don't.
These are purposefully exposed cryptographic materials used in tests of
the GitHub provider.

Thanks for understanding :heart:

## Self-Signed Certificate (cert.pem, domain.csr, key.pem)
A self-signed cert & key used in acceptance tests of the GitHub provider.

Generated via:

```sh
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```

You can "extend" the key expiration by using following commands:

```sh
# Generate CSR from existing certificate.
openssl x509 \
   -in cert.pem \
   -signkey key.pem \
   -x509toreq -out csr.pem

# Generate self-signed certificate using existing key and CSR.
openssl x509 \
   -signkey key.pem \
   -in csr.pem \
   -req -days 3650 -out cert.pem
```

## GPG Key (gpg-pubkey.asc)
Terraform's acceptance tests GPG public key.

## SSH Public Key (id_rsa.pub)

Terraform's acceptance tests SSH public key.

## GitHub App Keys (github-app-key.pem, github-app-key.pub)

Terraform's acceptance tests GitHub App public and private key.

You can re-generate them by using the following commands:

```
# Private Key
openssl rsa -in github-app-key.pem -pubout -outform PEM

# Public Key
openssl rsa -in github-app-key.pem -out github/test-fixtures/github-app-key.pub -pubout -outform PEM
```
