package github

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

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

func resourceGithubEnterpriseOrganizationCreate(data *schema.ResourceData, meta any) error {
	var mutate struct {
		CreateEnterpriseOrganization struct {
			Organization struct {
				ID githubv4.ID
			}
		} `graphql:"createEnterpriseOrganization(input:$input)"`
	}

	owner := meta.(*Owner)
	v3 := owner.v3client
	v4 := owner.v4client

	var adminLogins []githubv4.String
	for _, v := range data.Get("admin_logins").(*schema.Set).List() {
		adminLogins = append(adminLogins, githubv4.String(v.(string)))
	}

	input := githubv4.CreateEnterpriseOrganizationInput{
		EnterpriseID: data.Get("enterprise_id"),
		Login:        githubv4.String(data.Get("name").(string)),
		ProfileName:  githubv4.String(data.Get("name").(string)),
		BillingEmail: githubv4.String(data.Get("billing_email").(string)),
		AdminLogins:  adminLogins,
	}

	err := v4.Mutate(context.Background(), &mutate, input, nil)
	if err != nil {
		return err
	}
	data.SetId(fmt.Sprintf("%s", mutate.CreateEnterpriseOrganization.Organization.ID))

	// We use the V3 api to set the description of the org, because there is no mutator in the V4 API to edit the org's
	// description and display name

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

	description := data.Get("description").(string)
	displayName := data.Get("display_name").(string)
	if description != "" || displayName != "" {
		_, _, err = v3.Organizations.Edit(
			context.Background(),
			data.Get("name").(string),
			&github.Organization{
				Description: github.String(description),
				Name:        github.String(displayName),
			},
		)
		return err
	}
	return nil
}

