package github

import (
	"log"

	"github.com/google/go-github/v32/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubRepositoryPage() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryPageCreate,
		Read:   resourceGithubRepositoryPageRead,
		Update: resourceGithubRepositoryPageUpdate,
		Delete: resourceGithubRepositoryPageDelete,

		// TODO: Figure out importing
		// Importer: &schema.ResourceImporter{
		// 	State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
		// 		parts := strings.Split(d.Id(), "/")
		// 		if len(parts) != 2 {
		// 			return nil, fmt.Errorf("Invalid ID specified. Supplied ID must be written as <repository>/<Page_id>")
		// 		}
		// 		d.Set("repository", parts[0])
		// 		d.SetId(parts[1])
		// 		return []*schema.ResourceData{d}, nil
		// 	},
		// },

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": {
							Type:     schema.TypeString,
							Required: true,
							Default:  "main",
							// TODO: Update to current restrictsion`
							// ValidateFunc: validation.StringInSlice([]string{
							// 	"master",
							// 	"gh-pages",
							// }, false),
						},
						"path": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "/",
							// TODO: update to current restrictions
							// ValidateFunc: validation.StringInSlice([]string{
							// 	"/",
							// 	"/docs",
							// }, false),
						},
					},
				},
			},
			"cname": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// TODO: Add attributes
			// "custom_404": {
			// 	Type:     schema.TypeBool,
			// 	Computed: true,
			// },
			// "html_url": {
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },
			// "status": {
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },
			// "url": {
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },

		},
	}
}

func resourceGithubRepositoryPageCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	ctx := meta.(*Owner).StopContext

	log.Printf("[DEBUG] Creating repository Page: %s/%s", owner, repoName)

	pages := expandPages(d.Get("pages").([]interface{}))
	if pages != nil {
		_, _, err := client.Repositories.EnablePages(ctx, owner, repoName, pages)
		if err != nil {
			return err
		}
	}

	// TODO: Set to something specific
	d.SetId("0")

	return resourceGithubRepositoryPageRead(d, meta)
}

func resourceGithubRepositoryPageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	ctx := meta.(*Owner).StopContext

	log.Printf("[DEBUG] Reading repository Page: %s", d.Id())

	pages, _, err := client.Repositories.GetPagesInfo(ctx, owner, repoName)
	if err != nil {
		return err
	}

	// TODO: Remove
	// if err := d.Set("pages", flattenPages(pages)); err != nil {
	// 	return fmt.Errorf("error setting pages: %w", err)
	// }

	// Page, resp, err := client.Pages.GetPage(ctx, PageID)
	// if err != nil {
	// 	if ghErr, ok := err.(*github.ErrorResponse); ok {
	// 		if ghErr.Response.StatusCode == http.StatusNotModified {
	// 			return nil
	// 		}
	// 		if ghErr.Response.StatusCode == http.StatusNotFound {
	// 			log.Printf("[WARN] Removing repository Page %s from state because it no longer exists in GitHub",
	// 				d.Id())
	// 			d.SetId("")
	// 			return nil
	// 		}
	// 	}
	// return err
	// }

	// TODO:
	// d.Set("etag", resp.Header.Get("ETag"))
	// d.Set("name", Page.GetName())
	// d.Set("body", Page.GetBody())
	// d.Set("url", fmt.Sprintf("https://github.com/%s/%s/Pages/%d",
	// 	owner, d.Get("repository"), Page.GetNumber()))

	return nil
}

func resourceGithubRepositoryPageUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*Owner).v3client
	//
	// name := d.Get("name").(string)
	// body := d.Get("body").(string)
	//
	// options := github.PageOptions{
	// 	Name: &name,
	// 	Body: &body,
	// }
	//
	// PageID, err := strconv.ParseInt(d.Id(), 10, 64)
	// if err != nil {
	// 	return unconvertibleIdErr(d.Id(), err)
	// }
	// ctx := context.WithValue(context.Background(), ctxId, d.Id())
	//
	// log.Printf("[DEBUG] Updating repository Page: %s", d.Id())
	// _, _, err = client.Pages.UpdatePage(ctx, PageID, &options)
	// if err != nil {
	// 	return err
	// }

	return resourceGithubRepositoryPageRead(d, meta)
}

func resourceGithubRepositoryPageDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	ctx := meta.(*Owner).StopContext

	_, err := client.Repositories.DisablePages(ctx, owner, repoName)
	if err != nil {
		return err
	}

	return nil

}

func expandPages(input []interface{}) *github.Pages {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	pages := input[0].(map[string]interface{})
	pagesSource := pages["source"].([]interface{})[0].(map[string]interface{})
	source := &github.PagesSource{
		Branch: github.String(pagesSource["branch"].(string)),
	}
	if v, ok := pagesSource["path"].(string); ok {
		// Github Pages API only accepts "/docs" in source.Path;
		// to set to the root directory "/", leave source.Path unset
		if v != "" && v != "/" {
			source.Path = github.String(v)
		}
	}
	return &github.Pages{Source: source}
}
