package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubTeamMembershipV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type: schema.TypeString,
			},
			"username": {
				Type: schema.TypeString,
			},
			"role": {
				Type: schema.TypeString,
			},
			"etag": {
				Type: schema.TypeString,
			},
		},
	}
}

func resourceGithubTeamMembershipStateUpgradeV0(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	// Schema V1 changes from V0:
	//
	// 'username' was removed (delete from rawState)
	// 'user_id' was added as user-provided and required (must be populated in rawState)
	// 'id' was changed from 'team_id:username' format to 'team_id:user_id' format (change format in rawState)

	user, _, err := meta.(*Organization).client.Users.Get(context.Background(), rawState["username"].(string))
	if err != nil {
		return nil, err
	}

	delete(rawState, "username")
	teamIDString := rawState["team_id"].(string)
	userIDString := strconv.FormatInt(*user.ID, 10)
	rawState["user_id"] = userIDString
	rawState["id"] = buildTwoPartID(teamIDString, userIDString)

	return rawState, nil
}
