package field

import "github.com/shurcooL/githubv4"

type baseNode struct {
	ID       githubv4.String
	Name     githubv4.String
	DataType githubv4.ProjectV2FieldType
	Project  struct{ ID githubv4.String }
}

type singleSelectNode struct {
	baseNode
	Options []struct {
		ID          githubv4.String
		Name        githubv4.String
		Description githubv4.String
		Color       githubv4.ProjectV2SingleSelectFieldOptionColor
	}
}

type iterationValue struct {
	ID        githubv4.String
	Title     githubv4.String
	StartDate githubv4.String
	Duration  githubv4.Int
}

type iterationNode struct {
	baseNode
	Configuration struct {
		Duration            githubv4.Int
		Iterations          []iterationValue
		CompletedIterations []iterationValue
	}
}

type node struct {
	Typename     githubv4.String            `graphql:"__typename"`
	Field        struct{ baseNode }         `graphql:"... on ProjectV2Field"`
	SingleSelect struct{ singleSelectNode } `graphql:"... on ProjectV2SingleSelectField"`
	Iteration    struct{ iterationNode }    `graphql:"... on ProjectV2IterationField"`
}
