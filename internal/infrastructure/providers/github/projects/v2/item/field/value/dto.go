package value

import "github.com/shurcooL/githubv4"

type node struct {
	Typename githubv4.String `graphql:"__typename"`
	Field    struct {
		Typename     githubv4.String              `graphql:"__typename"`
		Field        struct{ ID githubv4.String } `graphql:"... on ProjectV2Field"`
		SingleSelect struct{ ID githubv4.String } `graphql:"... on ProjectV2SingleSelectField"`
		Iteration    struct{ ID githubv4.String } `graphql:"... on ProjectV2IterationField"`
	}
	Text         struct{ Text githubv4.String }        `graphql:"... on ProjectV2ItemFieldTextValue"`
	Number       struct{ Number githubv4.Float }       `graphql:"... on ProjectV2ItemFieldNumberValue"`
	Date         struct{ Date githubv4.String }        `graphql:"... on ProjectV2ItemFieldDateValue"`
	SingleSelect struct{ OptionID githubv4.String }    `graphql:"... on ProjectV2ItemFieldSingleSelectValue"`
	Iteration    struct{ IterationID githubv4.String } `graphql:"... on ProjectV2ItemFieldIterationValue"`
}

type pageInfo struct {
	EndCursor   githubv4.String
	HasNextPage githubv4.Boolean
}
