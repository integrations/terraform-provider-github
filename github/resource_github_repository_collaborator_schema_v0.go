package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubRepositoryCollaboratorV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"username": {
				Type: schema.TypeString,
			},
			"repository": {
				Type: schema.TypeString,
			},
			"permission": {
				Type: schema.TypeString,
			},
			"invitation_id": {
				Type: schema.TypeString,
			},
		},
	}
}

func resourceGithubRepositoryCollaboratorStateUpgradeV0(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	// Schema V1 changes from V0:
	//
	// 'username' was removed (delete from rawState)
	// 'user_id' was added as user-provided and required (must be populated in rawState)
	// 'id' was changed from 'repo:username' format to 'repo:user_id' format (change format in rawState)

	user, _, err := meta.(*Organization).client.Users.Get(context.Background(), rawState["username"].(string))
	if err != nil {
		return nil, err
	}

	delete(rawState, "username")
	userIDString := strconv.FormatInt(*user.ID, 10)
	rawState["user_id"] = userIDString
	rawState["id"] = buildTwoPartID(rawState["repository"].(string), userIDString)

	return rawState, nil
}
