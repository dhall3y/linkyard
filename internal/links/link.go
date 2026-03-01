package links

import "time"

type Link struct {
	ID           string
	Type         string
	DateAdded    time.Time
	LastModified time.Time
	URI          string
	Title        string
	ParentID     string
}
