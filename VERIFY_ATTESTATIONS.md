# Using artifact attestations to achieve SLSA v1 Build Level 3

This project started to use GitHub Action to create attestations for the release artifacts. Building software with artifact attestation streamlines supply chain security and helps us achieve [SLSA](https://slsa.dev/) v1.0 Build Level 3 for this project.

## Verifying release artifacts attestations with GitHub CLI

> [!WARNING]
> Not all artifacts may have attestations generated for them. Please check the [attestations](https://github.com/integrations/terraform-provider-github/attestations) page for this repository to see which artifacts have attestations available.

> [!CAUTION]
> The attestations are available only for the releases created since the version `v6.9.0` of this project.

### Prerequisites

First, install GitHub CLI if you haven't already. See the [installation instructions](https://github.com/cli/cli#installation) for your platform.

### Verifying attestations

To verify artifact attestations generated during the build process, use the `gh attestation verify` command from the GitHub CLI.

The `gh attestation verify` command requires either `--owner` or `--repo` flags to be used with it.

> [!NOTE]
> Make sure to replace x.y.z with the actual release tag you want to verify.

Download the release artifacts first:

```bash
version="x.y.z"
gh release download "v${version}" --repo integrations/terraform-provider-github -p "*.zip"
```

To verify the artifact attestations for this project, you can run the following command:

```bash
gh attestation verify --repo integrations/terraform-provider-github --source-ref "v${version}"\
  --signer-workflow integrations/terraform-provider-github/.github/workflows/release.yaml \
  "terraform-provider-github_${version}_darwin_amd64.zip"
```

### Using optional flags

The `gh attestation verify` command supports additional flags for more specific verification:

Use the `--signer-repo` flag to specify the repository:

```bash
gh attestation verify --owner integrations --signer-repo \
  integrations/terraform-provider-github \
  terraform-provider-github_${version}_darwin_amd64.zip
```

If you would like to require an artifact attestation to be signed with a specific workflow, use the `--signer-workflow` flag to indicate the workflow file that should be used.

```bash
gh attestation verify --owner integrations --signer-workflow \
  integrations/terraform-provider-github/.github/workflows/release.yaml \
  terraform-provider-github_${version}_darwin_amd64.zip
```

## Verifying release artifacts with Cosign

> [!WARNING]
> Not all the releases may have Cosign signature for the checksum files.

> [!CAUTION]
> The Cosign signatures are available only for the releases created since the version `v6.9.0` of this project.

In addition to artifact attestations, you can verify release artifacts using [Cosign](https://docs.sigstore.dev/cosign/overview/). Cosign is a tool for signing and verifying software artifacts and container images.

### Prerequisites

First, install Cosign if you haven't already. See the [installation instructions](https://docs.sigstore.dev/cosign/system_config/installation/) for your platform.

### Verify checksums file

> [!NOTE]
> Make sure to replace X.Y.Z with the actual release tag you want to verify.

Download the checksums file and its signature bundle:

```bash
gh release download v${version} --repo integrations/terraform-provider-github \
  -p "terraform-provider-github_${version}_SHA256SUMS" \
  -p "terraform-provider-github_${version}_SHA256SUMS.sbom.json.bundle"
```

Verify the checksums file signature:

```bash
cosign verify-blob \
  --bundle "terraform-provider-github_${version}_SHA256SUMS.sbom.json.bundle" \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  --certificate-identity "https://github.com/integrations/terraform-provider-github/.github/workflows/release.yaml@refs/tags/v${version}" \
  "terraform-provider-github_${version}_SHA256SUMS"
```

### Verify artifact checksums

After verifying the checksums file, verify your downloaded artifacts match the checksums:

Download the artifact you want to verify:

```bash
gh release download v${version} --repo integrations/terraform-provider-github \
  -p "terraform-provider-github_${version}_darwin_amd64.zip"
```

Verify the checksum:

```bash
shasum -a 256 -c terraform-provider-github_${version}_SHA256SUMS --ignore-missing
```

This will verify that your downloaded artifact matches the signed checksum, confirming its integrity and authenticity.
