package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/shurcooL/githubv4"
)

const (
	USER_EMAIL         = "email"
	USER_IS_SITE_ADMIN = "is_site_admin"
	USER_LOGIN         = "login"
	USER_NAME          = "name"
)

type User struct {
	Email       githubv4.String
	ID          githubv4.ID
	IsSiteAdmin githubv4.Boolean
	Login       githubv4.String
	Name        githubv4.String
}

func dataSourceGithubUser() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			// Input
			USER_LOGIN: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			// Computed
			USER_EMAIL: {
				Type:     schema.TypeString,
				Computed: true,
			},
			USER_IS_SITE_ADMIN: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			USER_NAME: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},

		Read: dataSourceGithubUserRead,
	}
}

func dataSourceGithubUserRead(d *schema.ResourceData, meta interface{}) error {
	var query struct {
		User User `graphql:"user(login: $login)"`
	}
	variables := map[string]interface{}{
		"login": githubv4.String(d.Get(USER_LOGIN).(string)),
	}

	ctx := context.Background()
	client := meta.(*Organization).v4client
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return err
	}

	err = d.Set(USER_EMAIL, query.User.Email)
	if err != nil {
		return err
	}

	err = d.Set(USER_IS_SITE_ADMIN, query.User.IsSiteAdmin)
	if err != nil {
		return err
	}

	err = d.Set(USER_NAME, query.User.Name)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s", query.User.ID))

	return nil
}
