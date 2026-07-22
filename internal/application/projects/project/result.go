package project

type OwnerKind string

const (
	OwnerOrganization OwnerKind = "organization"
	OwnerUser         OwnerKind = "user"
)

type Result struct {
	ID               string
	Number           int
	Title            string
	ShortDescription string
	Readme           string
	Public           bool
	Closed           bool
	URL              string
	OwnerKind        OwnerKind
	Owner            string
	OwnerID          int
}

type CreateInput struct {
	OwnerKind        OwnerKind
	Owner            string
	Title            string
	ShortDescription string
	Readme           string
	Public           bool
	Closed           bool
}

type UpdateInput struct {
	ID               string
	Title            string
	ShortDescription string
	Readme           string
	Public           bool
	Closed           bool
}
