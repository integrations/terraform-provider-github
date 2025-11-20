package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

type ExternalIdentities struct {
	Edges []struct {
		Node struct {
			User struct {
				Login githubv4.String
			}
			SamlIdentity struct {
				NameId     githubv4.String
				Username   githubv4.String
				GivenName  githubv4.String
				FamilyName githubv4.String
			}
			ScimIdentity struct {
				Username   githubv4.String
				GivenName  githubv4.String
				FamilyName githubv4.String
			}
		}
	}
	PageInfo struct {
		EndCursor   githubv4.String
		HasNextPage bool
	}
}

func dataSourceGithubOrganizationExternalIdentities() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationExternalIdentitiesRead,

		Schema: map[string]*schema.Schema{
			"identities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"login": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"saml_identity": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"scim_identity": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationExternalIdentitiesRead(d *schema.ResourceData, meta any) error {
	name := meta.(*Owner).name

	client4 := meta.(*Owner).v4client
	ctx := meta.(*Owner).StopContext

	var query struct {
		Organization struct {
			SamlIdentityProvider struct {
				ExternalIdentities `graphql:"externalIdentities(first: 100, after: $after)"`
			}
		} `graphql:"organization(login: $login)"`
	}
	variables := map[string]any{
		"login": githubv4.String(name),
		"after": (*githubv4.String)(nil),
	}

	var identities []map[string]any

	for {
		err := client4.Query(ctx, &query, variables)
		if err != nil {
			return err
		}
		for _, edge := range query.Organization.SamlIdentityProvider.Edges {
			identity := map[string]any{
				"login":         string(edge.Node.User.Login),
				"saml_identity": nil,
				"scim_identity": nil,
			}

			if edge.Node.SamlIdentity.NameId != "" {
				identity["saml_identity"] = map[string]string{
					"name_id":     string(edge.Node.SamlIdentity.NameId),
					"username":    string(edge.Node.SamlIdentity.Username),
					"given_name":  string(edge.Node.SamlIdentity.GivenName),
					"family_name": string(edge.Node.SamlIdentity.FamilyName),
				}
			}

			if edge.Node.ScimIdentity.Username != "" {
				identity["scim_identity"] = map[string]string{
					"username":    string(edge.Node.ScimIdentity.Username),
					"given_name":  string(edge.Node.ScimIdentity.GivenName),
					"family_name": string(edge.Node.ScimIdentity.FamilyName),
				}
			}

			identities = append(identities, identity)
		}
		if !query.Organization.SamlIdentityProvider.PageInfo.HasNextPage {
			break
		}
		variables["after"] = githubv4.NewString(query.Organization.SamlIdentityProvider.PageInfo.EndCursor)
	}

	d.SetId(name)
	_ = d.Set("identities", identities)

	return nil
}
