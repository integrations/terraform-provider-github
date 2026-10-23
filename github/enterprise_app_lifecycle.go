package github

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"

	"github.com/integrations/terraform-provider-github/v6/internal/ghclient"
)

var (
	errEnterpriseInventoryAbsent    = errors.New("organization absent from enterprise inventory")
	errAutomationInstallationAbsent = errors.New("automation installation absent")
)

func enterprisePropagationError(err error) bool {
	if errors.Is(err, errEnterpriseInventoryAbsent) || errors.Is(err, errAutomationInstallationAbsent) {
		return true
	}
	response, ok := errors.AsType[*github.ErrorResponse](err)
	return ok && response.Response != nil && response.Response.StatusCode == http.StatusNotFound
}

type enterpriseAppIdentity struct {
	ID          int64             `json:"id"`
	ClientID    string            `json:"client_id"`
	Permissions map[string]string `json:"permissions"`
}

type enterpriseAppInstallation struct {
	ID      int64 `json:"id"`
	AppID   int64 `json:"app_id"`
	Account struct {
		ID    int64  `json:"id"`
		Login string `json:"login"`
		Slug  string `json:"slug"`
	} `json:"account"`
	Enterprise struct {
		Slug string `json:"slug"`
	} `json:"enterprise"`
	TargetType  string            `json:"target_type"`
	SuspendedAt *time.Time        `json:"suspended_at"`
	Permissions map[string]string `json:"permissions"`
}

type enterpriseInventoryOrganization struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
}

func enterpriseAppRequest(ctx context.Context, client *github.Client, method, path string, body, result any) (*github.Response, error) {
	request, err := client.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}
	return client.Do(request, result)
}

func positiveAppID(value, field string) (int64, error) {
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil || parsed <= 0 {
		return 0, fmt.Errorf("%s must be a positive int64 ID", field)
	}
	return parsed, nil
}

func configureEnterpriseApp(ctx context.Context, version string, config *Config) (*Owner, error) {
	if config.AppID == nil || config.AppInstallationID == nil || len(config.AppPEM) == 0 || config.Token != "" {
		return nil, fmt.Errorf("enterprise_app_auth requires creation App credentials and forbids PAT authentication")
	}
	opts := ghclient.ClientOptions{BaseURL: config.BaseURL.String(), IsGHES: config.IsGHES, UserAgent: providerName + "/" + version}
	creationInstallation, err := positiveAppID(*config.AppInstallationID, "app_auth.installation_id")
	if err != nil {
		return nil, err
	}
	installerInstallation, err := positiveAppID(config.EnterpriseApp.installerInstallationID, "installer_installation_id")
	if err != nil {
		return nil, err
	}
	if _, err := positiveAppID(config.EnterpriseApp.automationID, "automation_app_id"); err != nil {
		return nil, err
	}
	lifecycle := &enterpriseAppLifecycle{config: config.EnterpriseApp, opts: opts}
	lifecycle.creationJWT, err = ghclient.NewAppRESTClient(*config.AppID, config.AppPEM, nil, opts)
	if err != nil {
		return nil, fmt.Errorf("invalid creation App credentials")
	}
	lifecycle.installerJWT, err = ghclient.NewAppRESTClient(lifecycle.config.installerID, []byte(lifecycle.config.installerPEM), nil, opts)
	if err != nil {
		return nil, fmt.Errorf("invalid installer App credentials")
	}
	lifecycle.automationJWT, err = ghclient.NewAppRESTClient(lifecycle.config.automationID, []byte(lifecycle.config.automationPEM), nil, opts)
	if err != nil {
		return nil, fmt.Errorf("invalid automation App credentials")
	}
	lifecycle.installer, err = ghclient.NewAppRESTClient(lifecycle.config.installerID, []byte(lifecycle.config.installerPEM), &installerInstallation, opts)
	if err != nil {
		return nil, err
	}
	owner := &Owner{enterpriseApp: lifecycle, maxPerPage: config.MaxPerPage}
	owner.v3client, err = ghclient.NewAppRESTClient(*config.AppID, config.AppPEM, &creationInstallation, opts)
	if err != nil {
		return nil, err
	}
	owner.v4client, err = ghclient.NewAppGraphQLClient(*config.AppID, config.AppPEM, &creationInstallation, opts)
	if err != nil {
		return nil, err
	}
	if err := lifecycle.validateEnterpriseInstallation(ctx, lifecycle.creationJWT, *config.AppID, creationInstallation, ""); err != nil {
		return nil, fmt.Errorf("creation App: %w", err)
	}
	if err := lifecycle.validateEnterpriseInstallation(ctx, lifecycle.installerJWT, lifecycle.config.installerID, installerInstallation, "enterprise_organization_installations"); err != nil {
		return nil, fmt.Errorf("installer App: %w", err)
	}
	lifecycle.automationPermissions, err = lifecycle.validateAutomationApp(ctx)
	if err != nil {
		return nil, err
	}
	lifecycle.enterpriseID, err = getEnterpriseID(ctx, owner.v4client, lifecycle.config.slug)
	if err != nil || lifecycle.enterpriseID == "" {
		return nil, fmt.Errorf("cannot validate enterprise %q using creation App: %v", lifecycle.config.slug, err)
	}
	lifecycle.orgClients = func(installationID int64) (*github.Client, *githubv4.Client, error) {
		rest, err := ghclient.NewAppRESTClient(lifecycle.config.automationID, []byte(lifecycle.config.automationPEM), &installationID, opts)
		if err != nil {
			return nil, nil, err
		}
		graphql, err := ghclient.NewAppGraphQLClient(lifecycle.config.automationID, []byte(lifecycle.config.automationPEM), &installationID, opts)
		return rest, graphql, err
	}
	return owner, nil
}

