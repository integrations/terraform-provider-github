package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubMembershipV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
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

func resourceGithubMembershipStateUpgradeV0(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	// Schema V1 changes from V0:
	//
	// 'username' was removed (delete from rawState)
	// 'user_id' was added as user-provided and required (must be populated in rawState)
	// 'id' was changed from 'orgname:username' format to 'orgname:user_id' format (change format in rawState)

	orgName, err := getOrganization(meta)
	if err != nil {
		return nil, err
	}

	user, _, err := meta.(*Organization).client.Users.Get(context.Background(), rawState["username"].(string))
	if err != nil {
		return nil, err
	}

	delete(rawState, "username")
	userIDString := strconv.FormatInt(*user.ID, 10)
	rawState["user_id"] = userIDString
	rawState["id"] = buildTwoPartID(orgName, userIDString)

	return rawState, nil
}