func resourceGithubEnterpriseOrganizationRead(data *schema.ResourceData, meta any) error {
	var query struct {
		Node struct {
			Organization struct {
				ID                       githubv4.ID
				DatabaseId               githubv4.Int
				Name                     githubv4.String
				Login                    githubv4.String
				Description              githubv4.String
				OrganizationBillingEmail githubv4.String
				MembersWithRole          struct {
					Edges []struct {
						User struct {
							Login githubv4.String
						} `graphql:"node"`
						Role githubv4.String
					} `graphql:"edges"`
					PageInfo PageInfo
				} `graphql:"membersWithRole(first:100, after:$cursor)"`
			} `graphql:"... on Organization"`
		} `graphql:"node(id: $id)"`
	}

	variables := map[string]any{
		"id":     data.Id(),
		"cursor": (*githubv4.String)(nil),
	}

	var adminLogins []any

	owner := meta.(*Owner)

	for {
		v4 := owner.v4client
		err := v4.Query(context.Background(), &query, variables)
		if err != nil {
			if strings.Contains(err.Error(), "Could not resolve to a node with the global id") {
				// The GraphQL error "Could not resolve to a node" can mean either:
				// 1. The org was actually deleted
				// 2. The org exists but the PAT hasn't been authorized for it yet (EMU/SSO)
				//
				// In EMU/SSO environments, both GraphQL and REST may return not-found errors
				// for orgs that exist but the PAT isn't authorized to access. We cannot
				// reliably distinguish "deleted" from "unauthorized" based on API responses.
				//
				// To avoid incorrectly removing the org from state (which causes Terraform
				// to destroy it), we always return an error and let the user investigate.
				// If the org was truly deleted, use: terraform state rm <resource_address>
				orgName := data.Get("name").(string)
				v3 := owner.v3client
				ctx := context.WithValue(context.Background(), ctxId, data.Id())
				_, _, restErr := v3.Organizations.Get(ctx, orgName)

				if restErr == nil {
					// REST succeeded - org definitely exists, GraphQL access issue
					return fmt.Errorf("organization %q exists but cannot be read via GraphQL. This typically occurs when the PAT has not been authorized for the organization yet. Please authorize the PAT via GitHub UI and retry. Original error: %w", orgName, err)
				}

				// REST also failed - could be deleted OR unauthorized
				// Do NOT remove from state to avoid accidental destruction
				return fmt.Errorf("cannot read organization %q via GraphQL or REST API. If the organization was deleted, remove it from state with: terraform state rm <resource_address>. GraphQL error: %v, REST error: %w", orgName, err, restErr)
			}
			return err
		}

		for _, v := range query.Node.Organization.MembersWithRole.Edges {
			if v.Role == "ADMIN" {
				adminLogins = append(adminLogins, string(v.User.Login))
			}
		}

		if !query.Node.Organization.MembersWithRole.PageInfo.HasNextPage {
			break
		}

		variables["cursor"] = githubv4.NewString(query.Node.Organization.MembersWithRole.PageInfo.EndCursor)
	}

	err := data.Set("admin_logins", schema.NewSet(schema.HashString, adminLogins))
	if err != nil {
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

	err = data.Set("billing_email", query.Node.Organization.OrganizationBillingEmail)
	if err != nil {
		return err
	}

	err = data.Set("database_id", query.Node.Organization.DatabaseId)
	if err != nil {
		return err
	}

	err = data.Set("description", query.Node.Organization.Description)
	return err
}

func resourceGithubEnterpriseOrganizationDelete(data *schema.ResourceData, meta any) error {
	owner := meta.(*Owner)
	v3 := owner.v3client

	ctx := context.WithValue(context.Background(), ctxId, data.Id())

	_, err := v3.Organizations.Delete(ctx, data.Get("name").(string))

	// We expect the delete to return with a 202 Accepted error so ignore those
	acceptedError := &github.AcceptedError{}
	if errors.As(err, &acceptedError) {
		return nil
	}

	return err
}

func resourceGithubEnterpriseOrganizationImport(data *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	parts := strings.Split(data.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid ID specified: supplied ID must be written as <enterprise_slug>/<org_name>")
	}

	v4 := meta.(*Owner).v4client
	ctx := context.Background()

	enterpriseId, err := getEnterpriseId(ctx, v4, parts[0])
	if err != nil {
		return nil, err
	}
	_ = data.Set("enterprise_id", enterpriseId)

	orgId, err := getOrganizationId(ctx, v4, parts[1])
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

func getEnterpriseId(ctx context.Context, v4 *githubv4.Client, enterpriseSlug string) (string, error) {
	var query struct {
		Enterprise struct {
			ID githubv4.String
		} `graphql:"enterprise(slug: $enterpriseSlug)"`
	}

	err := v4.Query(ctx, &query, map[string]any{"enterpriseSlug": githubv4.String(enterpriseSlug)})
	if err != nil {
		return "", err
	}
	return string(query.Enterprise.ID), nil
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

func updateDescription(ctx context.Context, data *schema.ResourceData, v3 *github.Client) error {
	orgName := data.Get("name").(string)
	oldDesc, newDesc := stringChanges(data.GetChange("description"))

	if oldDesc != newDesc {
		_, _, err := v3.Organizations.Edit(
			ctx,
			orgName,
			&github.Organization{
				Description: github.String(data.Get("description").(string)),
			},
		)
		return err
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
				Name: github.String(data.Get("display_name").(string)),
			},
		)
		return err
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

	membership.Role = github.String("member")
	_, _, err = v3.Organizations.EditOrgMembership(ctx, user, orgName, membership)
	return err
}

func updateAdminList(ctx context.Context, data *schema.ResourceData, orgName string, v3 *github.Client, v4 *githubv4.Client) error {
	oldSet, newSet := setChanges(data.GetChange("admin_logins"))
	toRemove := oldSet.Difference(newSet).List()
	toAdd := newSet.Difference(oldSet).List()

	err := addUsers(ctx, data, v4, toAdd)
	if err != nil {
		return err
	}

	return removeUsers(ctx, v3, v4, toRemove, orgName)
}

func addUsers(ctx context.Context, data *schema.ResourceData, v4 *githubv4.Client, toAdd []any) error {
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

		err = v4.Mutate(ctx, &mutate, input, nil)
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

func resourceGithubEnterpriseOrganizationUpdate(data *schema.ResourceData, meta any) error {
	v3 := meta.(*Owner).v3client
	v4 := meta.(*Owner).v4client
	ctx := context.Background()

	err := updateDisplayName(ctx, data, v3)
	if err != nil {
		return err
	}

	err = updateDescription(ctx, data, v3)
	if err != nil {
		return err
	}

	orgName := data.Get("name").(string)
	err = updateAdminList(ctx, data, orgName, v3, v4)
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