func (lifecycle *enterpriseAppLifecycle) validateEnterpriseInstallation(ctx context.Context, client *github.Client, issuer string, installationID int64, requiredPermission string) error {
	var app enterpriseAppIdentity
	if _, err := enterpriseAppRequest(ctx, client, http.MethodGet, "app", nil, &app); err != nil {
		return fmt.Errorf("cannot verify App JWT identity: %w", err)
	}
	if app.ID <= 0 || (issuer != strconv.FormatInt(app.ID, 10) && issuer != app.ClientID) {
		return fmt.Errorf("App JWT identity does not match configured issuer")
	}
	var installation enterpriseAppInstallation
	if _, err := enterpriseAppRequest(ctx, client, http.MethodGet, fmt.Sprintf("app/installations/%d", installationID), nil, &installation); err != nil {
		return err
	}
	enterpriseSlug := installation.Account.Slug
	if enterpriseSlug == "" {
		enterpriseSlug = installation.Enterprise.Slug
	}
	if enterpriseSlug == "" {
		enterpriseSlug = installation.Account.Login
	}
	if installation.ID != installationID || installation.AppID != app.ID || installation.SuspendedAt != nil || installation.TargetType != "Enterprise" ||
		!strings.EqualFold(enterpriseSlug, lifecycle.config.slug) {
		return fmt.Errorf("installation must be active, belong to this App, and target configured enterprise %q; no installation changes were made", lifecycle.config.slug)
	}
	if requiredPermission != "" && installation.Permissions[requiredPermission] != "write" {
		return fmt.Errorf("installer App installation requires %s=write; creation blocked, no installation changes were made", requiredPermission)
	}
	return nil
}

func checkAutomationPermissions(permissions map[string]string) error {
	if permissions["organization_administration"] != "write" || (permissions["members"] != "read" && permissions["members"] != "write") {
		return fmt.Errorf("automation App requires organization_administration=write and members=read or write")
	}
	for permission, level := range permissions {
		if level != "" && level != "none" && permission != "metadata" && permission != "members" && !strings.HasPrefix(permission, "organization_") {
			return fmt.Errorf("automation App requests unsupported permission %q; repository permissions are not supported with repository_selection=none", permission)
		}
	}
	return nil
}

