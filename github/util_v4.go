package github

import (
	"strings"

	"github.com/shurcooL/githubv4"
)

type PageInfo struct {
	EndCursor   githubv4.String
	HasNextPage bool
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

func githubv4IDSliceEmpty(ss []string) []githubv4.ID {
	vGh4 := make([]githubv4.ID, 0)
	for _, s := range ss {
		vGh4 = append(vGh4, githubv4.ID(s))
	}
	return vGh4
}

func githubv4NewStringSlice(v []githubv4.String) *[]githubv4.String { return &v }

func githubv4NewIDSlice(v []githubv4.ID) *[]githubv4.ID { return &v }

/**
 * Checks if the error message represents
 * that graphql query did not find node by its global id.
 * Graphql error responses return 200 OK http codes
 * so we have to inspect the messages =/
 */
func githubv4IsNodeNotFoundError(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), "Could not resolve to a node with the global id")
}
