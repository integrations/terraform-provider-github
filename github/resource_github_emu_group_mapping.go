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
	//name := groupMap["group_name"]
	//desc := groupMap["group_description"]

	//eg, resp, err := client.Teams.GetExternalGroup(ctx, orgName, groupId)
	//fmt.Printf("resp: %v", resp)

	eg := &github.ExternalGroup{
		GroupID: &groupId,
	}

	group, resp, err := client.Teams.UpdateConnectedExternalGroup(ctx, orgName, teamSlug, eg)
	fmt.Printf("group: %v, resp: %v", group, resp)

	// todo: set terraform state
	return nil
}

func resourceGithubEMUGroupMappingRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	managedTeam := d.Get("managed_team").(string)
	githubTeam := d.Get("github_team").(string)

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	opts := &github.ListExternalGroupsOptions{}

	filteredTeams := make([]*github.ExternalGroup, 0)
	teams, resp, err := client.Teams.ListExternalGroups(ctx, orgName, opts)
	fmt.Printf("[DEBUG] Response code: %v", resp.StatusCode)
	if err != nil {
		// might need to do something here to ignore expected errors
		return err
	}
	// example groupID: 28836 name: terraform-emu-test-group
	// gonna need to do another lookup here maybe with groupID to do a match?
	//client.Teams.GetTeamByID(ctx, )
	for _, team := range teams.Groups {
		if *team.GroupName == managedTeam {
			filteredTeams = append(filteredTeams, team)
		}
	}
	fmt.Printf("[DEBUG]: GitHub team: %v", githubTeam)
	// do the d.set stuff to set the resource information
	return nil
}

func resourceGithubEMUGroupMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGithubEMUGroupMappingDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