func (lifecycle *enterpriseAppLifecycle) validateAutomationApp(ctx context.Context) (map[string]string, error) {
	var app enterpriseAppIdentity
	if _, err := enterpriseAppRequest(ctx, lifecycle.automationJWT, http.MethodGet, "app", nil, &app); err != nil {
		return nil, fmt.Errorf("cannot verify automation App JWT identity: %w", err)
	}
	if strconv.FormatInt(app.ID, 10) != lifecycle.config.automationID || app.ClientID != lifecycle.config.automationClientID {
		return nil, fmt.Errorf("automation App numeric ID, client ID and private key do not identify the same App")
	}
	if err := checkAutomationPermissions(app.Permissions); err != nil {
		return nil, err
	}
	return app.Permissions, nil
}

func (lifecycle *enterpriseAppLifecycle) inventoryOrganization(ctx context.Context, name string) (*enterpriseInventoryOrganization, error) {
	var found *enterpriseInventoryOrganization
	for page := 1; ; page++ {
		var organizations []enterpriseInventoryOrganization
		path := fmt.Sprintf("enterprises/%s/apps/installable_organizations?per_page=100&page=%d", url.PathEscape(lifecycle.config.slug), page)
		response, err := enterpriseAppRequest(ctx, lifecycle.installer, http.MethodGet, path, nil, &organizations)
		if err != nil {
			return nil, fmt.Errorf("cannot verify enterprise inventory; state retained: %w", err)
		}
		for _, organization := range organizations {
			if strings.EqualFold(organization.Login, name) {
				if found != nil || organization.ID <= 0 {
					return nil, fmt.Errorf("ambiguous enterprise inventory for %q; state retained", name)
				}
				found = new(organization)
			}
		}
		if response.NextPage == 0 && len(organizations) < 100 {
			break
		}
		if response.NextPage > page {
			page = response.NextPage - 1
		}
	}
	if found == nil {
		return nil, fmt.Errorf("%w for %q; verify membership and installer access, state retained", errEnterpriseInventoryAbsent, name)
	}
	return found, nil
}

func (lifecycle *enterpriseAppLifecycle) installation(ctx context.Context, name string) (*enterpriseAppInstallation, bool, error) {
	var installation enterpriseAppInstallation
	response, err := enterpriseAppRequest(ctx, lifecycle.automationJWT, http.MethodGet, "orgs/"+url.PathEscape(name)+"/installation", nil, &installation)
	if err != nil {
		if response != nil && response.StatusCode == http.StatusNotFound {
			return nil, true, nil
		}
		return nil, false, err
	}
	if installation.ID <= 0 || strconv.FormatInt(installation.AppID, 10) != lifecycle.config.automationID || installation.SuspendedAt != nil || !strings.EqualFold(installation.Account.Login, name) {
		return nil, false, fmt.Errorf("automation installation is suspended or has an unexpected App/account identity; no changes were made")
	}
	if err := checkAutomationPermissions(installation.Permissions); err != nil {
		return nil, false, err
	}
	for permission, requested := range lifecycle.automationPermissions {
		granted := installation.Permissions[permission]
		if requested == "write" && granted != "write" || requested == "read" && granted != "read" && granted != "write" {
			return nil, false, fmt.Errorf("automation installation has not granted requested permission %s=%s; approve permissions externally, no installation upgrade was attempted", permission, requested)
		}
	}
	return &installation, false, nil
}

func (lifecycle *enterpriseAppLifecycle) organizationClients(ctx context.Context, name string, databaseID int64) (*github.Client, *githubv4.Client, int64, error) {
	installation, missing, err := lifecycle.installation(ctx, name)
	if err != nil {
		return nil, nil, 0, err
	}
	if missing {
		return nil, nil, 0, fmt.Errorf("%w; install the configured App through approved bootstrap, then import/refresh; read and import never install Apps", errAutomationInstallationAbsent)
	}
	if installation.Account.ID != databaseID {
		return nil, nil, 0, fmt.Errorf("automation installation organization ID mismatch; state retained")
	}
	rest, graphql, err := lifecycle.orgClients(installation.ID)
	return rest, graphql, installation.ID, err
}

func (lifecycle *enterpriseAppLifecycle) checkEnterprise(data *schema.ResourceData) error {
	if data.Get("enterprise_id").(string) != lifecycle.enterpriseID {
		return fmt.Errorf("resource enterprise_id does not match configured enterprise %q; no changes were made", lifecycle.config.slug)
	}
	return nil
}

