package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubUserRead,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username. Use an empty string to retrieve information about the currently authenticated user.",
			},
			"login": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user's login.",
			},
			"avatar_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user's avatar URL.",
			},
			"gravatar_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user's gravatar ID.",
			},
			"site_admin": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the user is a GitHub admin.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user's full name.",
			},
			"company": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user's company name.",
			},
			"blog": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user's blog location.",
			},
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user's location.",
			},
			"email": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user's email.",
			},
			"bio": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user's bio.",
			},
			"gpg_keys": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of user's GPG keys.",
			},
			"ssh_keys": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of user's SSH keys.",
			},
			"public_repos": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of public repositories.",
			},
			"public_gists": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of public gists.",
			},
			"followers": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of followers.",
			},
			"following": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of following users.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation date.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update date.",
			},
			"suspended_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The suspended date if the user is suspended.",
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Node ID of the user.",
			},
		},
	}
}

func dataSourceGithubUserRead(d *schema.ResourceData, meta any) error {
	username := d.Get("username").(string)

	client := meta.(*Owner).v3client
	ctx := context.Background()

	user, _, err := client.Users.Get(ctx, username)
	if err != nil {
		return err
	}

	gpg, _, err := client.Users.ListGPGKeys(ctx, user.GetLogin(), nil)
	if err != nil {
		return err
	}
	ssh, _, err := client.Users.ListKeys(ctx, user.GetLogin(), nil)
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
	if err = d.Set("login", user.GetLogin()); err != nil {
		return err
	}
	if err = d.Set("avatar_url", user.GetAvatarURL()); err != nil {
		return err
	}
	if err = d.Set("gravatar_id", user.GetGravatarID()); err != nil {
		return err
	}
	if err = d.Set("site_admin", user.GetSiteAdmin()); err != nil {
		return err
	}
	if err = d.Set("company", user.GetCompany()); err != nil {
		return err
	}
	if err = d.Set("blog", user.GetBlog()); err != nil {
		return err
	}
	if err = d.Set("location", user.GetLocation()); err != nil {
		return err
	}
	if err = d.Set("name", user.GetName()); err != nil {
		return err
	}
	if err = d.Set("email", user.GetEmail()); err != nil {
		return err
	}
	if err = d.Set("bio", user.GetBio()); err != nil {
		return err
	}
	if err = d.Set("gpg_keys", gpgKeys); err != nil {
		return err
	}
	if err = d.Set("ssh_keys", sshKeys); err != nil {
		return err
	}
	if err = d.Set("public_repos", user.GetPublicRepos()); err != nil {
		return err
	}
	if err = d.Set("public_gists", user.GetPublicGists()); err != nil {
		return err
	}
	if err = d.Set("followers", user.GetFollowers()); err != nil {
		return err
	}
	if err = d.Set("following", user.GetFollowing()); err != nil {
		return err
	}
	if err = d.Set("created_at", user.GetCreatedAt().String()); err != nil {
		return err
	}
	if err = d.Set("updated_at", user.GetUpdatedAt().String()); err != nil {
		return err
	}
	if err = d.Set("suspended_at", user.GetSuspendedAt().String()); err != nil {
		return err
	}
	if err = d.Set("node_id", user.GetNodeID()); err != nil {
		return err
	}

	return nil
}
