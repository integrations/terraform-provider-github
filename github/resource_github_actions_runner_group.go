package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v35/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceGithubActionsRunnerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsRunnerGroupCreate,
		Read:   resourceGithubActionsRunnerGroupRead,
		Update: resourceGithubActionsRunnerGroupUpdate,
		Delete: resourceGithubActionsRunnerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"allows_public_repositories": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inherited": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"runners_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"selected_repository_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:      schema.HashInt,
				Optional: true,
			},
			"selected_repositories_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"visibility": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"all", "selected", "private"}, false),
			},
		},
	}
}

func resourceGithubActionsRunnerGroupCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	name := d.Get("name").(string)
	visibility := d.Get("visibility").(string)
	selectedRepositories, hasSelectedRepositories := d.GetOk("selected_repository_ids")

	if visibility != "selected" && hasSelectedRepositories {
		return fmt.Errorf("Cannot use selected_repository_ids without visibility being set to selected")
	}

	selectedRepositoryIDs := []int64{}

	if hasSelectedRepositories {
		ids := selectedRepositories.(*schema.Set).List()

		for _, id := range ids {
			selectedRepositoryIDs = append(selectedRepositoryIDs, int64(id.(int)))
		}
	}

	// TODO: also get runners
	// runners := d.Get("runners").([]int64)
	ctx := context.Background()

	log.Printf("[DEBUG] Creating organization runner group: %s (%s)", name, orgName)
	runnerGroup, resp, err := client.Actions.CreateOrganizationRunnerGroup(ctx,
		orgName,
		github.CreateRunnerGroupRequest{
			Name:                  &name,
			Visibility:            &visibility,
			SelectedRepositoryIDs: selectedRepositoryIDs,
			// TODO
			// Runners:               runners,
		},
	)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(runnerGroup.GetID(), 10))
	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("allows_public_repositories", runnerGroup.GetAllowsPublicRepositories())
	d.Set("default", runnerGroup.GetDefault())
	d.Set("id", runnerGroup.GetID())
	d.Set("inherited", runnerGroup.GetInherited())
	d.Set("name", runnerGroup.GetName())
	d.Set("runners_url", runnerGroup.GetRunnersURL())
	d.Set("selected_repositories_url", runnerGroup.GetSelectedRepositoriesURL())
	d.Set("visibility", runnerGroup.GetVisibility())

	return resourceGithubActionsRunnerGroupRead(d, meta)
}

func resourceGithubActionsRunnerGroupRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	runnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	log.Printf("[DEBUG] Reading organization runner group: %s (%s)", d.Id(), orgName)
	runnerGroup, resp, err := client.Actions.GetOrganizationRunnerGroup(ctx, orgName, runnerGroupID)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing organization runner group %s/%s from state because it no longer exists in GitHub",
					orgName, d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("allows_public_repositories", runnerGroup.GetAllowsPublicRepositories())
	d.Set("default", runnerGroup.GetDefault())
	d.Set("id", runnerGroup.GetID())
	d.Set("inherited", runnerGroup.GetInherited())
	d.Set("name", runnerGroup.GetName())
	d.Set("runners_url", runnerGroup.GetRunnersURL())
	d.Set("selected_repositories_url", runnerGroup.GetSelectedRepositoriesURL())
	d.Set("visibility", runnerGroup.GetVisibility())

	return nil
}

func resourceGithubActionsRunnerGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	name := d.Get("name").(string)
	visibility := d.Get("visibility").(string)

	options := github.UpdateRunnerGroupRequest{
		Name:       &name,
		Visibility: &visibility,
	}

	runnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Updating organization runner group: %s (%s)", d.Id(), orgName)
	if _, _, err := client.Actions.UpdateOrganizationRunnerGroup(ctx, orgName, runnerGroupID, options); err != nil {
		return err
	}

	return resourceGithubActionsRunnerGroupRead(d, meta)
}

func resourceGithubActionsRunnerGroupDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	runnerGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Deleting organization runner group: %s (%s)", d.Id(), orgName)
	_, err = client.Actions.DeleteOrganizationRunnerGroup(ctx, orgName, runnerGroupID)
	return err
}
