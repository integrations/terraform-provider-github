package github

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v29/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubTeamSyncGroupMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubTeamSyncGroupMappingCreate,
		Read:   resourceGithubTeamSyncGroupMappingRead,
		Update: resourceGithubTeamSyncGroupMappingUpdate,
		Delete: resourceGithubTeamSyncGroupMappingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"team_slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
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

func resourceGithubTeamSyncGroupMappingCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client
	ctx := context.Background()
	orgName := meta.(*Organization).name
	slug := d.Get("team_slug").(string)

	team, err := getGithubTeamBySlug(ctx, client, orgName, slug)
	if err != nil {
		return err
	}

	idpGroupList, err := expandTeamSyncGroups(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Creating team-sync group mapping (Team: %s)", team.GetName())
	_, resp, err := client.Teams.CreateOrUpdateIDPGroupConnections(ctx, string(team.GetID()), *idpGroupList)
	if err != nil {
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.SetId(fmt.Sprintf("teams/%s/team-sync/group-mappings", slug))

	return resourceGithubTeamRead(d, meta)
}

func resourceGithubTeamSyncGroupMappingRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client
	orgName := meta.(*Organization).name
	slug := d.Get("team_slug").(string)

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	team, err := getGithubTeamBySlug(ctx, client, orgName, slug)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Reading team-sync group mapping (Team: %s)", team.GetName())
	idpGroupList, _, err := client.Teams.ListIDPGroupsForTeam(ctx, string(team.GetID()))
	if err != nil {
		return err
	}

	groups, err := flattenGithubIDPGroupList(idpGroupList)
	if err != nil {
		return err
	}

	d.Set("groups", groups)

	return nil
}

func resourceGithubTeamSyncGroupMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client
	orgName := meta.(*Organization).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	slug := d.Get("team_slug").(string)

	team, err := getGithubTeamBySlug(ctx, client, orgName, slug)
	if err != nil {
		return err
	}

	idpGroupList, err := expandTeamSyncGroups(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Updating team-sync group mapping (Team: %s)", team.GetName())
	_, _, err = client.Teams.CreateOrUpdateIDPGroupConnections(ctx, string(team.GetID()), *idpGroupList)
	if err != nil {
		return err
	}

	return resourceGithubTeamRead(d, meta)
}

func resourceGithubTeamSyncGroupMappingDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client
	orgName := meta.(*Organization).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	slug := d.Get("team_slug").(string)

	team, err := getGithubTeamBySlug(ctx, client, orgName, slug)
	if err != nil {
		return err
	}

	groups := make([]*github.IDPGroup, 0)
	emptyGroupList := github.IDPGroupList{Groups: groups}

	log.Printf("[DEBUG] Deleting team-sync group mapping (Team: %s)", team.GetName())
	_, _, err = client.Teams.CreateOrUpdateIDPGroupConnections(ctx, string(team.GetID()), emptyGroupList)

	return err
}

func flattenGithubIDPGroupList(idpGroupList *github.IDPGroupList) ([]interface{}, error) {
	if idpGroupList == nil {
		return make([]interface{}, 0), nil
	}
	results := make([]interface{}, 0)
	for _, group := range idpGroupList.Groups {
		result := make(map[string]interface{})
		result["group_id"] = group.GetGroupID()
		result["group_name"] = group.GetGroupName()
		result["group_description"] = group.GetGroupDescription()
		results = append(results, result)
	}

	return results, nil
}

func expandTeamSyncGroups(d *schema.ResourceData) (*github.IDPGroupList, error) {
	if v, ok := d.GetOk("groups"); ok {
		vL := v.([]interface{})
		idpGroupList := new(github.IDPGroupList)
		groups := make([]*github.IDPGroup, 0)
		for _, v := range vL {
			m := v.(map[string]interface{})
			group := new(github.IDPGroup)
			group.GroupID = m["group_id"].(*string)
			group.GroupName = m["group_name"].(*string)
			group.GroupDescription = m["group_description"].(*string)
			groups = append(groups, group)
		}
		idpGroupList.Groups = groups
		return idpGroupList, nil
	}
	return nil, nil
}
