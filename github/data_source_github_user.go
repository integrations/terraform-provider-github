package github

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGithubUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubUserRead,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"login": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"avatar_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gravatar_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_admin": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"company": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"blog": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bio": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gpg_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ssh_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"public_repos": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"public_gists": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"followers": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"following": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubUserRead(d *schema.ResourceData, meta interface{}) error {
	username := d.Get("username").(string)
	log.Printf("[INFO] Refreshing GitHub User: %s", username)

	client := meta.(*Organization).v3client
	ctx := context.Background()

	user, _, err := client.Users.Get(ctx, username)
	if err != nil {
		return err
	}

	gpg, _, err := client.Users.ListGPGKeys(ctx, username, nil)
	if err != nil {
		return err
	}
	ssh, _, err := client.Users.ListKeys(ctx, username, nil)
	if err != nil {
		return err
	}

	gpgKeys := []string{}
	for _, v := range gpg {
		gpgKeys = append(gpgKeys, v.GetPublicKey())
	}

	sshKeys := []string{}
	for _, v := range ssh {
		sshKeys = append(sshKeys, v.GetKey())
	}

	d.SetId(strconv.FormatInt(user.GetID(), 10))
	d.Set("login", user.GetLogin())
	d.Set("avatar_url", user.GetAvatarURL())
	d.Set("gravatar_id", user.GetGravatarID())
	d.Set("site_admin", user.GetSiteAdmin())
	d.Set("company", user.GetCompany())
	d.Set("blog", user.GetBlog())
	d.Set("location", user.GetLocation())
	d.Set("name", user.GetName())
	d.Set("email", user.GetEmail())
	d.Set("bio", user.GetBio())
	d.Set("gpg_keys", gpgKeys)
	d.Set("ssh_keys", sshKeys)
	d.Set("public_repos", user.GetPublicRepos())
	d.Set("public_gists", user.GetPublicGists())
	d.Set("followers", user.GetFollowers())
	d.Set("following", user.GetFollowing())
	d.Set("created_at", user.GetCreatedAt())
	d.Set("updated_at", user.GetUpdatedAt())

	return nil
}
