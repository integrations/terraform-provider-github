package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func isSAMLEnforcementError(err error) bool {
	if err == nil {
		return false
	}

	if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
		return ghErr.Response.StatusCode == 403 && strings.Contains(ghErr.Message, "SAML enforcement")
	}

	return strings.Contains(err.Error(), "Resource protected by organization SAML enforcement")
}

// organizationClients returns clients that can act on the organization. With a GitHub App these are
// the clients of the app installation on that organization; otherwise the provider clients.
func organizationClients(ctx context.Context, meta *Owner, orgName string) (*github.Client, *githubv4.Client, error) {
	if meta.appSource == nil {
		return meta.v3client, meta.v4client, nil
	}

	v3, err := meta.appSource.OwnerRESTClient(ctx, orgName)
	if err != nil {
		return nil, nil, err
	}

	v4, err := meta.appSource.OwnerGraphQLClient(ctx, orgName)
	if err != nil {
		return nil, nil, err
	}

	return v3, v4, nil
}

// installAppOnOrganization installs the authenticated GitHub App on an enterprise-owned organization
// so that the app can manage the organization it just created.
func installAppOnOrganization(ctx context.Context, meta *Owner, enterpriseId, orgName string) error {
	appClient, err := meta.appSource.RESTClient()
	if err != nil {
		return err
	}

	app, _, err := appClient.Apps.Get(ctx, "")
	if err != nil {
		return fmt.Errorf("could not read the authenticated GitHub App: %w", err)
	}

	enterpriseSlug, err := getEnterpriseSlug(ctx, meta.v4client, enterpriseId)
	if err != nil {
		return err
	}

	_, _, err = meta.v3client.Enterprise.InstallApp(ctx, enterpriseSlug, orgName, github.InstallAppRequest{
		ClientID:            app.GetClientID(),
		RepositorySelection: "all",
	})
	if err != nil {
		return fmt.Errorf("could not install GitHub App %q on organization %q: %w", app.GetSlug(), orgName, err)
	}

	return nil
}

func getEnterpriseSlug(ctx context.Context, v4 *githubv4.Client, enterpriseId string) (string, error) {
	var query struct {
		Node struct {
			Enterprise struct {
				Slug githubv4.String
			} `graphql:"... on Enterprise"`
		} `graphql:"node(id: $id)"`
	}

	err := v4.Query(ctx, &query, map[string]any{"id": githubv4.ID(enterpriseId)})
	if err != nil {
		return "", err
	}

	return string(query.Node.Enterprise.Slug), nil
}

func resourceGithubEnterpriseOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubEnterpriseOrganizationCreate,
		Read:   resourceGithubEnterpriseOrganizationRead,
		Delete: resourceGithubEnterpriseOrganizationDelete,
		Update: resourceGithubEnterpriseOrganizationUpdate,
		Importer: &schema.ResourceImporter{
			State: resourceGithubEnterpriseOrganizationImport,
		},
		Schema: map[string]*schema.Schema{
			"enterprise_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the enterprise.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The node ID of the organization.",
			},
			"database_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The database ID of the organization.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the organization.",
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The display name of the organization.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the organization.",
			},
			"admin_logins": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of organization owner usernames.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"billing_email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The billing email address.",
			},
		},
	}
}

