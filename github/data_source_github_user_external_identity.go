package github

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			"scim_identity": {
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

	var query struct {
		Organization struct {
			SamlIdentityProvider struct {
				ExternalIdentities `graphql:"externalIdentities(first: 1, login:$username)"` // There should only ever be one external identity configured
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
	if len(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes) == 0 {
		return fmt.Errorf("There was no external identity found for username %q in Organization %q", username, orgName)
	}

	samlIdentity := map[string]string{
		"family_name": string(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes[0].SamlIdentity.FamilyName),
		"given_name": string(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes[0].SamlIdentity.GivenName),
		"name_id": string(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes[0].SamlIdentity.NameID),
		"username": string(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes[0].SamlIdentity.Username),
	}

	scimIdentity := map[string]string{
		"family_name": string(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes[0].ScimIdentity.FamilyName),
		"given_name": string(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes[0].ScimIdentity.GivenName),
		"username": string(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes[0].ScimIdentity.Username),
	}

	// TODO: Is this a valid ID?
	d.SetId(string(query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes[0].ID))
	d.Set("saml_identity", samlIdentity)
	d.Set("scim_identity", scimIdentity)
	return nil
}
