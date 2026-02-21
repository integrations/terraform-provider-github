# Architecture Guide

This document serves as a guide for contributors implementing new features and resources in the Terraform Provider for GitHub.

- [Architecture Guide](#architecture-guide)
    - [Module Map](#module-map)
    - [Core Principles](#core-principles)
        - [1. One Resource = One API Entity](#1-one-resource--one-api-entity)
        - [2. Minimize API Calls](#2-minimize-api-calls)
        - [3. API-Only Operations](#3-api-only-operations)
    - [Resource Design](#resource-design)
        - [File Organization](#file-organization)
        - [Resource Structure](#resource-structure)
        - [Schema Field Guidelines](#schema-field-guidelines)
        - [ID Patterns](#id-patterns)
    - [Implementation Patterns](#implementation-patterns)
        - [CRUD Function Signatures](#crud-function-signatures)
        - [Accessing the API Client](#accessing-the-api-client)
        - [Error Handling](#error-handling)
        - [Import](#import)
        - [State Migrations](#state-migrations)
        - [Logging](#logging)
    - [Testing](#testing)
        - [Test Structure](#test-structure)
        - [Test Modes](#test-modes)
        - [Running Tests](#running-tests)
        - [Debugging Tests](#debugging-tests)
    - [Gotchas \& Known Issues](#gotchas--known-issues)
        - [API Preview Headers](#api-preview-headers)
        - [Deprecated Resources](#deprecated-resources)
        - [Known Limitations](#known-limitations)
        - [Workarounds in Code](#workarounds-in-code)
        - [Pending go-github Updates](#pending-go-github-updates)
    - [Appendix](#appendix)
        - [Common Utilities](#common-utilities)
        - [Naming Conventions](#naming-conventions)
    - [Decision Log](#decision-log)
        - [January 2026](#january-2026)
            - [Use StateUpgraders for State Migrations](#use-stateupgraders-for-state-migrations)
            - [Explicit Authentication Configuration](#explicit-authentication-configuration)
            - [Transport Layer Rework](#transport-layer-rework)
            - [No Local Git CLI Support](#no-local-git-cli-support)
        - [2025](#2025)
            - [Replace `log` Package with `tflog`](#replace-log-package-with-tflog)
            - [Finalize SDK v2 Migration](#finalize-sdk-v2-migration)

---

## Module Map

```text
terraform-provider-github/
├── github/
│   ├── provider.go              # Entry point, registers all resources/data sources
│   ├── config.go                # Auth setup, HTTP client, rate limiting, transport
│   │
│   ├── resource_github_*.go     # Resource implementations
│   ├── resource_*_migration.go  # Resource state migration functions (StateUpgraders)
│   ├── data_source_github_*.go  # Data source implementations
│   │
│   │
│   ├── util.go                  # Core utilities (ID parsing, validation)
│   ├── util_*.go                # Domain utilities (rules, labels, etc.)
│   │
│   └── transport.go             # Custom HTTP transport with ETag caching
│
├── ARCHITECTURE.md              # This file - implementation guide
├── MAINTAINERS.md               # Maintainers, decision log, contributors
└── CONTRIBUTING.md              # How to contribute
```

---

## Core Principles

### 1. One Resource = One API Entity

Each Terraform resource should map to a single GitHub API entity. Avoid creating resources that combine multiple API concerns.

**Do:**

```go
// Manages a single repository
func resourceGithubRepository() *schema.Resource { ... }

// Manages repository topics separately
func resourceGithubRepositoryTopics() *schema.Resource { ... }
```

**Don't:**

```go
// Don't combine unrelated API entities into one resource
func resourceGithubRepositoryWithTopicsAndLabels() *schema.Resource { ... }
```

### 2. Minimize API Calls

Each resource should use the minimum number of API calls necessary. Consolidate operations where possible and avoid redundant reads.

**Do:**

```go
func resourceExampleRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    // Single API call to get all needed data
    resource, _, err := client.Resources.Get(ctx, owner, name)
    if err != nil {
        return diag.FromErr(err)
    }
    // Set all attributes from single response
    d.Set("field1", resource.Field1)
    d.Set("field2", resource.Field2)
    return nil
}
```

**Don't:**

```go
func resourceExampleRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    // Don't make separate calls for each field
    field1, _ := client.Resources.GetField1(ctx, owner, name)
    field2, _ := client.Resources.GetField2(ctx, owner, name)
    // ...
}
```

### 3. API-Only Operations

We do not support using local git CLI to operate on repositories. All operations must go through the GitHub API.

**Do:**

```go
// Use go-github client for all git operations
_, _, err := client.Git.CreateRef(ctx, owner, repo, ref)
```

**Don't:**

```go
// Never shell out to git CLI
exec.Command("git", "push", "origin", "main")
```

---

## Resource Design

### File Organization

Resources follow a consistent file naming and organization pattern:

```text
github/
├── resource_github_<entity>.go           # Main resource implementation
├── resource_github_<entity>_test.go      # Acceptance tests
├── resource_github_<entity>_migration.go # State migration functions (if needed)
├── data_source_github_<entity>.go        # Data source implementation
├── data_source_github_<entity>_test.go   # Data source tests
└── util_<domain>.go                      # Domain-specific utilities
```

### Resource Structure

Use context-aware CRUD functions with the `*Context` suffix:

```go
func resourceGithubExample() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceGithubExampleCreate,
        ReadContext:   resourceGithubExampleRead,
        UpdateContext: resourceGithubExampleUpdate,
        DeleteContext: resourceGithubExampleDelete,
        Importer: &schema.ResourceImporter{
            StateContext: resourceGithubExampleImport,
        },

        // Include SchemaVersion and StateUpgraders if state migrations exist
        SchemaVersion: 1,
        StateUpgraders: []schema.StateUpgrader{
            {
                Type:    resourceGithubExampleResourceV0().CoreConfigSchema().ImpliedType(),
                Upgrade: resourceGithubExampleInstanceStateUpgradeV0,
                Version: 0,
            },
        },

        Schema: map[string]*schema.Schema{
            // Schema definition
        },
    }
}
```

### Schema Field Guidelines

Full reference: [Schema Behaviors](https://developer.hashicorp.com/terraform/plugin/sdkv2/schemas/schema-behaviors)

**Primitive options:**

- `Type` — field type (`TypeString`, `TypeBool`, `TypeInt`, `TypeFloat`, `TypeList`, `TypeSet`, `TypeMap`)
- `Description` — human-readable description (always include)
- `Elem` — element type for `TypeList`, `TypeSet`, and `TypeMap` fields (e.g., `&schema.Schema{Type: schema.TypeString}` or a nested `&schema.Resource{}`)
- `Default` — static default value when field is not set in config
- `DefaultFunc` — dynamic default (e.g., read from env var via `schema.EnvDefaultFunc`)

**Behavior flags:**

- `Required` — must be provided in config (mutually exclusive with `Optional`, `Computed`)
- `Optional` — may be omitted from config
- `Computed` — set by the provider (API-derived); combine with `Optional` for optional fields with server defaults
- `ForceNew` — changing this field destroys and recreates the resource
- `Sensitive` — value is masked in plan/state output (secrets, tokens)

**Validation:**

- `ValidateDiagFunc` — validate field value with diagnostics (preferred)

**Constraints:**

- `MaxItems` / `MinItems` — cardinality bounds for `TypeList` and `TypeSet`
- `ConflictsWith` — list of field paths that cannot be set together with this field
- `ExactlyOneOf` — exactly one of these fields must be set
- `AtLeastOneOf` — at least one of these fields must be set
- `RequiredWith` — these fields must all be set if this field is set

**Advanced:**

- `StateFunc` — transform value before storing in state (e.g., normalize to lowercase)
- `DiffSuppressFunc` — suppress plan diffs when old and new values are semantically equal
- `DiffSuppressOnRefresh` — also apply `DiffSuppressFunc` during refresh
- `Set` — custom hash function for `TypeSet` elements
- `Deprecated` — marks field as deprecated with a message shown to users

### ID Patterns

For single-part IDs (most common):

```go
d.SetId(resource.GetName())
```

For composite IDs, use `buildID` to create and `parseID2`/`parseID3`/`parseID4` to parse:

```go
// Two-part ID
id, err := buildID(owner, name)
d.SetId(id)
// Parse:
owner, name, err := parseID2(id)

// Three-part ID
id, err := buildID(owner, repo, name)
d.SetId(id)
// Parse:
owner, repo, name, err := parseID3(id)

// Four-part ID
owner, repo, env, name, err := parseID4(id)

// IDs with special characters (colons, etc.)
id, err := buildID(escapeIDPart(part1), part2)
```

> **Note:** The legacy functions `buildTwoPartID`, `parseTwoPartID`, `buildThreePartID`, and `parseThreePartID` are deprecated. Use `buildID` and `parseID2`/`parseID3`/`parseID4` instead.

---

## Implementation Patterns

### CRUD Function Signatures

```go
func resourceGithubExampleCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    client := meta.(*Owner).v3client
    owner := meta.(*Owner).name

    // Implementation
    return nil // Never call Read at end of Create, set any Computed fields in Create
}

func resourceGithubExampleRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    // Implementation
    return nil
}

func resourceGithubExampleUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    // Implementation
    return nil // Never call Read at end of Update, set any Computed fields in Update
}

func resourceGithubExampleDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    // Implementation
    return nil  // Never call Read after Delete
}
```

### Accessing the API Client

```go
func resourceExampleRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
    meta := m.(*Owner)
    // REST API client (go-github v82)
    client := meta.v3client

    // GraphQL client (for queries not available in REST)
    v4client := meta.v4client

    // Owner context
    owner := meta.name
    isOrg := meta.IsOrganization
    orgID := meta.id

    // ...
}
```

### Error Handling

Handle 404s gracefully by removing from state:

```go
resource, _, err := client.Resources.Get(ctx, owner, name)
if err != nil {
    var ghErr *github.ErrorResponse
    if errors.As(err, &ghErr) {
        if ghErr.Response.StatusCode == http.StatusNotFound {
            log.Printf("[INFO] Removing %s from state because it no longer exists", name)
            d.SetId("")
            return nil
        }
    }
    return diag.FromErr(err)
}
```

Or use the helper function:

```go
if err := deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "resource %s", name); err != nil {
    return diag.FromErr(err)
}
```

### Import

Import is registered via the `Importer` field with a `StateContext` function. After import runs, Terraform **automatically calls `Read`** — so the import function's only job is to set enough state for `Read` to succeed. Do not duplicate `Read` logic in the import function.

For resources with a single-part ID, the default passthrough importer is often sufficient:

```go
Importer: &schema.ResourceImporter{
    StateContext: schema.ImportStatePassthroughContext,
},
```

For resources with composite IDs, the import function must parse the user-provided ID and populate any schema attributes that `Read` depends on:

```go
func resourceGithubExampleImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
    owner, name, err := parseID2(d.Id())
    if err != nil {
        return nil, err
    }

    // Set attributes that Read needs to make API calls
    d.Set("owner", owner)
    // Re-build a normalized ID if needed
    id, err := buildID(owner, name)
    if err != nil {
        return nil, err
    }
    d.SetId(id)

    return []*schema.ResourceData{d}, nil
}
```

**Key principle:** Import sets the minimum state required for `Read` to fetch the full resource. `Read` then populates all remaining attributes.

### State Migrations

When adding new fields or changing schema, use `StateUpgraders` (not the deprecated `MigrateState`):

**Migration file (`resource_github_example_migration.go`):**

```go
// resourceGithubExampleResourceV0 returns the schema for version 0
func resourceGithubExampleResourceV0() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
            // Previous schema version
        },
    }
}

// resourceGithubExampleInstanceStateUpgradeV0 migrates from version 0 to 1
func resourceGithubExampleInstanceStateUpgradeV0(ctx context.Context, rawState map[string]any, meta any) (map[string]any, error) {
    log.Printf("[DEBUG] State before migration: %#v", rawState)

    // Add new field with default value
    if _, ok := rawState["new_field"]; !ok {
        rawState["new_field"] = "default_value"
    }

    log.Printf("[DEBUG] State after migration: %#v", rawState)
    return rawState, nil
}
```

**Register in resource:**

```go
SchemaVersion: 1,
StateUpgraders: []schema.StateUpgrader{
    {
        Type:    resourceGithubExampleResourceV0().CoreConfigSchema().ImpliedType(),
        Upgrade: resourceGithubExampleInstanceStateUpgradeV0,
        Version: 0,
    },
},
```

### Logging

Use `tflog` for structured logging (replacing `log` package):

```go
import "github.com/hashicorp/terraform-plugin-log/tflog"

func resourceExampleCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
    tflog.Debug(ctx, "Creating resource", map[string]any{
        "name":  name,
        "owner": owner,
    })
    // ...
}
```

**Note:** Migration from `log` to `tflog` is in progress. New code should use `tflog`.

---

## Testing

### Test Structure

```go
func TestAccGithubExample(t *testing.T) {
  
  t.Run("creates resource without error", func(t *testing.T) {
        randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
        testResourceName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
        config := fmt.Sprintf(`
            resource "github_example" "test" {
              name = "%s"
            }
        `, testResourceName)

        resource.Test(t, resource.TestCase{
            PreCheck:          func() { skipUnauthenticated(t) },
            ProviderFactories: providerFactories,
            Steps: []resource.TestStep{
                {
                    Config: config,
                    Check: resource.ComposeTestCheckFunc(
                        resource.TestCheckResourceAttr( "github_example.test", "name", testResourceName ),
                    ),
                },
            },
        })
    })
}
```

### Test Modes

Use `skipUnauthenticated(t)`, `skipUnlessHasOrgs(t)`, `skipUnlessHasPaidOrgs(t)`, `skipUnlessEnterprise(t)`, `skipUnlessMode(t, testModes...)` functions to run tests in appropriate contexts:

| Mode           | Environment Variables Required         |
| -------------- | -------------------------------------- |
| `anonymous`    | None (read-only, GitHub.com only)      |
| `individual`   | `GITHUB_TOKEN` + `GITHUB_OWNER`        |
| `organization` | `GITHUB_TOKEN` + `GITHUB_ORGANIZATION` |
| `enterprise`   | Enterprise deployment configuration    |

### Running Tests

```bash
# Run specific test
make testacc T=TestAccGithubExample
```

### Debugging Tests

```bash
# With debug logging
TF_LOG=DEBUG make testacc T=TestAccGithubExample 
```

---

## Gotchas & Known Issues

This section documents provider-specific quirks and known limitations discovered through development.

### API Preview Headers

- Stone Crop preview header is still required for some GraphQL features (`config.go:64-65`)
- The `previewHeaderInjectorTransport` handles automatic header injection for these cases

### Deprecated Resources

The following resources are deprecated and will be removed in future versions:

| Deprecated Resource                          | Replacement                                       |
| -------------------------------------------- | ------------------------------------------------- |
| `github_organization_security_manager`       | `github_organization_role_team`                   |
| `github_organization_custom_role`            | `github_organization_repository_role`             |
| `github_repository_deployment_branch_policy` | `github_repository_environment_deployment_policy` |
| `github_organization_project`                | None (Classic Projects API removed)               |
| `github_project_card`                        | None (Classic Projects API removed)               |
| `github_project_column`                      | None (Classic Projects API removed)               |
| `github_repository_project`                  | None (Classic Projects API removed)               |

### Known Limitations

- **Branch Protection `contexts`**: Deprecated, use the `checks` array instead
- **Runner Groups**: Selected repository IDs are not exposed via API (`resource_github_actions_runner_group.go:179`)
- **Organization Settings**: Test requires manual cleanup (`resource_github_organization_settings_test.go:11`)
- **Repository Search**: Tests may hit rate limits (`data_source_github_repositories_test.go:12`)

### Workarounds in Code

- **EMU with SSO**: Odd behavior with user tokens when using Enterprise Managed Users (`resource_github_enterprise_organization.go:122`)

### Pending go-github Updates

Several features are blocked waiting for go-github v68+:

- `data_source_github_organization_repository_role.go:56`
- `resource_github_organization_repository_role.go:102`
- `data_source_github_organization_role_users.go:41`
- `data_source_github_organization_role_teams.go:51`

---

## Appendix

### Common Utilities

| Function                                                    | Purpose                                        |
| ----------------------------------------------------------- | ---------------------------------------------- |
| `buildID(parts...)`                                         | Create composite ID (e.g., `"a:b"`, `"a:b:c"`) |
| `parseID2(id)`                                              | Parse two-part composite ID                    |
| `parseID3(id)`                                              | Parse three-part composite ID                  |
| `parseID4(id)`                                              | Parse four-part composite ID                   |
| `escapeIDPart(part)`                                        | Escape colons in ID parts                      |
| `wrapErrors([]error)`                                       | Convert errors to diagnostics                  |
| `checkOrganization(meta)`                                   | Verify org context                             |
| `getTeamID(idOrSlug, meta)`                                 | Resolve team ID from ID or slug                |
| `getTeamSlug(idOrSlug, meta)`                               | Resolve team slug from ID or slug              |
| `expandStringList([]any)`                                   | Convert to `[]string`                          |
| `flattenStringList([]string)`                               | Convert to `[]any`                             |
| `deleteResourceOn404AndSwallow304OtherwiseReturnError(...)` | Handle 404/304 responses                       |

### Naming Conventions

| Component            | Pattern                                          | Example                                          |
| -------------------- | ------------------------------------------------ | ------------------------------------------------ |
| Resource function    | `resourceGithub<Entity>`                         | `resourceGithubRepository`                       |
| Data source function | `dataSourceGithub<Entity>`                       | `dataSourceGithubRepository`                     |
| CRUD functions       | `resourceGithub<Entity><Op>`                     | `resourceGithubRepositoryCreate`                 |
| Migration function   | `resourceGithub<Entity>InstanceStateUpgradeV<N>` | `resourceGithubRepositoryInstanceStateUpgradeV0` |
| Schema function      | `resourceGithub<Entity>ResourceV<N>`             | `resourceGithubRepositoryResourceV0`             |
| Test function        | `TestAccGithub<Entity>`                          | `TestAccGithubRepository`                        |
| Utility file         | `util_<domain>.go`                               | `util_rules.go`                                  |

---

## Decision Log

### January 2026

#### Use StateUpgraders for State Migrations

**Decision:** Use `StateUpgraders` instead of the deprecated `MigrateState` function.

**Rationale:** `StateUpgraders` provides a cleaner, more maintainable approach to state migrations that works better with the SDK v2 architecture.

**Implementation:**

- Create `resource_github_<entity>_migration.go` with versioned schema and upgrade functions
- Register in resource with `SchemaVersion` and `StateUpgraders`
- See [ARCHITECTURE.md](ARCHITECTURE.md#state-migrations) for implementation pattern

#### Explicit Authentication Configuration

**Decision:** Make all authentication concerns of the provider entirely explicit. Users must explicitly configure their authentication method.

**Rationale:** Implicit auth detection can lead to confusion and security issues. Explicit configuration makes the provider's behavior predictable and auditable.

**Reference:** <https://github.com/integrations/terraform-provider-github/issues/3116>

#### Transport Layer Rework

**Decision:** Rework the transport layer to utilize:

- [`github-conditional-http-transport`](https://github.com/bored-engineer/github-conditional-http-transport) for conditional requests
- [`go-github-ratelimit`](https://github.com/gofri/go-github-ratelimit) for rate limiting

**Rationale:** These libraries provide better handling of GitHub API rate limits and conditional requests than our current custom implementation.

**Reference:** <https://github.com/integrations/terraform-provider-github/issues/2709#issuecomment-3811466444>

#### No Local Git CLI Support

**Decision:** Do not support using local git CLI to operate on repositories; use purely API operations.

**Rationale:** API-only operations ensure consistency, security, and avoid environment dependencies. The provider should not assume git is installed or configured on the user's machine.

### 2025

#### Replace `log` Package with `tflog`

**Decision:** Replace all usage of the standard `log` package with `tflog` from terraform-plugin-log.

**Rationale:** `tflog` provides structured logging that integrates better with Terraform's logging infrastructure and supports log filtering, structured fields, and proper log levels.

**Migration pattern:**

```go
// Before
log.Printf("[DEBUG] Creating resource: %s", name)

// After
tflog.Debug(ctx, "Creating resource", map[string]any{"name": name})
```

#### Finalize SDK v2 Migration

**Decision:** Complete migration to Terraform Plugin SDK v2.

**Rationale:** SDK v2 provides better diagnostics, context-aware functions, and improved schema validation.

**Key changes:**

- Use `*Context` functions (`CreateContext`, `ReadContext`, etc.)
- Use `ValidateDiagFunc` instead of `ValidateFunc`
- Use `diag.Diagnostics` for error returns
- Use `any` instead of `interface{}`

**Reference:** <https://developer.hashicorp.com/terraform/plugin/sdkv2/guides/v2-upgrade-guide>

---