func resourceGithubEnterpriseOrganizationCreate(data *schema.ResourceData, m any) error {
	meta, _ := m.(*Owner)
	var mutate struct {
		CreateEnterpriseOrganization struct {
			Organization struct {
				ID         githubv4.ID
				DatabaseId githubv4.Int
			}
		} `graphql:"createEnterpriseOrganization(input:$input)"`
	}

	v3 := meta.v3client
	v4 := meta.v4client

	var adminLogins []githubv4.String
	for _, v := range data.Get("admin_logins").(*schema.Set).List() {
		adminLogins = append(adminLogins, githubv4.String(v.(string)))
	}

	profileName := data.Get("display_name").(string)
	if profileName == "" {
		profileName = data.Get("name").(string)
	}

	input := githubv4.CreateEnterpriseOrganizationInput{
		EnterpriseID: data.Get("enterprise_id"),
		Login:        githubv4.String(data.Get("name").(string)),
		ProfileName:  githubv4.String(profileName),
		BillingEmail: githubv4.String(data.Get("billing_email").(string)),
		AdminLogins:  adminLogins,
	}

	err := v4.Mutate(context.Background(), &mutate, input, nil)
	if err != nil {
		return err
	}
	data.SetId(fmt.Sprintf("%s", mutate.CreateEnterpriseOrganization.Organization.ID))

	// The provider does not read after write, so database_id has to be populated here or it stays
	// unset until the next refresh, which breaks same-apply references such as
	// github_enterprise_actions_runner_group.selected_organization_ids.
	if err := data.Set("database_id", mutate.CreateEnterpriseOrganization.Organization.DatabaseId); err != nil {
		return err
	}

	// We use the V3 api to set the description of the org, because there is no mutator in the V4 API to edit the org's
	// description

	//NOTE: There is some odd behavior here when using an EMU with SSO. If the user token has been granted permission to
	//ANY ORG in the enterprise, then this works, provided that our token has sufficient permission. If the user token
	//has not been added to any orgs, then this will fail.
	//
	//Unfortunately, there is no way in the api to grant a token permission to access an org. This needs to be done
	//via the UI. This means our resource will work fine if the user has sufficient admin permissions and at least one
	//org exists. It also means that we can't use terraform to automate creation of the very first org in an enterprise.
	//That sucks a little, but seems like a restriction we can live with.
	//
	//It would be nice if there was an API available in github to enable a token for SSO.

	ctx := context.Background()
	orgName := data.Get("name").(string)

	if meta.appSource != nil {
		err = installAppOnOrganization(ctx, meta, data.Get("enterprise_id").(string), orgName)
		if err != nil {
			return err
		}
		v3, _, err = organizationClients(ctx, meta, orgName)
		if err != nil {
			return err
		}
	}

	description := data.Get("description").(string)
	if description != "" {
		_, _, err = v3.Organizations.Edit(
			ctx,
			orgName,
			&github.Organization{
				Description: new(description),
			},
		)
		if err != nil {
			if isSAMLEnforcementError(err) {
				// The org was created but we can't set description until the PAT is authorized.
				// Clear it from state so next plan will show drift and retry after PAT authorization.
				log.Printf("[WARN] Organization %q created but could not set description due to SAML enforcement. Authorize the PAT and run apply again.", orgName)
				_ = data.Set("description", "")
				return nil
			}
			return err
		}
	}
	return nil
}

func resourceGithubEnterpriseOrganizationRead(data *schema.ResourceData, m any) error {
	meta, _ := m.(*Owner)
	v4 := meta.v4client
	ctx := context.Background()

	var query struct {
		Node struct {
			Organization struct {
				ID          githubv4.ID
				DatabaseId  githubv4.Int
				Name        githubv4.String
				Login       githubv4.String
				Description githubv4.String
			} `graphql:"... on Organization"`
		} `graphql:"node(id: $id)"`
	}

	err := v4.Query(ctx, &query, map[string]any{"id": data.Id()})
	if err != nil {
		if strings.Contains(err.Error(), "Could not resolve to a node with the global id") {
			log.Printf("[INFO] Removing organization (%s) from state because it no longer exists in GitHub", data.Id())
			data.SetId("")
			return nil
		}
		return err
	}

	err = data.Set("name", query.Node.Organization.Login)
	if err != nil {
		return err
	}

	if query.Node.Organization.Name != query.Node.Organization.Login {
		err = data.Set("display_name", query.Node.Organization.Name)
		if err != nil {
			return err
		}
	}

	err = data.Set("database_id", query.Node.Organization.DatabaseId)
	if err != nil {
		return err
	}

	err = data.Set("description", query.Node.Organization.Description)
	if err != nil {
		return err
	}

	orgName := string(query.Node.Organization.Login)
	orgV3, orgV4, err := organizationClients(ctx, meta, orgName)
	if err != nil {
		return err
	}

	adminLogins, err := readOrganizationAdmins(ctx, orgV4, data.Id(), meta.maxPerPage)
	if err != nil {
		return err
	}

	err = data.Set("admin_logins", schema.NewSet(schema.HashString, adminLogins))
	if err != nil {
		return err
	}

	// The GraphQL field organizationBillingEmail is not available to app installations; REST is.
	org, _, err := orgV3.Organizations.Get(ctx, orgName)
	if err != nil {
		return err
	}

	return data.Set("billing_email", org.GetBillingEmail())
}

