//go:build gramps_schema180

package grampsxml

import "encoding/xml"

const grampsXMLVersion180 = "http://gramps-project.org/xml/1.8.0/"

func unmarshalDatabase180(db *Database, d *xml.Decoder, start *xml.StartElement) error {
	var src Database180
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
	db.DNATests = src.DNATests
	db.DNAMatches = src.DNAMatches
	db.Bookmarks = src.Bookmarks
	db.Namemaps = src.Namemaps
	return nil
}

// Database180 is the versioned XML element for Gramps XML schema 1.8.0.
// Use to decode or encode XML in this exact schema version.
type Database180 struct {
	XMLName      xml.Name      `xml:"http://gramps-project.org/xml/1.8.0/ database"`
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
	DNATests     *DNATests     `xml:"dnatests,omitempty"`
	DNAMatches   *DNAMatches   `xml:"dnamatches,omitempty"`
	Bookmarks    *Bookmarks    `xml:"bookmarks,omitempty"`
	Namemaps     *Namemaps     `xml:"namemaps,omitempty"`
}

type DNATests struct {
	DNATest []DNATest `xml:"dnatest,omitempty"`
}

type DNAMatches struct {
	DNAMatch []DNAMatch `xml:"dnamatch,omitempty"`
}

// PersonLink is a reference element carrying only an hlink attribute,
// used wherever a child element links to a person (e.g. in dnatest and shared_ancestor).
type PersonLink struct {
	Hlink string `xml:"hlink,attr"`
}

// SubjectTest is the <subject_test hlink="..."/> element in a dnamatch.
type SubjectTest struct {
	Hlink string `xml:"hlink,attr"`
}

// MatchTest is the <match_test hlink="..."/> element in a dnamatch.
type MatchTest struct {
	Hlink string `xml:"hlink,attr"`
}

// DNATest represents one DNA kit for one person at one provider.
type DNATest struct {
	ID          *string       `xml:"id,attr,omitempty"`
	Handle      string        `xml:"handle,attr"`
	Priv        *bool         `xml:"priv,attr,omitempty"`
	Change      string        `xml:"change,attr"`
	Person      *PersonLink   `xml:"person,omitempty"`
	AccountName *string       `xml:"account_name,omitempty"`
	Provider    *string       `xml:"provider,omitempty"`
	KitID       *string       `xml:"kit_id,omitempty"`
	TestType    *string       `xml:"test_type,omitempty"`
	GenomeBuild *string       `xml:"genome_build,omitempty"`
	Daterange   *Daterange    `xml:"daterange,omitempty"`
	Datespan    *Datespan     `xml:"datespan,omitempty"`
	Dateval     *Dateval      `xml:"dateval,omitempty"`
	Datestr     *Datestr      `xml:"datestr,omitempty"`
	Haplogroup  *string       `xml:"haplogroup,omitempty"`
	Attribute   []Attribute   `xml:"attribute,omitempty"`
	Objref      []Objref      `xml:"objref,omitempty"`
	Noteref     []Noteref     `xml:"noteref,omitempty"`
	Citationref []Citationref `xml:"citationref,omitempty"`
	Tagref      []Tagref      `xml:"tagref,omitempty"`
}

// DNAMatch represents a pairwise DNA match between two kits.
type DNAMatch struct {
	ID                    *string               `xml:"id,attr,omitempty"`
	Handle                string                `xml:"handle,attr"`
	Priv                  *bool                 `xml:"priv,attr,omitempty"`
	Change                string                `xml:"change,attr"`
	SubjectTest           *SubjectTest          `xml:"subject_test,omitempty"`
	MatchTest             *MatchTest            `xml:"match_test,omitempty"`
	SharedCM              *SharedCM             `xml:"shared_cm,omitempty"`
	PercentShared         *PercentShared        `xml:"percent_shared,omitempty"`
	SegmentCount          *SegmentCount         `xml:"segment_count,omitempty"`
	LargestSegmentCM      *LargestSegmentCM     `xml:"largest_segment_cm,omitempty"`
	PredictedRelationship *string               `xml:"predicted_relationship,omitempty"`
	PredictedGenerations  *PredictedGenerations `xml:"predicted_generations,omitempty"`
	SharedAncestor        []SharedAncestor      `xml:"shared_ancestor,omitempty"`
	DNASegment            []DNASegment          `xml:"dna_segment,omitempty"`
	Attribute             []Attribute           `xml:"attribute,omitempty"`
	Objref                []Objref              `xml:"objref,omitempty"`
	Noteref               []Noteref             `xml:"noteref,omitempty"`
	Citationref           []Citationref         `xml:"citationref,omitempty"`
	Tagref                []Tagref              `xml:"tagref,omitempty"`
}

// SharedCM holds the total shared centimorgans as a val attribute.
type SharedCM struct {
	Val float64 `xml:"val,attr"`
}

// PercentShared holds the percentage of genome shared as a val attribute.
type PercentShared struct {
	Val float64 `xml:"val,attr"`
}

// SegmentCount holds the number of shared segments as a val attribute.
type SegmentCount struct {
	Val int `xml:"val,attr"`
}

// LargestSegmentCM holds the largest segment size in cM as a val attribute.
type LargestSegmentCM struct {
	Val float64 `xml:"val,attr"`
}

// PredictedGenerations holds the estimated generations to MRCA as a val attribute.
type PredictedGenerations struct {
	Val float64 `xml:"val,attr"`
}

// SharedAncestor records a proposed or confirmed MRCA for a DNAMatch.
type SharedAncestor struct {
	Confidence  string        `xml:"confidence,attr"`
	Description *string       `xml:"description,omitempty"`
	Person      *PersonLink   `xml:"person,omitempty"`
	Noteref     []Noteref     `xml:"noteref,omitempty"`
	Citationref []Citationref `xml:"citationref,omitempty"`
}

// DNASegment is a single shared chromosomal segment within a DNAMatch.
type DNASegment struct {
	Chromosome string  `xml:"chromosome,attr"`
	StartBP    int     `xml:"start_bp,attr"`
	EndBP      int     `xml:"end_bp,attr"`
	SharedCM   float64 `xml:"shared_cm,attr"`
	SNPCount   int     `xml:"snp_count,attr"`
	Phase      int     `xml:"phase,attr"`
	IBDState   *int    `xml:"ibd_state,attr,omitempty"`
	StartRSID  *string `xml:"start_rsid,attr,omitempty"`
	EndRSID    *string `xml:"end_rsid,attr,omitempty"`
}
