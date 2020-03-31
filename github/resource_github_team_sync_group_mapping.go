package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v29/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"retrieve_by": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"id",
					"slug",
				}, false),
			},
			"org_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"team_slug": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"team_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_description": {
							Type:     schema.TypeString,
							Optional: true,
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

	retrieveBy := d.Get("retrieve_by").(string)
	var team *github.Team

	if retrieveBy == "id" {
		orgID := meta.(*Organization).id
		teamID, err := strconv.ParseInt(d.Get("team_id").(string), 10, 64)
		if err != nil {
			return err
		}
		team, _, err = client.Teams.GetTeamByID(ctx, orgID, teamID)
	} else {
		orgName := meta.(*Organization).name
		slug := d.Get("team_slug").(string)
		team, err = getGithubTeamBySlug(ctx, client, orgName, slug)
	}

	if err != nil {
		return err
	}

	if g, ok := d.GetOk("group"); ok {
		var idpGL github.IDPGroupList
		gL := g.([]interface{})
		if len(gL) > 0 {
			idpGL = github.IDPGroupList{Groups: createIDPGroups(gL)}
		}
		_, resp, err := client.Teams.CreateOrUpdateIDPGroupConnections(ctx, string(team.GetID()), idpGL)
		if err != nil {
			return err
		}
		d.Set("etag", resp.Header.Get("ETag"))
	}

	d.SetId("github-team-sync-group-mappings")
	d.Set("org_name", meta.(*Organization).name)
	d.Set("org_id", meta.(*Organization).id)
	d.Set("team_id", string(team.GetID()))
	return resourceGithubTeamRead(d, meta)
}

func resourceGithubTeamSyncGroupMappingRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client
	teamID := d.Get("team_id").(string)

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	groupList, _, err := client.Teams.ListIDPGroupsForTeam(ctx, teamID)
	if err != nil {
		return err
	}

	var groups []map[string]string
	for _, g := range groupList.Groups {
		group := map[string]string{
			"group_id":          g.GetGroupID(),
			"group_name":        g.GetGroupName(),
			"group_description": g.GetGroupDescription(),
		}
		groups = append(groups, group)
	}

	d.Set("group", groups)

	return nil
}

func resourceGithubTeamSyncGroupMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client
	teamID := d.Get("team_id").(string)
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	if g, ok := d.GetOk("group"); ok {
		var idpGL github.IDPGroupList
		gL := g.([]interface{})
		if len(gL) > 0 {
			idpGL = github.IDPGroupList{Groups: createIDPGroups(gL)}
		}
		_, _, err := client.Teams.CreateOrUpdateIDPGroupConnections(ctx, teamID, idpGL)
		if err != nil {
			return err
		}
	}

	return resourceGithubTeamRead(d, meta)
}

func resourceGithubTeamSyncGroupMappingDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client
	teamID := d.Get("team_id").(string)
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	_, _, err = client.Teams.CreateOrUpdateIDPGroupConnections(ctx, teamID, github.IDPGroupList{})

	return err
}

func createIDPGroups(gL []interface{}) []*github.IDPGroup {
	var idpGroups []*github.IDPGroup
	for _, group := range gL {
		g := group.(map[string]interface{})
		id := g["group_id"].(string)
		name := g["group_name"].(string)
		description := g["group_description"].(string)
		idpGroup := github.IDPGroup{GroupID: &id, GroupName: &name, GroupDescription: &description}
		idpGroups = append(idpGroups, &idpGroup)
	}
	return idpGroups
}