func readOrganizationAdmins(ctx context.Context, v4 *githubv4.Client, orgId string, perPage int) ([]any, error) {
	var query struct {
		Node struct {
			Organization struct {
				MembersWithRole struct {
					Edges []struct {
						User struct {
							Login githubv4.String
						} `graphql:"node"`
						Role githubv4.String
					} `graphql:"edges"`
					PageInfo PageInfo
				} `graphql:"membersWithRole(first:$first, after:$cursor)"`
			} `graphql:"... on Organization"`
		} `graphql:"node(id: $id)"`
	}

	variables := map[string]any{
		"id":     orgId,
		"first":  githubv4.Int(perPage),
		"cursor": (*githubv4.String)(nil),
	}

	var adminLogins []any

	for {
		err := v4.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}

		for _, v := range query.Node.Organization.MembersWithRole.Edges {
			if v.Role == "ADMIN" {
				adminLogins = append(adminLogins, string(v.User.Login))
			}
		}

		if !query.Node.Organization.MembersWithRole.PageInfo.HasNextPage {
			break
		}

		variables["cursor"] = new(query.Node.Organization.MembersWithRole.PageInfo.EndCursor)
	}

	return adminLogins, nil
}

func resourceGithubEnterpriseOrganizationDelete(data *schema.ResourceData, m any) error {
	meta, _ := m.(*Owner)
	ctx := context.WithValue(context.Background(), ctxId, data.Id())

	v3, _, err := organizationClients(ctx, meta, data.Get("name").(string))
	if err != nil {
		return err
	}

	_, err = v3.Organizations.Delete(ctx, data.Get("name").(string))

	// We expect the delete to return with a 202 Accepted error so ignore those
	acceptedError := &github.AcceptedError{}
	if errors.As(err, &acceptedError) {
		return nil
	}

	return err
}

func resourceGithubEnterpriseOrganizationImport(data *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	parts := strings.Split(data.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid ID specified: supplied ID must be written as <enterprise_slug>/<org_name>")
	}

	v4 := meta.v4client
	ctx := context.Background()

	enterpriseId, err := getEnterpriseID(ctx, v4, parts[0])
	if err != nil {
		return nil, err
	}
	_ = data.Set("enterprise_id", enterpriseId)

	// An enterprise app installation cannot resolve organizations by login, only through the enterprise.
	orgId, err := getOrganizationId(ctx, v4, parts[1])
	if err != nil {
		orgId, err = getEnterpriseOrganizationId(ctx, v4, parts[0], parts[1])
	}
	if err != nil {
		return nil, err
	}
	data.SetId(orgId)

	err = resourceGithubEnterpriseOrganizationRead(data, meta)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{data}, nil
}

func getOrganizationId(ctx context.Context, v4 *githubv4.Client, orgName string) (string, error) {
	var query struct {
		Organization struct {
			Id githubv4.String
		} `graphql:"organization(login: $orgName)"`
	}

	err := v4.Query(ctx, &query, map[string]any{"orgName": githubv4.String(orgName)})
	if err != nil {
		return "", err
	}
	return string(query.Organization.Id), nil
}

