package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRef() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRefRead,

		Schema: map[string]*schema.Schema{
			"ref": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubRefRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner, ok := d.Get("owner").(string)
	if !ok {
		owner = meta.name
	}
	repoName := d.Get("repository").(string)
	ref := d.Get("ref").(string)

	refData, resp, err := client.Git.GetRef(ctx, owner, repoName, ref)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Debug(ctx, "Missing GitHub ref", map[string]any{"owner": owner, "repoName": repoName, "ref": ref})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(repoName, ref))
	if meta.storeEtag {
		if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
			return diag.FromErr(err)
		}
	}
	err = d.Set("sha", *refData.Object.SHA)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
