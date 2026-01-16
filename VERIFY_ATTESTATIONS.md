# Using Artifact Attestations to Achieve SLSA v1 Build Level 3

This project started to use GitHub Action to create attestations for the release artifacts. Building software with artifact attestation streamlines supply chain security and helps us achieve [SLSA](https://slsa.dev/) v1.0 Build Level 3 for this project.

> [!NOTE]
> Not all artifacts may have attestations generated for them. Please check the [repository attestations](https://github.com/integrations/terraform-provider-github/attestations) to see which artifacts have attestations available.
>
> Attestations are only available for releases from `v6.9.0`.

## Verifying with GitHub CLI

### Prerequisites

First, install GitHub CLI if you haven't already. See the [installation instructions](https://github.com/cli/cli#installation) for your platform.

### Verifying Attestations

To verify artifact attestations generated during the build process, use the `gh attestation verify` command from the GitHub CLI.

The `gh attestation verify` command requires either `--owner` or `--repo` flags to be used with it.

> [!NOTE]
> Make sure to replace x.y.z with the actual release tag you want to verify.
> Replace artifact name with the actual artifact you want to verify.

Download the release artifacts first:

```bash
version="x.y.z"
artifact="terraform-provider-github_${version}_darwin_amd64.zip"

gh release download "v${version}" --repo integrations/terraform-provider-github -p "*.zip" --clobber
```

To verify the artifact attestations for this project, you can run the following command:

```bash
gh attestation verify --repo integrations/terraform-provider-github --source-ref "refs/tags/v${version}"\
  --signer-workflow integrations/terraform-provider-github/.github/workflows/release.yaml@refs/tags/v${version} \
  "$artifact"
```

### Verifying All Artifacts

Alternatively, you can verify all downloaded artifacts with a loop that provides individual status reporting:

```bash
for artifact in terraform-provider-github_${version}_*.zip; do
  echo "Verifying: $artifact"
  gh attestation verify --repo integrations/terraform-provider-github --source-ref "refs/tags/v${version}" \
    --signer-workflow integrations/terraform-provider-github/.github/workflows/release.yaml@refs/tags/v${version} \
    "$artifact" && echo "✓ Verified" || echo "✗ Failed"
done
```

### Using optional flags

The `gh attestation verify` command supports additional flags for more specific verification:

Use the `--signer-repo` flag to specify the repository:

```bash
gh attestation verify --owner integrations --signer-repo \
  integrations/terraform-provider-github \
  "$artifact"
```

If you would like to require an artifact attestation to be signed with a specific workflow, use the `--signer-workflow` flag to indicate the workflow file that should be used.

```bash
gh attestation verify --owner integrations --signer-workflow \
  integrations/terraform-provider-github/.github/workflows/release.yaml@refs/tags/v${version} \
  "$artifact"
```

## Verifying checksums file signature with Cosign and checking artifact integrity

> [!NOTE]
> Not all artifacts may have attestations generated for them. Please check the [repository attestations](https://github.com/integrations/terraform-provider-github/attestations) to see which artifacts have attestations available.
>
> Attestations are only available for releases from `v6.9.0`.

In addition to artifact attestations, you can verify release artifacts using [Cosign](https://docs.sigstore.dev/cosign/overview/). Cosign is a tool for signing and verifying software artifacts and container images.

### Prerequisites

First, install Cosign if you haven't already. See the [installation instructions](https://docs.sigstore.dev/cosign/system_config/installation/) for your platform.

### Verify Checksums File

Download the checksums file and its signature bundle:

```bash
gh release download v${version} --repo integrations/terraform-provider-github \
  -p "terraform-provider-github_${version}_SHA256SUMS" \
  -p "terraform-provider-github_${version}_SHA256SUMS.sbom.json.bundle" --clobber
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
artifact="terraform-provider-github_${version}_darwin_amd64.zip"
gh release download v${version} --repo integrations/terraform-provider-github \
  -p "$artifact" --clobber
```

Verify the checksum:

```bash
shasum -a 256 -c terraform-provider-github_${version}_SHA256SUMS --ignore-missing
```

This will verify that your downloaded artifact matches the signed checksum, confirming its integrity and authenticity.

## Verifying SLSA Provenance Attestations with Cosign

In addition to using the GitHub CLI, you can verify SLSA provenance attestations using Cosign by downloading the attestation and verifying it against your local artifact.

### Prerequisites

1. Install `cosign` for verifying attestations. See the [installation instructions](https://docs.sigstore.dev/cosign/system_config/installation/).
2. Install `gh` (GitHub CLI) if you haven't already. See the [installation instructions](https://github.com/cli/cli#installation).

### Download and Verify Attestation

> [!NOTE]
> Make sure to replace x.y.z with the actual release tag you want to verify.
> Replace artifact name with the actual artifact you want to verify.

> [!NOTE]
> Not all artifacts may have attestations generated for them. Please check the [repository attestations](https://github.com/integrations/terraform-provider-github/attestations) to see which artifacts have attestations available.
>
> Attestations are only available for releases from `v6.9.0`.

First, download the artifact you want to verify:

```bash
version="x.y.z"
artifact="terraform-provider-github_${version}_darwin_amd64.zip"

gh release download "v${version}" --repo integrations/terraform-provider-github \
  -p "$artifact" --clobber
```

Then, download the attestation associated with the artifact:

```bash
gh attestation download "$artifact" \
  --repo integrations/terraform-provider-github
```

This will create a file named `sha256:[digest].jsonl` in the current directory.

Verify the attestation using Cosign:

```bash
# Calculate the digest and verify using the specific bundle file
digest=$(shasum -a 256 "$artifact" | awk '{ print $1 }')
cosign verify-blob-attestation \
  --bundle "sha256:${digest}.jsonl" \
  --new-bundle-format \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  --certificate-identity "https://github.com/integrations/terraform-provider-github/.github/workflows/release.yaml@refs/tags/v${version}" \
  "$artifact"
```

A successful verification will output `Verified OK`, confirming that the artifact was built by the trusted GitHub Actions workflow and its provenance is securely recorded.

### Verifying all release artifacts

To verify all release artifacts for a specific version:

```bash
version="x.y.z"

# Download all release artifacts
gh release download "v${version}" --repo integrations/terraform-provider-github -p "*.zip" --clobber

# Download attestations for all artifacts
for artifact in terraform-provider-github_${version}_*.zip; do
  gh attestation download "$artifact" --repo integrations/terraform-provider-github
done

# Verify all artifacts using specific digest-based bundle files
for artifact in terraform-provider-github_${version}_*.zip; do
  echo "Verifying: $artifact"
  digest=$(shasum -a 256 "$artifact" | awk '{ print $1 }')
  cosign verify-blob-attestation \
    --bundle "sha256:${digest}.jsonl" \
    --new-bundle-format \
    --certificate-oidc-issuer https://token.actions.githubusercontent.com \
    --certificate-identity "https://github.com/integrations/terraform-provider-github/.github/workflows/release.yaml@refs/tags/v${version}" \
    "$artifact" > /dev/null && echo "✓ Verified" || echo "✗ Failed"
done
```

This approach calculates the digest for each artifact and uses the corresponding specific bundle file, ensuring each artifact is verified against its own attestation.
