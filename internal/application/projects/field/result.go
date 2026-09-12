package field

import "time"

type Result struct {
	ID                     string
	ProjectID              string
	Name                   string
	DataType               string
	SingleSelectOptions    []SingleSelectOption
	IterationConfiguration *IterationConfiguration
}

type SingleSelectOption struct {
	ID          string
	Name        string
	Description string
	Color       string
}

type IterationConfiguration struct {
	Duration            int
	Iterations          []Iteration
	CompletedIterations []Iteration
}

type Iteration struct {
	ID        string
	Title     string
	StartDate time.Time
	Duration  int
}

type Configuration struct {
	SingleSelectOptions []SingleSelectOptionInput
	Iteration           *IterationConfigurationInput
}

type SingleSelectOptionInput struct {
	Name        string
	Description string
	Color       string
}

type IterationConfigurationInput struct {
	StartDate  time.Time
	Duration   int
	Iterations []IterationInput
}

type IterationInput struct {
	Title     string
	StartDate time.Time
	Duration  int
}

type CreateInput struct {
	ProjectID     string
	Name          string
	DataType      string
	Configuration Configuration
}

type UpdateInput struct {
	ID            string
	Name          string
	Configuration *Configuration
}
