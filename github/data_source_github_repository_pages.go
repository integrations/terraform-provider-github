package github

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryPages() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve GitHub Pages configuration for a repository.",
		ReadContext: dataSourceGithubRepositoryPagesRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The repository name to get GitHub Pages information for.",
			},
			// TODO: Uncomment this when we are ready to support owner fields properly. https://github.com/integrations/terraform-provider-github/pull/3166#discussion_r2816053082
			// "owner": {
			// 	Type:        schema.TypeString,
			// 	Required:    true,
			// 	Description: "The owner of the repository.",
			// },
			"source": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The source branch and directory for the rendered Pages site.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repository branch used to publish the site's source files.",
						},
						"path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repository directory from which the site publishes.",
						},
					},
				},
			},
			"build_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of GitHub Pages site. Can be 'legacy' or 'workflow'.",
			},
			"cname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The custom domain for the repository.",
			},
			"custom_404": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the rendered GitHub Pages site has a custom 404 page.",
			},
			"html_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL (with scheme) to the rendered GitHub Pages site.",
			},
			"build_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The GitHub Pages site's build status e.g. 'building' or 'built'.",
			},
			"api_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The API URL of the GitHub Pages resource.",
			},
			"public": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.",
			},
			"https_enforced": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the rendered GitHub Pages site will only be served over HTTPS.",
			},
		},
	}
}

func dataSourceGithubRepositoryPagesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client

	owner := meta.name // TODO: Add owner support // d.Get("owner").(string)
	repoName := d.Get("repository").(string)

	pages, resp, err := client.Repositories.GetPagesInfo(ctx, owner, repoName)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return diag.Errorf("GitHub Pages not found for repository %s/%s", owner, repoName)
		}
		return diag.Errorf("error reading repository pages: %s", err.Error())
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(int(repo.GetID())))

	if err := d.Set("build_type", pages.GetBuildType()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cname", pages.GetCNAME()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("custom_404", pages.GetCustom404()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("html_url", pages.GetHTMLURL()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("build_status", pages.GetStatus()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("api_url", pages.GetURL()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("public", pages.GetPublic()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("https_enforced", pages.GetHTTPSEnforced()); err != nil {
		return diag.FromErr(err)
	}
	// Set source only for legacy build type
	if pages.GetBuildType() == "legacy" && pages.GetSource() != nil {
		source := []map[string]any{
			{
				"branch": pages.GetSource().GetBranch(),
				"path":   pages.GetSource().GetPath(),
			},
		}
		if err := d.Set("source", source); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("source", []map[string]any{}); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}
