package github

import (
	"context"
	"log"
	"strconv"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGithubUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubUserRead,

		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"username"},
				ValidateFunc:  validateNumericIDFunc,
			},
			"username": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"user_id"},
			},
			"minimal": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
	client := meta.(*Organization).client
	ctx := context.Background()

	var user *github.User
	var err error

	if res, ok := d.GetOk("username"); ok {
		username := res.(string)
		log.Printf("[INFO] Refreshing GitHub User: %s", username)
		user, _, err = client.Users.Get(ctx, username)
		if err != nil {
			return err
		}
	}
	if res, ok := d.GetOk("user_id"); ok {
		userID, err := strconv.ParseInt(res.(string), 10, 64)
		if err != nil {
			return err
		}
		log.Printf("[INFO] Refreshing GitHub User: %d", userID)
		user, _, err = client.Users.GetByID(ctx, userID)
		if err != nil {
			return err
		}
	}

	minimal := d.Get("minimal").(bool)

	gpgKeys := []string{}
	sshKeys := []string{}

	if !minimal {
		gpg, _, err := client.Users.ListGPGKeys(ctx, user.GetLogin(), nil)
		if err != nil {
			return err
		}
		ssh, _, err := client.Users.ListKeys(ctx, user.GetLogin(), nil)
		if err != nil {
			return err
		}

		for _, v := range gpg {
			gpgKeys = append(gpgKeys, v.GetPublicKey())
		}

		for _, v := range ssh {
			sshKeys = append(sshKeys, v.GetKey())
		}
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
	d.Set("public_repos", user.GetPublicRepos())
	d.Set("public_gists", user.GetPublicGists())
	d.Set("followers", user.GetFollowers())
	d.Set("following", user.GetFollowing())
	d.Set("created_at", user.GetCreatedAt())
	d.Set("updated_at", user.GetUpdatedAt())
	d.Set("gpg_keys", gpgKeys)
	d.Set("ssh_keys", sshKeys)

	return nil
}