func (lifecycle *enterpriseAppLifecycle) read(ctx context.Context, data *schema.ResourceData) error {
	if err := lifecycle.checkEnterprise(data); err != nil {
		return err
	}
	name := data.Get("name").(string)
	inventory, err := lifecycle.inventoryOrganization(ctx, name)
	if err != nil {
		return err
	}
	if stored := data.Get("database_id").(int); stored != 0 && int64(stored) != inventory.ID {
		return fmt.Errorf("enterprise inventory database ID mismatch; state retained")
	}
	rest, graphql, installationID, err := lifecycle.organizationClients(ctx, name, inventory.ID)
	if err != nil {
		return err
	}
	organization, _, err := rest.Organizations.Get(ctx, name)
	if err != nil {
		return fmt.Errorf("cannot read organization; state retained: %w", err)
	}
	if organization.GetNodeID() == "" || organization.GetNodeID() != data.Id() || organization.GetID() != inventory.ID || !strings.EqualFold(organization.GetLogin(), name) {
		return fmt.Errorf("REST organization node/database/login identity mismatch; state retained")
	}
	if organization.BillingEmail == nil {
		return fmt.Errorf("organization billing_email is missing or null; check automation App organization administration permission; state retained")
	}
	var owners []any
	var cursor *githubv4.String
	for {
		var query struct {
			Organization struct {
				ID              githubv4.ID
				DatabaseID      githubv4.Int
				Login           githubv4.String
				MembersWithRole struct {
					Edges []struct {
						Node struct{ Login githubv4.String }
						Role githubv4.String
					}
					PageInfo PageInfo
				} `graphql:"membersWithRole(first:100, after:$cursor)"`
			} `graphql:"organization(login:$login)"`
		}
		if err := graphql.Query(ctx, &query, map[string]any{"login": githubv4.String(name), "cursor": cursor}); err != nil {
			return fmt.Errorf("cannot read organization owners; state retained: %w", err)
		}
		if fmt.Sprint(query.Organization.ID) != data.Id() || int64(query.Organization.DatabaseID) != inventory.ID || !strings.EqualFold(string(query.Organization.Login), name) {
			return fmt.Errorf("GraphQL organization identity mismatch; state retained")
		}
		for _, edge := range query.Organization.MembersWithRole.Edges {
			if edge.Role == "ADMIN" {
				owners = append(owners, string(edge.Node.Login))
			}
		}
		if !query.Organization.MembersWithRole.PageInfo.HasNextPage {
			break
		}
		next := query.Organization.MembersWithRole.PageInfo.EndCursor
		if next == "" || (cursor != nil && next == *cursor) {
			return fmt.Errorf("invalid owners pagination cursor; state retained")
		}
		cursor = new(next)
	}
	displayName := organization.GetName()
	if displayName == organization.GetLogin() && data.Get("display_name").(string) == "" {
		displayName = ""
	}
	for field, value := range map[string]any{
		"name": organization.GetLogin(), "database_id": organization.GetID(), "billing_email": *organization.BillingEmail,
		"display_name": displayName, "description": organization.GetDescription(),
		"admin_logins": owners, "automation_installation_id": installationID,
	} {
		if err := data.Set(field, value); err != nil {
			return err
		}
	}
	return nil
}

func (lifecycle *enterpriseAppLifecycle) importOrganization(ctx context.Context, data *schema.ResourceData, owner *Owner) ([]*schema.ResourceData, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	parts := strings.Split(data.Id(), "/")
	if len(parts) != 2 || !strings.EqualFold(parts[0], lifecycle.config.slug) || parts[1] == "" {
		return nil, fmt.Errorf("import ID must be <configured enterprise_slug>/<organization login>")
	}
	enterpriseID, err := getEnterpriseID(ctx, owner.v4client, parts[0])
	if err != nil {
		return nil, err
	}
	if enterpriseID != lifecycle.enterpriseID {
		return nil, fmt.Errorf("enterprise identity changed; import blocked")
	}
	inventory, err := lifecycle.inventoryOrganization(ctx, parts[1])
	if err != nil {
		return nil, err
	}
	rest, _, _, err := lifecycle.organizationClients(ctx, parts[1], inventory.ID)
	if err != nil {
		return nil, err
	}
	organization, _, err := rest.Organizations.Get(ctx, parts[1])
	if err != nil {
		return nil, err
	}
	if organization.GetID() != inventory.ID || organization.GetNodeID() == "" || !strings.EqualFold(organization.GetLogin(), parts[1]) {
		return nil, fmt.Errorf("organization identity mismatch; import blocked")
	}
	for field, value := range map[string]any{"enterprise_id": enterpriseID, "name": organization.GetLogin(), "database_id": organization.GetID()} {
		if err := data.Set(field, value); err != nil {
			return nil, err
		}
	}
	data.SetId(organization.GetNodeID())
	if err := lifecycle.read(ctx, data); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{data}, nil
}

