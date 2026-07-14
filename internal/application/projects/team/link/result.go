package link

type Result struct {
	ProjectID    string
	TeamID       string
	Organization string
	Slug         string
}

type AttachInput struct {
	ProjectID    string
	Organization string
	Slug         string
}