func getEnterpriseOrganizationId(ctx context.Context, v4 *githubv4.Client, enterpriseSlug, orgName string) (string, error) {
	var query struct {
		Enterprise struct {
			Organizations struct {
				Nodes []struct {
					Id    githubv4.String
					Login githubv4.String
				}
			} `graphql:"organizations(first: 1, query: $orgName)"`
		} `graphql:"enterprise(slug: $slug)"`
	}

	err := v4.Query(ctx, &query, map[string]any{
		"slug":    githubv4.String(enterpriseSlug),
		"orgName": githubv4.String(orgName),
	})
	if err != nil {
		return "", err
	}

	for _, node := range query.Enterprise.Organizations.Nodes {
		if string(node.Login) == orgName {
			return string(node.Id), nil
		}
	}

	return "", fmt.Errorf("organization %q not found in enterprise %q", orgName, enterpriseSlug)
}

func updateDescription(ctx context.Context, data *schema.ResourceData, v3 *github.Client) error {
	orgName := data.Get("name").(string)
	oldDesc, newDesc := stringChanges(data.GetChange("description"))

	if oldDesc != newDesc {
		_, _, err := v3.Organizations.Edit(
			ctx,
			orgName,
			&github.Organization{
				Description: new(newDesc),
			},
		)
		if err != nil {
			if isSAMLEnforcementError(err) {
				// Reset state to old value so next plan shows drift
				log.Printf("[WARN] Could not update description for %q due to SAML enforcement. Authorize the PAT and run apply again.", orgName)
				_ = data.Set("description", oldDesc)
				return nil
			}
			return err
		}
	}
	return nil
}

func updateDisplayName(ctx context.Context, data *schema.ResourceData, v4 *github.Client) error {
	orgName := data.Get("name").(string)
	oldDisplayName, newDisplayName := stringChanges(data.GetChange("display_name"))

	if oldDisplayName != newDisplayName {
		_, _, err := v4.Organizations.Edit(
			ctx,
			orgName,
			&github.Organization{
				Name: new(newDisplayName),
			},
		)
		if err != nil {
			if isSAMLEnforcementError(err) {
				// Reset state to old value so next plan shows drift
				log.Printf("[WARN] Could not update display_name for %q due to SAML enforcement. Authorize the PAT and run apply again.", orgName)
				_ = data.Set("display_name", oldDisplayName)
				return nil
			}
			return err
		}
	}
	return nil
}

func removeUsers(ctx context.Context, v3 *github.Client, v4 *githubv4.Client, toRemove []any, orgName string) error {
	for _, user := range toRemove {
		err := removeUser(ctx, v3, v4, user.(string), orgName)
		if err != nil {
			return err
		}
	}
	return nil
}

func removeUser(ctx context.Context, v3 *github.Client, v4 *githubv4.Client, user, orgName string) error {
	//How we remove an admin user from an enterprise organization depends on if the user is a member of any teams.
	//If they are a member of any teams, we shouldn't delete them, instead we edit their membership role to be
	//'MEMBER' instead of 'ADMIN'. If the user is not a member of any teams, then we remove from the org.

	// First, use the v4 API to count how many teams the user is in
	var query struct {
		Organization struct {
			Teams struct {
				TotalCount githubv4.Int
			} `graphql:"teams(first:1, userLogins:[$user])"`
		} `graphql:"organization(login: $org)"`
	}

	err := v4.Query(
		ctx,
		&query,
		map[string]any{
			"org":  githubv4.String(orgName),
			"user": githubv4.String(user),
		},
	)
	if err != nil {
		return err
	}

	if query.Organization.Teams.TotalCount == 0 {
		_, err = v3.Organizations.RemoveOrgMembership(ctx, user, orgName)
		return err
	}

	membership, _, err := v3.Organizations.GetOrgMembership(ctx, user, orgName)
	if err != nil {
		return err
	}

	membership.Role = new("member")
	_, _, err = v3.Organizations.EditOrgMembership(ctx, user, orgName, membership)
	return err
}

