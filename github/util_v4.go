package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/shurcooL/githubv4"
)

type PageInfo struct {
	EndCursor   githubv4.String
	HasNextPage bool
}

func expandNestedSet(m map[string]interface{}, target string) []string {
	res := make([]string, 0)
	if v, ok := m[target]; ok {
		vL := v.(*schema.Set).List()
		for _, v := range vL {
			res = append(res, v.(string))
		}
	}
	return res
}

func githubv4StringSlice(ss []string) []githubv4.String {
	var vGh4 []githubv4.String
	for _, s := range ss {
		vGh4 = append(vGh4, githubv4.String(s))
	}
	return vGh4
}

func githubv4IDSlice(ss []string) []githubv4.ID {
	var vGh4 []githubv4.ID
	for _, s := range ss {
		vGh4 = append(vGh4, githubv4.ID(s))
	}
	return vGh4
}

func githubv4NewStringSlice(v []githubv4.String) *[]githubv4.String { return &v }

func githubv4NewIDSlice(v []githubv4.ID) *[]githubv4.ID { return &v }
