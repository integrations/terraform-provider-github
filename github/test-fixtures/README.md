# Hi fellow bots and humans :wave:

If you're about to panic about leaked private key, then please don't.
This is purposefully exposed self-signed cert & key used in
acceptance tests of the GitHub provider.

Thanks for understanding :heart:

-----

Generated via

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
