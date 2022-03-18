package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v42/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubEMUGroupMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubEMUGroupMappingCreate,
		Read:   resourceGithubEMUGroupMappingRead,
		Update: resourceGithubEMUGroupMappingUpdate,
		Delete: resourceGithubEMUGroupMappingDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				fmt.Printf("resource data: %v, meta: %v", d, meta)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"team_slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"group_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"group_description": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubEMUGroupMappingCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	teamSlug := d.Get("team_slug").(string)
	groupMap := d.Get("group").(map[string]interface{})
	id := groupMap["group_id"].(string)

	groupId, _ := strconv.ParseInt(id, 10, 64)

	eg := &github.ExternalGroup{
		GroupID: &groupId,
	}

	group, resp, err := client.Teams.UpdateConnectedExternalGroup(ctx, orgName, teamSlug, eg)
	fmt.Printf("group: %v, resp: %v", group, resp)

	d.SetId(fmt.Sprintf("organizations/%s/team/%s/external-groups", orgName, teamSlug))
	return resourceGithubEMUGroupMappingRead(d, meta)
}

func resourceGithubEMUGroupMappingRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	//teamSlug := d.Get("team_slug").(string)
	groupMap := d.Get("group").(map[string]interface{})
	id := groupMap["group_id"].(string)
	groupId, _ := strconv.ParseInt(id, 10, 64)

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	group, resp, err := client.Teams.GetExternalGroup(ctx, orgName, groupId)
	if err != nil {
		// might need to do something here to ignore expected errors
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("group", group)
	return nil
}

func resourceGithubEMUGroupMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGithubEMUGroupMappingDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	teamSlug := d.Get("team_slug").(string)

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	resp, err := client.Teams.RemoveConnectedExternalGroup(ctx, orgName, teamSlug)
	fmt.Printf("resp: %v", resp)
	if err != nil {
		return err
	}
	return nil
}