func (lifecycle *enterpriseAppLifecycle) create(ctx context.Context, data *schema.ResourceData, owner *Owner) error {
	if err := lifecycle.checkEnterprise(data); err != nil {
		return err
	}
	permissions, err := lifecycle.validateAutomationApp(ctx)
	if err != nil {
		return err
	}
	if !maps.Equal(permissions, lifecycle.automationPermissions) {
		return fmt.Errorf("automation App requested permissions changed since provider configuration; re-plan before creation")
	}
	installerID, err := positiveAppID(lifecycle.config.installerInstallationID, "installer_installation_id")
	if err != nil {
		return err
	}
	if err := lifecycle.validateEnterpriseInstallation(ctx, lifecycle.installerJWT, lifecycle.config.installerID, installerID, "enterprise_organization_installations"); err != nil {
		return err
	}
	var organizations []enterpriseInventoryOrganization
	if _, err := enterpriseAppRequest(ctx, lifecycle.installer, http.MethodGet, "enterprises/"+url.PathEscape(lifecycle.config.slug)+"/apps/installable_organizations?per_page=1", nil, &organizations); err != nil {
		return fmt.Errorf("installer cannot access enterprise inventory; creation blocked: %w", err)
	}
	var mutation struct {
		CreateEnterpriseOrganization struct {
			Organization struct {
				ID         githubv4.ID
				DatabaseID githubv4.Int
			}
		} `graphql:"createEnterpriseOrganization(input:$input)"`
	}
	var admins []githubv4.String
	for _, login := range data.Get("admin_logins").(*schema.Set).List() {
		admins = append(admins, githubv4.String(login.(string)))
	}
	err = owner.v4client.Mutate(ctx, &mutation, githubv4.CreateEnterpriseOrganizationInput{
		EnterpriseID: lifecycle.enterpriseID, Login: githubv4.String(data.Get("name").(string)),
		ProfileName: githubv4.String(data.Get("name").(string)), BillingEmail: githubv4.String(data.Get("billing_email").(string)), AdminLogins: admins,
	}, nil)
	created := mutation.CreateEnterpriseOrganization.Organization
	if nodeID, ok := created.ID.(string); ok && nodeID != "" {
		data.SetId(nodeID)
		if setErr := data.Set("database_id", int(created.DatabaseID)); setErr != nil {
			return setErr
		}
	}
	if err != nil {
		return fmt.Errorf("enterprise organization creation failed or has an ambiguous outcome; do not retry blindly, check enterprise inventory and import %s/%s if created: %w", lifecycle.config.slug, data.Get("name"), err)
	}
	if data.Id() == "" || created.DatabaseID <= 0 {
		return fmt.Errorf("creation returned incomplete organization identity; check enterprise inventory and import %s/%s before retrying", lifecycle.config.slug, data.Get("name"))
	}
	if err := lifecycle.bootstrap(ctx, data); err != nil {
		return fmt.Errorf("organization %s/%s was created (ID %s) but bootstrap failed; ID retained. Correct App permissions/installation, then import %s/%s if Terraform did not persist state; do not recreate: %w", lifecycle.config.slug, data.Get("name"), data.Id(), lifecycle.config.slug, data.Get("name"), err)
	}
	return nil
}

