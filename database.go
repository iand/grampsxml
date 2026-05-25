//go:build !gramps_schema180

package grampsxml

import (
	"encoding/xml"
	"fmt"
)

// Database is the top-level container for a Gramps XML export.
// It supports schema versions 1.7.1 and 1.7.2.
type Database struct {
	Header       Header        `xml:"header"`
	NameFormats  []NameFormat  `xml:"name-formats>format"`
	Tags         *Tags         `xml:"tags,omitempty"`
	Events       *Events       `xml:"events,omitempty"`
	People       *People       `xml:"people,omitempty"`
	Families     *Families     `xml:"families,omitempty"`
	Citations    *Citations    `xml:"citations,omitempty"`
	Sources      *Sources      `xml:"sources,omitempty"`
	Places       *Places       `xml:"places,omitempty"`
	Objects      *Objects      `xml:"objects,omitempty"`
	Repositories *Repositories `xml:"repositories,omitempty"`
	Notes        *Notes        `xml:"notes,omitempty"`
	Bookmarks    *Bookmarks    `xml:"bookmarks,omitempty"`
	Namemaps     *Namemaps     `xml:"namemaps,omitempty"`
}

func (db *Database) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Space {
	case grampsXMLVersion171:
		return unmarshalDatabase171(db, d, &start)
	case grampsXMLVersion172:
		return unmarshalDatabase172(db, d, &start)
	default:
		return fmt.Errorf("grampsxml: unsupported XML namespace %q", start.Name.Space)
	}
}
