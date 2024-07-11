package grampsxml

import "encoding/xml"

type Database struct {
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

type Header struct {
	Created    Created     `xml:"created"`
	Researcher *Researcher `xml:"researcher,omitempty"`
	Mediapath  *string     `xml:"mediapath,omitempty"`
}

type Created struct {
	Date    string `xml:"date,attr"`
	Version string `xml:"version,attr"`
}

type Researcher struct {
	Resname     *string `xml:"resname,omitempty"`
	Resaddr     *string `xml:"resaddr,omitempty"`
	Reslocality *string `xml:"reslocality,omitempty"`
	Rescity     *string `xml:"rescity,omitempty"`
	Resstate    *string `xml:"resstate,omitempty"`
	Rescountry  *string `xml:"rescountry,omitempty"`
	Respostal   *string `xml:"respostal,omitempty"`
	Resphone    *string `xml:"resphone,omitempty"`
	Resemail    *string `xml:"resemail,omitempty"`
}

type People struct {
	Default *string  `xml:"default,omitempty"`
	Home    *string  `xml:"home,omitempty"`
	Person  []Person `xml:"person,omitempty"`
}

type Person struct {
	ID          *string       `xml:"id,attr,omitempty"`
	Handle      string        `xml:"handle,attr"`
	Priv        *bool         `xml:"priv,attr,omitempty"`
	Change      string        `xml:"change,attr"`
	Gender      string        `xml:"gender"`
	Name        []Name        `xml:"name,omitempty"`
	Eventref    []Eventref    `xml:"eventref,omitempty"`
	LdsOrd      []LdsOrd      `xml:"lds_ord,omitempty"`
	Objref      []Objref      `xml:"objref,omitempty"`
	Address     []Address     `xml:"address,omitempty"`
	Attribute   []Attribute   `xml:"attribute,omitempty"`
	Url         []Url         `xml:"url,omitempty"`
	Childof     []Childof     `xml:"childof,omitempty"`
	Parentin    []Parentin    `xml:"parentin,omitempty"`
	Personref   []Personref   `xml:"personref,omitempty"`
	Noteref     []Noteref     `xml:"noteref,omitempty"`
	Citationref []Citationref `xml:"citationref,omitempty"`
	Tagref      []Tagref      `xml:"tagref,omitempty"`
}

type Name struct {
	Alt         *bool         `xml:"alt,attr,omitempty"`
	Type        *string       `xml:"type,attr,omitempty"`
	Priv        *bool         `xml:"priv,attr,omitempty"`
	Sort        *string       `xml:"sort,attr,omitempty"`
	Display     *string       `xml:"display,attr,omitempty"`
	First       *string       `xml:"first,omitempty"`
	Call        *string       `xml:"call,omitempty"`
	Suffix      *string       `xml:"suffix,omitempty"`
	Title       *string       `xml:"title,omitempty"`
	Nick        *string       `xml:"nick,omitempty"`
	Familynick  *string       `xml:"familynick,omitempty"`
	Group       *string       `xml:"group,omitempty"`
	Surname     []Surname     `xml:"surname"`
	Daterange   *Daterange    `xml:"daterange,omitempty"`
	Datespan    *Datespan     `xml:"datespan,omitempty"`
	Dateval     *Dateval      `xml:"dateval,omitempty"`
	Datestr     *Datestr      `xml:"datestr,omitempty"`
	Noteref     []Noteref     `xml:"noteref,omitempty"`
	Citationref []Citationref `xml:"citationref,omitempty"`
}

type Surname struct {
	Prefix     *string `xml:"prefix,attr,omitempty"`
	Prim       *bool   `xml:"prim,attr,omitempty"`
	Derivation *string `xml:"derivation,attr,omitempty"`
	Connector  *string `xml:"connector,attr,omitempty"`
	Surname    string  `xml:",chardata"`
}

type Childof struct {
	Hlink string `xml:"hlink,attr"`
}

type Parentin struct {
	Hlink string `xml:"hlink,attr"`
}

type Personref struct {
	Hlink       string        `xml:"hlink,attr"`
	Priv        *bool         `xml:"priv,attr,omitempty"`
	Rel         string        `xml:"rel,attr"`
	Citationref []Citationref `xml:"citationref,omitempty"`
	Noteref     []Noteref     `xml:"noteref,omitempty"`
}

type Address struct {
	Daterange   *Daterange    `xml:"daterange,omitempty"`
	Datespan    *Datespan     `xml:"datespan,omitempty"`
	Dateval     *Dateval      `xml:"dateval,omitempty"`
	Datestr     *Datestr      `xml:"datestr,omitempty"`
	Street      *string       `xml:"street,omitempty"`
	Locality    *string       `xml:"locality,omitempty"`
	City        *string       `xml:"city,omitempty"`
	County      *string       `xml:"county,omitempty"`
	State       *string       `xml:"state,omitempty"`
	Country     *string       `xml:"country,omitempty"`
	Postal      *string       `xml:"postal,omitempty"`
	Phone       *string       `xml:"phone,omitempty"`
	Noteref     []Noteref     `xml:"noteref,omitempty"`
	Citationref []Citationref `xml:"citationref,omitempty"`
}

type Families struct {
	Family []Family `xml:"family,omitempty"`
}

type Family struct {
	ID          *string       `xml:"id,attr,omitempty"`
	Handle      string        `xml:"handle,attr"`
	Priv        *bool         `xml:"priv,attr,omitempty"`
	Change      string        `xml:"change,attr"`
	Rel         *Rel          `xml:"rel,omitempty"`
	Father      *Father       `xml:"father,omitempty"`
	Mother      *Mother       `xml:"mother,omitempty"`
	Eventref    []Eventref    `xml:"eventref,omitempty"`
	LdsOrd      []LdsOrd      `xml:"lds_ord,omitempty"`
	Objref      []Objref      `xml:"objref,omitempty"`
	Childref    []Childref    `xml:"childref,omitempty"`
	Attribute   []Attribute   `xml:"attribute,omitempty"`
	Noteref     []Noteref     `xml:"noteref,omitempty"`
	Citationref []Citationref `xml:"citationref,omitempty"`
	Tagref      []Tagref      `xml:"tagref,omitempty"`
}

type Father struct {
	Hlink string `xml:"hlink,attr"`
}

type Mother struct {
	Hlink string `xml:"hlink,attr"`
}

type Childref struct {
	Hlink string  `xml:"hlink,attr"`
	Priv  *bool   `xml:"priv,attr,omitempty"`
	Mrel  *string `xml:"mrel,attr,omitempty"`
	Frel  *string `xml:"frel,attr,omitempty"`
}

type Rel struct {
	Type string `xml:"type,attr"`
}

type Events struct {
	Event []Event `xml:"event,omitempty"`
}

type Event struct {
	ID          *string       `xml:"id,attr,omitempty"`
	Handle      string        `xml:"handle,attr"`
	Priv        *bool         `xml:"priv,attr,omitempty"`
	Change      string        `xml:"change,attr"`
	Type        *string       `xml:"type"`
	Daterange   *Daterange    `xml:"daterange,omitempty"`
	Datespan    *Datespan     `xml:"datespan,omitempty"`
	Dateval     *Dateval      `xml:"dateval,omitempty"`
	Datestr     *Datestr      `xml:"datestr,omitempty"`
	Place       *Place        `xml:"place,omitempty"`
	Cause       *string       `xml:"cause,omitempty"`
	Description *string       `xml:"description,omitempty"`
	Attribute   []Attribute   `xml:"attribute,omitempty"`
	Noteref     []Noteref     `xml:"noteref,omitempty"`
	Citationref []Citationref `xml:"citationref,omitempty"`
	Objref      []Objref      `xml:"objref,omitempty"`
	Tagref      []Tagref      `xml:"tagref,omitempty"`
}

type Sources struct {
	Source []Source `xml:"source,omitempty"`
}

type Source struct {
	ID           *string        `xml:"id,attr,omitempty"`
	Handle       string         `xml:"handle,attr"`
	Priv         *bool          `xml:"priv,attr,omitempty"`
	Change       string         `xml:"change,attr"`
	Stitle       *string        `xml:"stitle,omitempty"`
	Sauthor      *string        `xml:"sauthor,omitempty"`
	Spubinfo     *string        `xml:"spubinfo,omitempty"`
	Sabbrev      *string        `xml:"sabbrev,omitempty"`
	Noteref      []Noteref      `xml:"noteref,omitempty"`
	Objref       []Objref       `xml:"objref,omitempty"`
	Srcattribute []Srcattribute `xml:"srcattribute,omitempty"`
	Reporef      []Reporef      `xml:"reporef,omitempty"`
	Tagref       []Tagref       `xml:"tagref,omitempty"`
}

type Places struct {
	Place []Placeobj `xml:"placeobj,omitempty"`
}

type Placeobj struct {
	ID          *string       `xml:"id,attr,omitempty"`
	Handle      string        `xml:"handle,attr"`
	Priv        *bool         `xml:"priv,attr,omitempty"`
	Change      string        `xml:"change,attr"`
	Type        string        `xml:"type,attr"`
	Ptitle      *string       `xml:"ptitle,omitempty"`
	Pname       []Pname       `xml:"pname"`
	Code        *string       `xml:"code,omitempty"`
	Coord       *Coord        `xml:"coord,omitempty"`
	Placeref    []Placeref    `xml:"placeref,omitempty"`
	Location    []Location    `xml:"location,omitempty"`
	Url         []Url         `xml:"url,omitempty"`
	Objref      []Objref      `xml:"objref,omitempty"`
	Noteref     []Noteref     `xml:"noteref,omitempty"`
	Citationref []Citationref `xml:"citationref,omitempty"`
	Tagref      []Tagref      `xml:"tagref,omitempty"`
}

type Pname struct {
	Lang      *string    `xml:"lang,attr,omitempty"`
	Value     string     `xml:"value,attr"`
	Daterange *Daterange `xml:"daterange,omitempty"`
	Datespan  *Datespan  `xml:"datespan,omitempty"`
	Dateval   *Dateval   `xml:"dateval,omitempty"`
	Datestr   *Datestr   `xml:"datestr,omitempty"`
}

type Coord struct {
	Long string `xml:"long,attr"`
	Lat  string `xml:"lat,attr"`
}

type Location struct {
	Street   *string `xml:"street,attr,omitempty"`
	Locality *string `xml:"locality,attr,omitempty"`
	City     *string `xml:"city,attr,omitempty"`
	Parish   *string `xml:"parish,attr,omitempty"`
	County   *string `xml:"county,attr,omitempty"`
	State    *string `xml:"state,attr,omitempty"`
	Country  *string `xml:"country,attr,omitempty"`
	Postal   *string `xml:"postal,attr,omitempty"`
	Phone    *string `xml:"phone,attr,omitempty"`
}

type Objects struct {
	Object []Object `xml:"object,omitempty"`
}

type Object struct {
	ID          *string       `xml:"id,attr,omitempty"`
	Handle      string        `xml:"handle,attr"`
	Priv        *bool         `xml:"priv,attr,omitempty"`
	Change      string        `xml:"change,attr"`
	File        File          `xml:"file"`
	Attribute   []Attribute   `xml:"attribute,omitempty"`
	Noteref     []Noteref     `xml:"noteref,omitempty"`
	Daterange   *Daterange    `xml:"daterange,omitempty"`
	Datespan    *Datespan     `xml:"datespan,omitempty"`
	Dateval     *Dateval      `xml:"dateval,omitempty"`
	Datestr     *Datestr      `xml:"datestr,omitempty"`
	Citationref []Citationref `xml:"citationref,omitempty"`
	Tagref      []Tagref      `xml:"tagref,omitempty"`
}

type File struct {
	Src         string  `xml:"src,attr"`
	Mime        string  `xml:"mime,attr"`
	Checksum    *string `xml:"checksum,attr,omitempty"`
	Description string  `xml:"description,attr"`
}

type Repositories struct {
	Repository []Repository `xml:"repository,omitempty"`
}

type Repository struct {
	ID      *string   `xml:"id,attr,omitempty"`
	Handle  string    `xml:"handle,attr"`
	Priv    *bool     `xml:"priv,attr,omitempty"`
	Change  string    `xml:"change,attr"`
	Rname   string    `xml:"rname"`
	Type    string    `xml:"type"`
	Address []Address `xml:"address,omitempty"`
	Url     []Url     `xml:"url,omitempty"`
	Noteref []Noteref `xml:"noteref,omitempty"`
	Tagref  []Tagref  `xml:"tagref,omitempty"`
}

type Notes struct {
	Note []Note `xml:"note,omitempty"`
}

type Note struct {
	ID     *string  `xml:"id,attr,omitempty"`
	Handle string   `xml:"handle,attr"`
	Priv   *bool    `xml:"priv,attr,omitempty"`
	Change string   `xml:"change,attr"`
	Format *bool    `xml:"format,attr,omitempty"`
	Type   string   `xml:"type,attr"`
	Text   string   `xml:"text"`
	Style  []Style  `xml:"style,omitempty"`
	Tagref []Tagref `xml:"tagref,omitempty"`
}

type Style struct {
	Name  string  `xml:"name,attr"`
	Value *string `xml:"value,attr,omitempty"`
	Range []Range `xml:"range"`
}

type Range struct {
	Start int `xml:"start,attr"`
	End   int `xml:"end,attr"`
}

type Tags struct {
	Tag []Tag `xml:"tag,omitempty"`
}

type Tag struct {
	Handle   string `xml:"handle,attr"`
	Name     string `xml:"name,attr"`
	Color    string `xml:"color,attr"`
	Priority string `xml:"priority,attr"`
	Change   string `xml:"change,attr"`
}

type Citations struct {
	Citation []Citation `xml:"citation,omitempty"`
}

type Citation struct {
	ID           *string        `xml:"id,attr,omitempty"`
	Handle       string         `xml:"handle,attr"`
	Priv         *bool          `xml:"priv,attr,omitempty"`
	Change       string         `xml:"change,attr"`
	Daterange    *Daterange     `xml:"daterange,omitempty"`
	Datespan     *Datespan      `xml:"datespan,omitempty"`
	Dateval      *Dateval       `xml:"dateval,omitempty"`
	Datestr      *Datestr       `xml:"datestr,omitempty"`
	Page         *string        `xml:"page,omitempty"`
	Confidence   string         `xml:"confidence"`
	Noteref      []Noteref      `xml:"noteref,omitempty"`
	Objref       []Objref       `xml:"objref,omitempty"`
	Srcattribute []Srcattribute `xml:"srcattribute,omitempty"`
	Sourceref    *Sourceref     `xml:"sourceref,omitempty"`
	Tagref       []Tagref       `xml:"tagref,omitempty"`
}

type Bookmarks struct {
	Bookmark []Bookmark `xml:"bookmark,omitempty"`
}

type Bookmark struct {
	Target string `xml:"target,attr"`
	Hlink  string `xml:"hlink,attr"`
}

type Namemaps struct {
	Map []Map `xml:"map,omitempty"`
}

type Map struct {
	Type  string `xml:"type,attr"`
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

type NameFormat struct {
	Number string `xml:"number,attr"`
	Name   string `xml:"name,attr"`
	Fmtstr string `xml:"fmt_str,attr"`
	Active *bool  `xml:"active,attr,omitempty"`
}

type Daterange struct {
	Start     string  `xml:"start,attr"`
	Stop      string  `xml:"stop,attr"`
	Quality   *string `xml:"quality,attr,omitempty"`
	Cformat   *string `xml:"cformat,attr,omitempty"`
	Dualdated *bool   `xml:"dualdated,attr,omitempty"`
	Newyear   *string `xml:"newyear,attr,omitempty"`
}

type Datespan struct {
	Start     string  `xml:"start,attr"`
	Stop      string  `xml:"stop,attr"`
	Quality   *string `xml:"quality,attr,omitempty"`
	Cformat   *string `xml:"cformat,attr,omitempty"`
	Dualdated *bool   `xml:"dualdated,attr,omitempty"`
	Newyear   *string `xml:"newyear,attr,omitempty"`
}

type Dateval struct {
	Val       string  `xml:"val,attr"`
	Type      *string `xml:"type,attr,omitempty"`
	Quality   *string `xml:"quality,attr,omitempty"`
	Cformat   *string `xml:"cformat,attr,omitempty"`
	Dualdated *bool   `xml:"dualdated,attr,omitempty"`
	Newyear   *string `xml:"newyear,attr,omitempty"`
}

type Datestr struct {
	Val string `xml:"val,attr"`
}

type Citationref struct {
	Hlink string `xml:"hlink,attr"`
}

type Sourceref struct {
	Hlink string `xml:"hlink,attr"`
}

type Eventref struct {
	Hlink     string      `xml:"hlink,attr"`
	Priv      *bool       `xml:"priv,attr,omitempty"`
	Role      *string     `xml:"role,attr,omitempty"`
	Attribute []Attribute `xml:"attribute,omitempty"`
	Noteref   []Noteref   `xml:"noteref,omitempty"`
}

type Reporef struct {
	Hlink  string  `xml:"hlink,attr"`
	Priv   *bool   `xml:"priv,attr,omitempty"`
	Callno *string `xml:"callno,attr,omitempty"`
	Medium *string `xml:"medium,attr,omitempty"`
}

type Noteref struct {
	Hlink string `xml:"hlink,attr"`
}

type Tagref struct {
	Hlink string `xml:"hlink,attr"`
}

type Attribute struct {
	Priv        *bool         `xml:"priv,attr,omitempty"`
	Type        string        `xml:"type,attr"`
	Value       string        `xml:"value,attr"`
	Citationref []Citationref `xml:"citationref,omitempty"`
}

type Srcattribute struct {
	Priv  *bool  `xml:"priv,attr,omitempty"`
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

type Place struct {
	Hlink string `xml:"hlink,attr"`
}

type Url struct {
	Priv        *bool   `xml:"priv,attr,omitempty"`
	Type        *string `xml:"type,attr,omitempty"`
	Href        string  `xml:"href,attr"`
	Description *string `xml:"description,attr,omitempty"`
}

type Objref struct {
	Hlink  string  `xml:"hlink,attr"`
	Priv   *bool   `xml:"priv,attr,omitempty"`
	Region *Region `xml:"region,omitempty"`
}

type Region struct {
	Corner1x *int `xml:"corner1_x,attr,omitempty"`
	Corner1y *int `xml:"corner1_y,attr,omitempty"`
	Corner2x *int `xml:"corner2_x,attr,omitempty"`
	Corner2y *int `xml:"corner2_y,attr,omitempty"`
}

type Placeref struct {
	Hlink string `xml:"hlink,attr"`
}

type LdsOrd struct {
	Priv        *bool         `xml:"priv,attr,omitempty"`
	Type        string        `xml:"type,attr"`
	Daterange   *Daterange    `xml:"daterange,omitempty"`
	Datespan    *Datespan     `xml:"datespan,omitempty"`
	Dateval     *Dateval      `xml:"dateval,omitempty"`
	Datestr     *Datestr      `xml:"datestr,omitempty"`
	Temple      *Temple       `xml:"temple,omitempty"`
	Place       *Place        `xml:"place,omitempty"`
	Status      *Status       `xml:"status,omitempty"`
	SealedTo    *SealedTo     `xml:"sealed_to,omitempty"`
	Noteref     []Noteref     `xml:"noteref,omitempty"`
	Citationref []Citationref `xml:"citationref,omitempty"`
}

type Temple struct {
	Val string `xml:"val,attr"`
}

type Status struct {
	Val string `xml:"val,attr"`
}

type SealedTo struct {
	Hlink string `xml:"hlink,attr"`
}
