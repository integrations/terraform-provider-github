package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubOrganizationExternalIdentities() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationExternalIdentitiesRead,

		Schema: map[string]*schema.Schema{
			"identities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationExternalIdentitiesRead(d *schema.ResourceData, meta interface{}) error {
	name := meta.(*Owner).name

	client4 := meta.(*Owner).v4client
	ctx := meta.(*Owner).StopContext

	var query struct {
		Organization struct {
			SamlIdentityProvider struct {
				ExternalIdentities struct {
					Edges []struct {
						Node struct {
							User struct {
								Login githubv4.String
							}
							SamlIdentity struct {
								NameId githubv4.String
							}
						}
					}
					PageInfo struct {
						EndCursor   githubv4.String
						HasNextPage bool
					}
				} `graphql:"externalIdentities(first: 100, after: $after)"`
			}
		} `graphql:"organization(login: $login)"`
	}
	variables := map[string]interface{}{
		"login": githubv4.String(name),
		"after": (*githubv4.String)(nil),
	}

	var identities []map[string]string

	for {
		err := client4.Query(ctx, &query, variables)
		if err != nil {
			return err
		}
		for _, edge := range query.Organization.SamlIdentityProvider.ExternalIdentities.Edges {
			identities = append(identities, map[string]string{
				"login":              string(edge.Node.User.Login),
				"samlIdentityNameId": string(edge.Node.SamlIdentity.NameId),
			})
		}
		if !query.Organization.SamlIdentityProvider.ExternalIdentities.PageInfo.HasNextPage {
			break
		}
		variables["after"] = githubv4.NewString(query.Organization.SamlIdentityProvider.ExternalIdentities.PageInfo.EndCursor)
	}

	d.SetId(name)
	d.Set("identities", identities)

	return nil
}
