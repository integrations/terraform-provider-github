package item

type Result struct {
	ID        string
	ProjectID string
	ContentID string
	Archived  bool
}

type AddInput struct {
	ProjectID string
	ContentID string
	Archived  bool
}