func propagationWait(ctx context.Context, attempt int) error {
	delay := min(time.Second*time.Duration(1<<min(attempt, 3)), 8*time.Second)
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func (lifecycle *enterpriseAppLifecycle) enterpriseListsAutomationInstallation(ctx context.Context, name string) (bool, error) {
	for page := 1; ; page++ {
		var installations []struct {
			ID       int64  `json:"id"`
			ClientID string `json:"client_id"`
		}
		path := fmt.Sprintf("enterprises/%s/apps/organizations/%s/installations?per_page=100&page=%d", url.PathEscape(lifecycle.config.slug), url.PathEscape(name), page)
		response, err := enterpriseAppRequest(ctx, lifecycle.installer, http.MethodGet, path, nil, &installations)
		if err != nil {
			return false, fmt.Errorf("cannot verify existing enterprise organization installations; installation POST blocked: %w", err)
		}
		for _, installation := range installations {
			if installation.ID <= 0 || installation.ClientID == "" {
				return false, fmt.Errorf("enterprise organization installation listing has incomplete identity; installation POST blocked")
			}
			if installation.ClientID == lifecycle.config.automationClientID {
				return true, nil
			}
		}
		if response.NextPage == 0 && len(installations) < 100 {
			return false, nil
		}
		if response.NextPage > page {
			page = response.NextPage - 1
		}
	}
}

func (lifecycle *enterpriseAppLifecycle) bootstrap(ctx context.Context, data *schema.ResourceData) error {
	name := data.Get("name").(string)
	var inventory *enterpriseInventoryOrganization
	var err error
	for attempt := range 8 {
		inventory, err = lifecycle.inventoryOrganization(ctx, name)
		if err == nil {
			break
		}
		if !enterprisePropagationError(err) || attempt == 7 {
			return err
		}
		if waitErr := propagationWait(ctx, attempt); waitErr != nil {
			return waitErr
		}
	}
	if err != nil {
		return err
	}
	if inventory.ID != int64(data.Get("database_id").(int)) {
		return fmt.Errorf("created organization does not match enterprise inventory; installation blocked")
	}
	_, missing, err := lifecycle.installation(ctx, name)
	if err != nil {
		return err
	}
	if missing {
		listed, err := lifecycle.enterpriseListsAutomationInstallation(ctx, name)
		if err != nil {
			return err
		}
		if listed {
			return fmt.Errorf("enterprise already lists the automation App installation but the App cannot access it; inspect suspension, grants and propagation externally, no installation POST or upgrade was attempted")
		}
		body := struct {
			ClientID            string `json:"client_id"`
			RepositorySelection string `json:"repository_selection"`
		}{lifecycle.config.automationClientID, "none"}
		path := "enterprises/" + url.PathEscape(lifecycle.config.slug) + "/apps/organizations/" + url.PathEscape(name) + "/installations"
		if _, err := enterpriseAppRequest(ctx, lifecycle.installer, http.MethodPost, path, body, nil); err != nil {
			return fmt.Errorf("installation request failed or outcome is ambiguous; inspect installation before retrying (no automatic POST retry): %w", err)
		}
	}
	var rest *github.Client
	for attempt := range 8 {
		var installationID int64
		rest, _, installationID, err = lifecycle.organizationClients(ctx, name, inventory.ID)
		if err == nil {
			organization, _, getErr := rest.Organizations.Get(ctx, name)
			err = getErr
			if err == nil {
				if organization.GetNodeID() != data.Id() || organization.GetID() != inventory.ID {
					return fmt.Errorf("automation organization identity mismatch; profile changes blocked")
				}
				if err := data.Set("automation_installation_id", installationID); err != nil {
					return err
				}
				break
			}
		}
		if !enterprisePropagationError(err) || attempt == 7 {
			return err
		}
		if waitErr := propagationWait(ctx, attempt); waitErr != nil {
			return waitErr
		}
	}
	if err != nil {
		return err
	}
	if data.Get("description").(string) != "" || data.Get("display_name").(string) != "" {
		profile := &github.Organization{}
		if description := data.Get("description").(string); description != "" {
			profile.Description = new(description)
		}
		if displayName := data.Get("display_name").(string); displayName != "" {
			profile.Name = new(displayName)
		}
		_, _, err = rest.Organizations.Edit(ctx, name, profile)
	}
	return err
}
