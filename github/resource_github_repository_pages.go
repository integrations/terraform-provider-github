package github

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryPages() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages GitHub Pages for a repository.",
		CreateContext: resourceGithubRepositoryPagesCreate,
		ReadContext:   resourceGithubRepositoryPagesRead,
		UpdateContext: resourceGithubRepositoryPagesUpdate,
		DeleteContext: resourceGithubRepositoryPagesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryPagesImport,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The repository name to configure GitHub Pages for.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the repository to configure GitHub Pages for.",
			},
			// TODO: Uncomment this when we are ready to support owner fields properly. https://github.com/integrations/terraform-provider-github/pull/3166#discussion_r2816053082
			// "owner": {
			// 	Type:        schema.TypeString,
			// 	Required:    true,
			// 	ForceNew:    true,
			// 	Description: "The owner of the repository to configure GitHub Pages for.",
			// },
			"source": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The source branch and directory for the rendered Pages site.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The repository branch used to publish the site's source files. (i.e. 'main' or 'gh-pages')",
						},
						"path": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "/",
							Description: "The repository directory from which the site publishes (Default: '/')",
						},
					},
				},
			},
			"build_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "legacy",
				Description:      "The type of GitHub Pages site to build. Can be 'legacy' or 'workflow'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"legacy", "workflow"}, false)),
			},
			"cname": {
				Type:        schema.TypeString,
				Optional:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.",
			},
			"https_enforced": {
				Type:         schema.TypeBool,
				Optional:     true,
				Computed:     true,
				RequiredWith: []string{"cname"},
				Description:  "Whether the rendered GitHub Pages site will only be served over HTTPS. Requires 'cname' to be set.",
			},
		},
		CustomizeDiff: customdiff.All(resourceGithubRepositoryPagesDiff, diffRepository),
	}
}

func resourceGithubRepositoryPagesCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client

	owner := meta.name // TODO: Add owner support // d.Get("owner").(string)
	repoName := d.Get("repository").(string)

	pagesReq := expandPagesForCreate(d)
	pages, _, err := client.Repositories.EnablePages(ctx, owner, repoName, pagesReq)
	if err != nil {
		return diag.FromErr(err)
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(int(repo.GetID())))

	if err = d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}

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

	// Determine if we need to update the page with CNAME or public flag
	shouldUpdatePage := false
	update := &github.PagesUpdate{}
	cname, cnameExists := d.GetOk("cname")
	if cnameExists && cname.(string) != "" {
		shouldUpdatePage = true
		update.CNAME = github.Ptr(cname.(string))
	}
	public, publicExists := d.GetOkExists("public") // nolint:staticcheck // SA1019: There is no better alternative for checking if boolean value is set
	if publicExists && public != nil {
		shouldUpdatePage = true
		update.Public = github.Ptr(public.(bool))
	} else {
		if err := d.Set("public", pages.GetPublic()); err != nil {
			return diag.FromErr(err)
		}
	}
	httpsEnforced, httpsEnforcedExists := d.GetOkExists("https_enforced") // nolint:staticcheck // SA1019: There is no better alternative for checking if boolean value is set
	if httpsEnforcedExists && httpsEnforced != nil {
		shouldUpdatePage = true
		update.HTTPSEnforced = github.Ptr(httpsEnforced.(bool))
	} else {
		if err := d.Set("https_enforced", pages.GetHTTPSEnforced()); err != nil {
			return diag.FromErr(err)
		}
	}

	if shouldUpdatePage {
		_, err = client.Repositories.UpdatePages(ctx, owner, repoName, update)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubRepositoryPagesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client

	owner := meta.name // TODO: Add owner support // d.Get("owner").(string)
	repoName := d.Get("repository").(string)

	pages, resp, err := client.Repositories.GetPagesInfo(ctx, owner, repoName)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		return diag.Errorf("error reading repository pages: %s", err.Error())
	}

	if err := d.Set("build_type", pages.GetBuildType()); err != nil {
		return diag.FromErr(err)
	}
	upstreamCname := pages.GetCNAME()
	if upstreamCname != "" {
		if err := d.Set("cname", upstreamCname); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("cname", nil); err != nil {
			return diag.FromErr(err)
		}
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
		if err := d.Set("source", nil); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubRepositoryPagesUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client

	owner := meta.name // TODO: Add owner support // d.Get("owner").(string)
	repoName := d.Get("repository").(string)

	update := &github.PagesUpdate{}

	if d.HasChange("cname") {
		cname := d.Get("cname").(string)
		if cname != "" {
			update.CNAME = github.Ptr(cname)
		}
	}

	if d.HasChange("public") {
		public, ok := d.Get("public").(bool)
		if ok {
			update.Public = github.Ptr(public)
		}
	}

	if d.HasChange("https_enforced") {
		httpsEnforced, ok := d.Get("https_enforced").(bool)
		if ok {
			update.HTTPSEnforced = github.Ptr(httpsEnforced)
		}
	}

	if d.HasChange("build_type") {
		buildType := d.Get("build_type").(string)
		update.BuildType = github.Ptr(buildType)
	}

	if d.HasChange("source") || d.HasChange("build_type") {
		buildType := d.Get("build_type").(string)
		if buildType == "legacy" {
			if source, ok := d.GetOk("source"); ok {
				sourceList := source.([]any)
				if len(sourceList) > 0 {
					sourceMap := sourceList[0].(map[string]any)
					branch := sourceMap["branch"].(string)
					path := sourceMap["path"].(string)
					update.Source = &github.PagesSource{
						Branch: &branch,
						Path:   &path,
					}
				}
			}
		}
	}

	_, err := client.Repositories.UpdatePages(ctx, owner, repoName, update)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryPagesDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client

	owner := meta.name // TODO: Add owner support // d.Get("owner").(string)
	repoName := d.Get("repository").(string)

	_, err := client.Repositories.DisablePages(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(handleArchivedRepoDelete(err, "repository pages", d.Id(), owner, repoName))
	}

	return nil
}

func resourceGithubRepositoryPagesImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	repoName := d.Id()
	if strings.Contains(repoName, " ") {
		return nil, fmt.Errorf("invalid ID specified: supplied ID must be the slug of the repository name")
	}
	// if err := d.Set("owner", owner); err != nil { // TODO: Add owner support
	// 	return nil, err
	// }
	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}

	meta := m.(*Owner)
	owner := meta.name
	client := meta.v3client

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}
	if err = d.Set("repository_id", int(repo.GetID())); err != nil {
		return nil, err
	}

	d.SetId(strconv.Itoa(int(repo.GetID())))

	return []*schema.ResourceData{d}, nil
}

func resourceGithubRepositoryPagesDiff(ctx context.Context, d *schema.ResourceDiff, _ any) error {
	if d.Id() == "" {
		return nil
	}

	buildType := d.Get("build_type").(string)
	_, ok := d.GetOk("source")

	if buildType == "workflow" && ok {
		return fmt.Errorf("'source' is not supported for workflow build type")
	}
	if buildType == "legacy" && !ok {
		return fmt.Errorf("'source' is required for legacy build type")
	}

	return nil
}

func expandPagesForCreate(d *schema.ResourceData) *github.Pages {
	pages := &github.Pages{}

	buildType := d.Get("build_type").(string)
	pages.BuildType = github.Ptr(buildType)

	if buildType == "legacy" {
		if source, ok := d.GetOk("source"); ok {
			sourceList := source.([]any)
			if len(sourceList) > 0 {
				sourceMap := sourceList[0].(map[string]any)
				branch := sourceMap["branch"].(string)
				pagesSource := &github.PagesSource{
					Branch: github.Ptr(branch),
				}
				if path, ok := sourceMap["path"].(string); ok && path != "" && path != "/" {
					pagesSource.Path = github.Ptr(path)
				}
				pages.Source = pagesSource
			}
		}
		// Default to main branch if no source specified
		if pages.Source == nil {
			pages.Source = &github.PagesSource{
				Branch: github.Ptr("main"),
			}
		}
	}

	return pages
}
