package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v77/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubProjectItem() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubProjectItemCreate,
		Read:   resourceGithubProjectItemRead,
		Update: resourceGithubProjectItemUpdate,
		Delete: resourceGithubProjectItemDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubProjectItemImport,
		},
		Schema: map[string]*schema.Schema{
			"project_number": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The number of the project (Projects V2).",
			},
			"content_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the issue or pull request to add to the project.",
			},
			"content_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v != "Issue" && v != "PullRequest" {
						errs = append(errs, fmt.Errorf("%q must be either 'Issue' or 'PullRequest', got: %s", key, v))
					}
					return
				},
				Description: "Must be either 'Issue' or 'PullRequest'.",
			},
			"archived": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the item is archived.",
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The node ID of the project item.",
			},
			"item_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the project item.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubProjectItemCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	projectNumber := d.Get("project_number").(int)
	contentID := int64(d.Get("content_id").(int))
	contentType := d.Get("content_type").(string)

	options := &github.AddProjectItemOptions{
		Type: contentType,
		ID:   contentID,
	}

	ctx := context.Background()
	log.Printf("[DEBUG] Adding %s %d to project %d", contentType, contentID, projectNumber)

	item, _, err := client.Projects.AddOrganizationProjectItem(ctx, orgName, projectNumber, options)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(item.GetID(), 10))
	if err = d.Set("item_id", item.GetID()); err != nil {
		return err
	}
	if err = d.Set("node_id", item.GetNodeID()); err != nil {
		return err
	}

	// If archived is set to true, update the item
	if d.Get("archived").(bool) {
		updateOpts := &github.UpdateProjectItemOptions{
			Archived: github.Ptr(true),
		}
		_, _, err = client.Projects.UpdateOrganizationProjectItem(ctx, orgName, projectNumber, item.GetID(), updateOpts)
		if err != nil {
			return err
		}
	}

	return resourceGithubProjectItemRead(d, meta)
}

func resourceGithubProjectItemRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	projectNumber := d.Get("project_number").(int)

	itemID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	log.Printf("[DEBUG] Reading project item: %d", itemID)
	item, resp, err := client.Projects.GetOrganizationProjectItem(ctx, orgName, projectNumber, itemID, nil)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing project item %s from state because it no longer exists in GitHub", d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return err
	}
	if err = d.Set("item_id", item.GetID()); err != nil {
		return err
	}
	if err = d.Set("node_id", item.GetNodeID()); err != nil {
		return err
	}
	if err = d.Set("content_type", item.GetContentType()); err != nil {
		return err
	}
	archived := false
	if item.ArchivedAt != nil {
		archived = true
	}
	if err = d.Set("archived", archived); err != nil {
		return err
	}

	return nil
}

func resourceGithubProjectItemUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	projectNumber := d.Get("project_number").(int)

	itemID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	log.Printf("[DEBUG] Updating project item: %d", itemID)

	// Only archived status can be updated for project items
	if d.HasChange("archived") {
		archived := d.Get("archived").(bool)
		options := &github.UpdateProjectItemOptions{
			Archived: &archived,
		}

		ctx := context.WithValue(context.Background(), ctxId, d.Id())
		_, _, err := client.Projects.UpdateOrganizationProjectItem(ctx, orgName, projectNumber, itemID, options)
		if err != nil {
			return err
		}
	}

	return resourceGithubProjectItemRead(d, meta)
}

func resourceGithubProjectItemDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	projectNumber := d.Get("project_number").(int)

	itemID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Deleting project item: %d", itemID)
	_, err = client.Projects.DeleteOrganizationProjectItem(ctx, orgName, projectNumber, itemID)
	if err != nil {
		return err
	}

	return nil
}

func resourceGithubProjectItemImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// For project items, we need: org/project_number/item_id
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid ID format: expected 'org/project_number/item_id', got '%s'", d.Id())
	}

	projectNumber, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid project number: %v", err)
	}

	itemIDStr := parts[2]
	itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid item ID: %v", err)
	}

	// Set the computed ID to just the item ID
	d.SetId(itemIDStr)
	if err = d.Set("project_number", projectNumber); err != nil {
		return []*schema.ResourceData{d}, err
	}

	log.Printf("[DEBUG] Imported project item %d from project %d", itemID, projectNumber)
	return []*schema.ResourceData{d}, nil
}
