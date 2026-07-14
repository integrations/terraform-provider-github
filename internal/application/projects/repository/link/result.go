package link

type Result struct {
	ProjectID    string
	RepositoryID string
	Owner        string
	Name         string
}

type AttachInput struct {
	ProjectID string
	Owner     string
	Name      string
}
