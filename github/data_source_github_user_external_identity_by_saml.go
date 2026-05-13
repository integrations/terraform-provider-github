package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubUserExternalIdentityBySaml() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubUserExternalIdentityBySamlRead,

		Schema: map[string]*schema.Schema{
			"saml_name_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The SAML NameID (typically an email address) to look up.",
			},
			"login": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The GitHub username linked to this SAML identity.",
			},
			"username": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The GitHub username linked to this SAML identity (same as login).",
			},
			"saml_identity": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The SAML identity attributes.",
			},
			"scim_identity": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The SCIM identity attributes.",
			},
		},
	}
}

func dataSourceGithubUserExternalIdentityBySamlRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	samlNameId := d.Get("saml_name_id").(string)

	client := meta.(*Owner).v4client
	orgName := meta.(*Owner).name

	var query struct {
		Organization struct {
			SamlIdentityProvider struct {
				ExternalIdentities `graphql:"externalIdentities(first: 1, userName:$userName)"`
			}
		} `graphql:"organization(login: $orgName)"`
	}

	variables := map[string]any{
		"orgName":  githubv4.String(orgName),
		"userName": githubv4.String(samlNameId),
	}

	err := client.Query(meta.(*Owner).StopContext, &query, variables)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(query.Organization.SamlIdentityProvider.Edges) == 0 {
		return diag.Errorf("no external identity found for SAML NameID %q in organization %q", samlNameId, orgName)
	}

	node := query.Organization.SamlIdentityProvider.ExternalIdentities.Edges[0].Node

	samlIdentity := map[string]string{
		"family_name": string(node.SamlIdentity.FamilyName),
		"given_name":  string(node.SamlIdentity.GivenName),
		"name_id":     string(node.SamlIdentity.NameId),
		"username":    string(node.SamlIdentity.Username),
	}

	scimIdentity := map[string]string{
		"family_name": string(node.ScimIdentity.FamilyName),
		"given_name":  string(node.ScimIdentity.GivenName),
		"username":    string(node.ScimIdentity.Username),
	}

	login := string(node.User.Login)

	d.SetId(fmt.Sprintf("%s/%s", orgName, samlNameId))
	_ = d.Set("saml_identity", samlIdentity)
	_ = d.Set("scim_identity", scimIdentity)
	_ = d.Set("login", login)
	_ = d.Set("username", login)
	return nil
}
