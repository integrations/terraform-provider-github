# Security policy

Terraform Provider Tester is an internal maintainer tool for
[terraform-provider-github](https://github.com/integrations/terraform-provider-github).
It runs acceptance tests that create real GitHub resources and it reads real
credentials from the environment.

## Reporting a vulnerability

Do not open a public issue for a security problem. Report it privately:

- Use GitHub private vulnerability reporting on this repository (the Security
  tab, "Report a vulnerability"), or
- Ping a maintainer listed in `MAINTAINERS.md` directly.

Please include what you found, how to reproduce it, and the impact you expect.
We will confirm receipt and keep you updated while we work on a fix.

## How Terraform Provider Tester handles credentials

Terraform Provider Tester never asks you to put a token in a file or on the command line. It reads
PATs and app keys from environment variables, validates them, and redacts them
in all output and saved state. If you ever see a credential printed in full,
treat it as a bug and report it.

- State files (`.pulsar-state*`) and failure logs (`.pulsar-failures/`) are
  written only inside the provider checkout and are gitignored. They are
  redacted, but treat them as sensitive and do not share them.
- Acceptance tests cost real API calls and create real resources. Use a
  dedicated test org and a least-privilege token, and run `terraform-provider-tester sweep` to
  clean up leaked `tf-acc-test-*` resources.
