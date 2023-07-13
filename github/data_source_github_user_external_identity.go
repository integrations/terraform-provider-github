package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubUserExternalIdentity() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubUserExternalIdentityRead,

		// TODO: fix this
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"organization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Can't set named Keys from what I understand (https://discuss.hashicorp.com/t/custom-provider-how-to-reference-computed-attribute-of-typemap-list-set-defined-as-nested-block/22898/2)
			"saml_identity": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceGithubUserExternalIdentityRead(d *schema.ResourceData, meta interface{}) error {
	username := d.Get("username").(string)

	client := meta.(*Owner).v4client
	orgName := meta.(*Owner).name

	if configuredOrg := d.Get("organization").(string); configuredOrg != "" {
		orgName = configuredOrg
	}

	var query struct {
		Organization struct {
			SamlIdentityProvider struct {
				ExternalIdentities struct {
					Nodes []struct {
						ID   githubv4.String
						User struct {
							Login githubv4.String
						}
						SamlIdentity struct {
							Username githubv4.String
							NameID   githubv4.String
						}
						// TODO: scimIdentity whenever I can actually test that
					}
				} `graphql:"externalIdentities(first: 1, login:$username)"` // TODO: First 1 should be fine?
			}
		} `graphql:"organization(login: $orgName)"`
	}

	variables := map[string]interface{}{
		"orgName":  githubv4.String(orgName),
		"username": githubv4.String(username),
	}

	err := client.Query(meta.(*Owner).StopContext, &query, variables)
	if err != nil {
		return err
	}

	externalIdentity := map[string]string{
		"username": string(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes[0].SamlIdentity.Username),
		"name_id": string(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes[0].SamlIdentity.NameID),
	}

	// TODO: Is this a valid ID?
	d.SetId(string(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes[0].ID))
	d.Set("saml_identity", externalIdentity)
	return nil
}
