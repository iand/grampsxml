package grampsxml

import "encoding/xml"

const grampsXMLVersion171 = "http://gramps-project.org/xml/1.7.1/"

// Database171 is the versioned XML element for Gramps XML schema 1.7.1.
// Use to decode or encode XML in this exact schema version.
type Database171 struct {
	XMLName      xml.Name      `xml:"http://gramps-project.org/xml/1.7.1/ database"`
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

func unmarshalDatabase171(db *Database, d *xml.Decoder, start *xml.StartElement) error {
	var src Database171
	if err := d.DecodeElement(&src, start); err != nil {
		return err
	}
	db.Header = src.Header
	db.NameFormats = src.NameFormats
	db.Tags = src.Tags
	db.Events = src.Events
	db.People = src.People
	db.Families = src.Families
	db.Citations = src.Citations
	db.Sources = src.Sources
	db.Places = src.Places
	db.Objects = src.Objects
	db.Repositories = src.Repositories
	db.Notes = src.Notes
	db.Bookmarks = src.Bookmarks
	db.Namemaps = src.Namemaps
	return nil
}
