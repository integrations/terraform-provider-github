# Decision Log

## April 2026

### Replace Legacy Documentation Website

**Decision:** Replace legacy documentation website with generated docs using `github.com/hashicorp/terraform-plugin-docs`.

**Rationale:** Generated documentation ensures consistency, reduces maintenance overhead, and keeps documentation up-to-date with the codebase.

**Implementation:**

- Use `terraform-plugin-docs` to generate documentation from provider schema and examples
- Integrate documentation generation into the build process
- Deprecate and remove legacy documentation website

## January 2026

### Use StateUpgraders for State Migrations

**Decision:** Use `StateUpgraders` instead of the deprecated `MigrateState` function.

**Rationale:** `StateUpgraders` provides a cleaner, more maintainable approach to state migrations that works better with the SDK v2 architecture.

**Implementation:**

- Create `resource_github_<entity>_migration.go` with versioned schema and upgrade functions
- Register in resource with `SchemaVersion` and `StateUpgraders`
- See [ARCHITECTURE.md](ARCHITECTURE.md#state-migrations) for implementation pattern

### Explicit Authentication Configuration

**Decision:** Make all authentication concerns of the provider entirely explicit. Users must explicitly configure their authentication method.

**Rationale:** Implicit auth detection can lead to confusion and security issues. Explicit configuration makes the provider's behavior predictable and auditable.

**Reference:** <https://github.com/integrations/terraform-provider-github/issues/3116>

### Transport Layer Rework

**Decision:** Rework the transport layer to utilize:

- [`github-conditional-http-transport`](https://github.com/bored-engineer/github-conditional-http-transport) for conditional requests
- [`go-github-ratelimit`](https://github.com/gofri/go-github-ratelimit) for rate limiting

**Rationale:** These libraries provide better handling of GitHub API rate limits and conditional requests than our current custom implementation.

**Reference:** <https://github.com/integrations/terraform-provider-github/issues/2709#issuecomment-3811466444>

### Migrate to `terraform-plugin-testing`

**Decision:** Migrate from the SDK testing package (`github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource`) to `terraform-plugin-testing` (`github.com/hashicorp/terraform-plugin-testing`). Use `ConfigStateChecks` and `ConfigPlanChecks` as the preferred assertion patterns, replacing `Check:` + `resource.ComposeTestCheckFunc`.

**Rationale:** `terraform-plugin-testing` is the standalone testing framework that decouples test utilities from the SDK. `ConfigStateChecks` and `ConfigPlanChecks` provide type-safe, composable assertions with better error messages.

**Reference:** <https://developer.hashicorp.com/terraform/plugin/testing>

### No Local Git CLI Support

**Decision:** Do not support using local git CLI to operate on repositories; use purely API operations.

**Rationale:** API-only operations ensure consistency, security, and avoid environment dependencies. The provider should not assume git is installed or configured on the user's machine.

## 2025

### Replace `log` Package with `tflog`

**Decision:** Replace all usage of the standard `log` package with `tflog` from terraform-plugin-log.

**Rationale:** `tflog` provides structured logging that integrates better with Terraform's logging infrastructure and supports log filtering, structured fields, and proper log levels.

**Migration pattern:**

```go
// Before
log.Printf("[DEBUG] Creating resource: %s", name)

// After
tflog.Debug(ctx, "Creating resource", map[string]any{"name": name})
```

### Finalize SDK v2 Migration

**Decision:** Complete migration to Terraform Plugin SDK v2.

**Rationale:** SDK v2 provides better diagnostics, context-aware functions, and improved schema validation.

**Key changes:**

- Use `*Context` functions (`CreateContext`, `ReadContext`, etc.)
- Use `ValidateDiagFunc` instead of `ValidateFunc`
- Use `diag.Diagnostics` for error returns
- Use `any` instead of `interface{}`

**Reference:** <https://developer.hashicorp.com/terraform/plugin/sdkv2/guides/v2-upgrade-guide>

---
