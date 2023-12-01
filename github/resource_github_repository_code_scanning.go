package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v55/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceGithubRepositoryCodeScanning() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryCodeScanningCreate,
		Read:   resourceGithubRepositoryCodeScanningRead,
		Update: resourceGithubRepositoryCodeScanningUpdate,
		Delete: resourceGithubRepositoryCodeScanningDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		// editing repository collaborators are not supported by github api so forcing new on any changes
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GitHub repository",
			},
			"owner": {
				Type:     schema.TypeString,
				Required: true,
			},
			"languages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"query_suite": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"default",
					"extended",
				}, false),
			},
			"state": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"configured",
					"not-configured",
				}, false),
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubRepositoryCodeScanningCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	createUpdateOpts := createUpdateCodeScanning(d, meta)
	ctx := context.Background()

	_, _, err := client.CodeScanning.UpdateDefaultSetupConfiguration(ctx,
		owner,
		repoName,
		&createUpdateOpts,
	)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(owner, repoName))

	return resourceGithubRepositoryCodeScanningRead(d, meta)
}

func resourceGithubRepositoryCodeScanningDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner, repoName, err := parseTwoPartID(d.Id(), "owner", "repository")
	if err != nil {
		return err
	}

	createUpdateOpts := createUpdateCodeScanning(d, meta)
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	_, _, err = client.CodeScanning.UpdateDefaultSetupConfiguration(ctx, owner, repoName, &createUpdateOpts)
	return err
}

func resourceGithubRepositoryCodeScanningRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner, repoName, err := parseTwoPartID(d.Id(), "owner", "repository")
	if err != nil {
		return err
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	config, _, err := client.CodeScanning.GetDefaultSetupConfiguration(ctx, owner, repoName)
	if err != nil {
		return err
	}

	d.Set("repository", repoName)
	d.Set("owner", owner)
	d.Set("state", config.State)
	d.Set("query_suite", config.QuerySuite)
	d.Set("languages", config.Languages)
	d.Set("updated_at", config.UpdatedAt.String())

	return nil
}

func resourceGithubRepositoryCodeScanningUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner, repoName, err := parseTwoPartID(d.Id(), "owner", "repository")
	if err != nil {
		return err
	}

	createUpdateOpts := createUpdateCodeScanning(d, meta)
	ctx := context.Background()

	_, _, err = client.CodeScanning.UpdateDefaultSetupConfiguration(ctx,
		owner,
		repoName,
		&createUpdateOpts,
	)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(owner, repoName))

	return resourceGithubRepositoryCodeScanningRead(d, meta)
}

func createUpdateCodeScanning(d *schema.ResourceData, meta interface{}) github.UpdateDefaultSetupConfigurationOptions {
	data := github.UpdateDefaultSetupConfigurationOptions{}

	var languages []string
	if v, ok := d.GetOk("languages"); ok {
		for _, lang := range v.([]interface{}) {
			languages = append(languages, fmt.Sprintf("%s", lang))
		}

		data.Languages = languages
	}

	if v, ok := d.GetOk("query_suite"); ok {
		querySuite := v.(string)
		data.QuerySuite = &querySuite
	}

	data.State = d.Get("state").(string)

	return data
}
