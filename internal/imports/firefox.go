package imports

import (
	"fmt"
	"linkyard/internal/links"

	"github.com/gofrs/uuid"
)

type FirefoxLink struct {
	ID           *uuid.UUID    `json:"-"`
	TypeCode     uint8         `json:"typeCode"`
	DateAdded    uint          `json:"dateAdded"`
	LastModified uint          `json:"lastModified"`
	URI          string        `json:"uri"`
	Title        string        `json:"title"`
	Children     []FirefoxLink `json:"children"`
	ParentID     *uuid.UUID    `json:"-"`
}

func (f *FirefoxLink) goThroughLinks(parentID *uuid.UUID, links *[]links.Link, uuidGen *uuid.Gen) {
	id, err := uuidGen.NewV4()
	if err != nil {
		fmt.Println("failed to generate uuid id", err)
		return
	}
	// append current link
	f.ParentID = parentID
	f.ID = &id
	*links = append(*links, *f.format())

	// if it has children go through them
	for i := range f.Children {
		f.Children[i].goThroughLinks(&id, links, uuidGen)
	}
}

func (f *FirefoxLink) format() *links.Link {
	newLink := links.Link{
		Title:        f.Title,
		URI:          f.URI,
		DateAdded:    f.DateAdded,
		LastModified: f.LastModified,
	}

	if f.ParentID != nil {
		newLink.ParentID = f.ParentID.String()
	}

	if f.ID != nil {
		newLink.ID = f.ID.String()
	}

	if f.TypeCode == 1 {
		newLink.Type = "link"
	} else {
		newLink.Type = "folder"
	}

	return &newLink
}
