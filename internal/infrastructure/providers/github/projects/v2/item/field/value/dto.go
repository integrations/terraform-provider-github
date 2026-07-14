package value

import "github.com/shurcooL/githubv4"

type node struct {
	Typename     githubv4.String                       `graphql:"__typename"`
	Text         struct{ Text githubv4.String }        `graphql:"... on ProjectV2ItemFieldTextValue"`
	Number       struct{ Number githubv4.Float }       `graphql:"... on ProjectV2ItemFieldNumberValue"`
	Date         struct{ Date githubv4.String }        `graphql:"... on ProjectV2ItemFieldDateValue"`
	SingleSelect struct{ OptionID githubv4.String }    `graphql:"... on ProjectV2ItemFieldSingleSelectValue"`
	Iteration    struct{ IterationID githubv4.String } `graphql:"... on ProjectV2ItemFieldIterationValue"`
}
