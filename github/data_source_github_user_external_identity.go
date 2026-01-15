package github

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubUserExternalIdentity() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubUserExternalIdentityRead,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
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
			"login": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubUserExternalIdentityRead(d *schema.ResourceData, meta any) error {
	username := d.Get("username").(string)

	client := meta.(*Owner).v4client
	orgName := meta.(*Owner).name

	var query struct {
		Organization struct {
			SamlIdentityProvider struct {
				ExternalIdentities `graphql:"externalIdentities(first: 1, login:$username)"` // There should only ever be one external identity configured
			}
		} `graphql:"organization(login: $orgName)"`
	}

	variables := map[string]any{
		"orgName":  githubv4.String(orgName),
		"username": githubv4.String(username),
	}

	err := client.Query(meta.(*Owner).StopContext, &query, variables)
	if err != nil {
		return err
	}
	if len(query.Organization.SamlIdentityProvider.Edges) == 0 {
		return fmt.Errorf("there was no external identity found for username %q in Organization %q", username, orgName)
	}

	externalIdentityNode := query.Organization.SamlIdentityProvider.ExternalIdentities.Edges[0].Node // There should only be one user in this list

	samlIdentity := map[string]string{
		"family_name": string(externalIdentityNode.SamlIdentity.FamilyName),
		"given_name":  string(externalIdentityNode.SamlIdentity.GivenName),
		"name_id":     string(externalIdentityNode.SamlIdentity.NameId),
		"username":    string(externalIdentityNode.SamlIdentity.Username),
	}

	scimIdentity := map[string]string{
		"family_name": string(externalIdentityNode.ScimIdentity.FamilyName),
		"given_name":  string(externalIdentityNode.ScimIdentity.GivenName),
		"username":    string(externalIdentityNode.ScimIdentity.Username),
	}

	login := string(externalIdentityNode.User.Login)

	d.SetId(fmt.Sprintf("%s/%s", orgName, username))
	_ = d.Set("saml_identity", samlIdentity)
	_ = d.Set("scim_identity", scimIdentity)
	_ = d.Set("login", login)
	_ = d.Set("username", login)
	return nil
}
