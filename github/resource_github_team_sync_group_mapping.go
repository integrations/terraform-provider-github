package github

import (
	"context"

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
				Optional: true,
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
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"org_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"team_id": {
				Type:     schema.TypeInt,
				Computed: true,
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
	orgName := meta.(*Organization).name
	ctx := context.Background()

	slug := d.Get("team_slug").(string)
	team, err := getGithubTeamBySlug(ctx, client, orgName, slug)
	if err != nil {
		return err
	}

	if g, ok := d.GetOk("groups"); ok {
		var idpGL github.IDPGroupList
		gL := g.([]map[string]string)
		if len(gL) > 0 {
			idpGL.Groups = createIDPGroups(gL)
		}
		_, resp, err := client.Teams.CreateOrUpdateIDPGroupConnections(ctx, string(team.GetID()), idpGL)
		if err != nil {
			return err
		}
		d.Set("etag", resp.Header.Get("ETag"))
	}

	d.SetId("github-team-sync-group-mappings")
	d.Set("org_id", meta.(*Organization).id)
	d.Set("org_name", orgName)
	d.Set("team_id", team.GetID())
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
			"group_id":    g.GetGroupID(),
			"group_name":  g.GetGroupName(),
			"description": g.GetGroupDescription(),
		}
		groups = append(groups, group)
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
	teamID := d.Get("team_id").(string)
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	if g, ok := d.GetOk("groups"); ok {
		var idpGL github.IDPGroupList
		gL := g.([]map[string]string)
		if len(gL) > 0 {
			idpGL.Groups = createIDPGroups(gL)
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

func createIDPGroups(gL []map[string]string) []*github.IDPGroup {
	var idpGroups []*github.IDPGroup
	for _, group := range gL {
		id := group["group_id"]
		name := group["group_name"]
		description := group["description"]
		idpGroup := github.IDPGroup{&id, &name, &description}
		idpGroups = append(idpGroups, &idpGroup)
	}
	return idpGroups
}
