## Using artifact attestations to achieve SLSA v1 Build Level 3

This project started to use GitHub Action to create attestations for the release artifacts. Building software with artifact attestation streamlines supply chain security and helps us achieve [SLSA](https://slsa.dev/) v1.0 Build Level 3 for this project.

### Verifying artifact attestations built with a reusable workflow

To verify artifact attestations generated during the build process, use the `gh attestation verify` command from the GitHub CLI.

The `gh attestation verify` command requires either `--owner` or `--repo` flags to be used with it.

> [!NOTE]
> Make sure to replace X.Y.Z with the actual release tag you want to verify.

> [!WARNING]
> Not all artifacts may have attestations generated for them. Please check the [attestations](https://github.com/integrations/terraform-provider-github/attestations) page for this repository to see which artifacts have attestations available.

Download the release artifacts first:

```bash
gh release download vX.Y.Z -R integrations/terraform-provider-github -p "*.zip"
```

To verify the artifact attestations for this project, you can run the following command:

```bash
gh attestation verify --repo integrations/terraform-provider-github terraform-provider-github_X.Y.Z_darwin_amd64.zip
```

### Using optional flags

The `gh attestation verify` command supports additional flags for more specific verification:

Use the `--signer-repo` flag to specify the repository:

```bash
gh attestation verify --owner integrations --signer-repo integrations/terraform-provider-github \
terraform-provider-github_X.Y.Z_darwin_amd64.zip
```

If you would like to require an artifact attestation to be signed with a specific workflow, use the `--signer-workflow` flag to indicate the workflow file that should be used.

```bash
gh attestation verify --owner integrations --signer-workflow integrations/terraform-provider-github/.github/workflows/release.yml \
terraform-provider-github_X.Y.Z_darwin_amd64.zip
```
