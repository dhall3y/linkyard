package links

type Link struct {
	ID           string
	Type         string
	DateAdded    uint
	LastModified uint
	URI          string
	Title        string
	ParentID     string
}
