# Hi fellow bots and humans :wave:

If you're about to panic about leaked private key, then please don't.
This is purposefully exposed self-signed cert & key used in
acceptance tests of the GitHub provider.

Thanks for understanding :heart:

-----

Generated via

```
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```