// Owners are added through an enterprise mutation (enterpriseV4) and looked up or removed through the
// organization (v3, v4).
func updateAdminList(ctx context.Context, data *schema.ResourceData, orgName string, v3 *github.Client, v4, enterpriseV4 *githubv4.Client) error {
	oldSet, newSet := setChanges(data.GetChange("admin_logins"))
	toRemove := oldSet.Difference(newSet).List()
	toAdd := newSet.Difference(oldSet).List()

	err := addUsers(ctx, data, v4, enterpriseV4, toAdd)
	if err != nil {
		return err
	}

	return removeUsers(ctx, v3, v4, toRemove, orgName)
}

func addUsers(ctx context.Context, data *schema.ResourceData, v4, enterpriseV4 *githubv4.Client, toAdd []any) error {
	if len(toAdd) != 0 {
		var mutate struct {
			AddEnterpriseOrganizationMember struct {
				Ignored string `graphql:"clientMutationId"`
			} `graphql:"addEnterpriseOrganizationMember(input: $input)"`
		}

		adminRole := githubv4.OrganizationMemberRoleAdmin
		userIds, err := getUserIds(v4, toAdd)
		if err != nil {
			return err
		}

		input := githubv4.AddEnterpriseOrganizationMemberInput{
			EnterpriseID:   data.Get("enterprise_id"),
			OrganizationID: data.Id(),
			UserIDs:        userIds,
			Role:           &adminRole,
		}

		err = enterpriseV4.Mutate(ctx, &mutate, input, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateBillingEmail(ctx context.Context, data *schema.ResourceData, orgName string, v3 *github.Client) error {
	oldBilling, newBilling := stringChanges(data.GetChange("billing_email"))
	if oldBilling != newBilling {
		_, _, err := v3.Organizations.Edit(
			ctx,
			orgName,
			&github.Organization{
				BillingEmail: &newBilling,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceGithubEnterpriseOrganizationUpdate(data *schema.ResourceData, m any) error {
	meta, _ := m.(*Owner)
	ctx := context.Background()
	orgName := data.Get("name").(string)

	v3, v4, err := organizationClients(ctx, meta, orgName)
	if err != nil {
		return err
	}

	err = updateDisplayName(ctx, data, v3)
	if err != nil {
		return err
	}

	err = updateDescription(ctx, data, v3)
	if err != nil {
		return err
	}

	err = updateAdminList(ctx, data, orgName, v3, v4, meta.v4client)
	if err != nil {
		return err
	}

	return updateBillingEmail(ctx, data, orgName, v3)
}

func getUserIds(v4 *githubv4.Client, loginNames []any) ([]githubv4.ID, error) {
	var query struct {
		User struct {
			ID githubv4.String
		} `graphql:"user(login: $login)"`
	}

	var ret []githubv4.ID

	for _, l := range loginNames {
		err := v4.Query(context.Background(), &query, map[string]any{"login": githubv4.String(l.(string))})
		if err != nil {
			return nil, err
		}
		ret = append(ret, query.User.ID)
	}
	return ret, nil
}

func stringChanges(oldValue, newValue any) (string, string) {
	oldString, _ := oldValue.(string)
	newString, _ := newValue.(string)

	return oldString, newString
}

func setChanges(oldValue, newValue any) (*schema.Set, *schema.Set) {
	oldSet, _ := oldValue.(*schema.Set)
	newSet, _ := newValue.(*schema.Set)

	if oldSet == nil {
		oldSet = schema.NewSet(schema.HashString, nil)
	}

	if newSet == nil {
		newSet = schema.NewSet(schema.HashString, nil)
	}

	return oldSet, newSet
}
