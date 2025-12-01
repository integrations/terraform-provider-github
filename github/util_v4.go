package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

type PageInfo struct {
	EndCursor   githubv4.String
	HasNextPage bool
}

func expandNestedSet(m map[string]any, target string) []string {
	res := make([]string, 0)
	if v, ok := m[target]; ok {
		vL := v.(*schema.Set).List()
		for _, v := range vL {
			res = append(res, v.(string))
		}
	}
	return res
}

func githubv4StringSliceEmpty(ss []string) []githubv4.String {
	vGh4 := make([]githubv4.String, 0)
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

func githubv4IDSliceEmpty(ss []string) []githubv4.ID {
	vGh4 := make([]githubv4.ID, 0)
	for _, s := range ss {
		vGh4 = append(vGh4, githubv4.ID(s))
	}
	return vGh4
}

func githubv4NewStringSlice(v []githubv4.String) *[]githubv4.String { return &v }

func githubv4NewIDSlice(v []githubv4.ID) *[]githubv4.ID { return &v }
