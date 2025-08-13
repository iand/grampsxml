package grampsxml

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func ptr[T any](v T) *T { return &v }

type wellFormedCase struct {
	name  string // optional
	input string
	want  *Database
}

var wellFormedCases = []wellFormedCase{
	{
		input: `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE database PUBLIC "-//Gramps//DTD Gramps XML 1.7.1//EN"
"http://gramps-project.org/xml/1.7.1/grampsxml.dtd">
<database xmlns="http://gramps-project.org/xml/1.7.1/">
	<header>
		<created date="2024-03-04" version="5.1.5"/>
		<researcher>
			<resname>grampsxml tester name</resname>
			<resaddr>grampsxml tester addr</resaddr>
			<reslocality>grampsxml tester locality</reslocality>
			<rescity>grampsxml tester city</rescity>
			<resstate>grampsxml tester state</resstate>
			<rescountry>grampsxml tester country</rescountry>
			<respostal>grampsxml tester postal</respostal>
			<resphone>grampsxml tester phone</resphone>
			<resemail>grampsxml tester email</resemail>
		</researcher>
		<mediapath>/user/grampsxml/media</mediapath>
	</header>
</database>`,
		want: &Database{
			Header: Header{
				Created: Created{
					Date:    "2024-03-04",
					Version: "5.1.5",
				},
				Researcher: &Researcher{
					Resname:     ptr("grampsxml tester name"),
					Resaddr:     ptr("grampsxml tester addr"),
					Reslocality: ptr("grampsxml tester locality"),
					Rescity:     ptr("grampsxml tester city"),
					Resstate:    ptr("grampsxml tester state"),
					Rescountry:  ptr("grampsxml tester country"),
					Respostal:   ptr("grampsxml tester postal"),
					Resphone:    ptr("grampsxml tester phone"),
					Resemail:    ptr("grampsxml tester email"),
				},
				Mediapath: ptr("/user/grampsxml/media"),
			},
		},
	},

	{
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <name-formats>
    <format number="-1" name="Surname, Name|Common Suffix (Nickname)" fmt_str="surname, Name|common suffix (nickname)" active="1"/>
  </name-formats>
</database>`,
		want: &Database{
			NameFormats: []NameFormat{
				{
					Number: "-1",
					Name:   "Surname, Name|Common Suffix (Nickname)",
					Fmtstr: "surname, Name|common suffix (nickname)",
					Active: ptr(true),
				},
			},
		},
	},

	{
		name: "people",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <people>
    <person handle="_f8c683ce6176f7422e564f7262d" change="1709571921" id="I0000">
      <gender>M</gender>
      <name type="Birth Name" sort="2">
        <first>Testy</first>
        <call>Testy</call>
        <surname prefix="von" derivation="Inherited">Testface</surname>
        <surname prim="0" derivation="Patronymic">Testface</surname>
        <title>Mr</title>
        <nick>tt</nick>
        <daterange start="1920" stop="1928"/>
      </name>
    </person>
    <person handle="_076KQC7HG6P8BL5E35" change="1185438865" id="I0667">
      <gender>M</gender>
      <name type="Birth Name">
        <first>Edward</first>
        <surname>Потылицин</surname>
      </name>
      <eventref hlink="_a5af0ecd609332c4202" role="Primary"/>
      <eventref hlink="_a5af0ecd61d3ea3dda9" role="Primary"/>
      <eventref hlink="_a5af0ecd62a6af10f0c" role="Primary"/>
      <parentin hlink="_S66KQC0PH4U8544UYW"/>
      <childof hlink="_A54KQCIQUDO6Y2RR7I"/>
      <citationref hlink="_c140d245e167ee93109"/>
    </person>
  </people>
</database>`,
		want: &Database{
			People: &People{
				Person: []Person{
					{
						ID:     ptr("I0000"),
						Handle: "_f8c683ce6176f7422e564f7262d",
						Change: "1709571921",
						Priv:   nil,
						Gender: "M",
						Name: []Name{
							{
								Type:  ptr("Birth Name"),
								Sort:  ptr("2"),
								First: ptr("Testy"),
								Call:  ptr("Testy"),
								Title: ptr("Mr"),
								Nick:  ptr("tt"),
								Surname: []Surname{
									{
										Prefix:     ptr("von"),
										Prim:       nil,
										Derivation: ptr("Inherited"),
										Surname:    "Testface",
									},
									{
										Prim:       ptr(false),
										Derivation: ptr("Patronymic"),
										Surname:    "Testface",
									},
								},
								Daterange: &Daterange{
									Start:   "1920",
									Stop:    "1928",
									Quality: nil,
								},
							},
						},
					},
					{
						Handle: "_076KQC7HG6P8BL5E35",
						ID:     ptr("I0667"),
						Change: "1185438865",
						Gender: "M",
						Name: []Name{
							{
								Type:  ptr("Birth Name"),
								First: ptr("Edward"),
								Surname: []Surname{
									{
										Surname: "Потылицин",
									},
								},
							},
						},
						Eventref: []Eventref{
							{
								Hlink: "_a5af0ecd609332c4202",
								Role:  ptr("Primary"),
							},
							{
								Hlink: "_a5af0ecd61d3ea3dda9",
								Role:  ptr("Primary"),
							},
							{
								Hlink: "_a5af0ecd62a6af10f0c",
								Role:  ptr("Primary"),
							},
						},
						Parentin: []Parentin{
							{
								Hlink: "_S66KQC0PH4U8544UYW",
							},
						},
						Childof: []Childof{
							{
								Hlink: "_A54KQCIQUDO6Y2RR7I",
							},
						},
						Citationref: []Citationref{
							{
								Hlink: "_c140d245e167ee93109",
							},
						},
					},
				},
			},
		},
	},

	{
		name: "events",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <events>
    <event handle="_a5af0eb667015e355db" change="1284030602" id="E0000">
      <type>Birth</type>
      <dateval val="1987-08-29"/>
      <place hlink="_08TJQCCFIX31BXPNXN"/>
      <description>Birth of Warner, Sarah Suzanne</description>
    </event>
    <event handle="_a5af0eb696917232725" change="1284030602" id="E0001">
      <type>LVG</type>
      <description>Custom FTW5 tag to specify LIVING not specified in GEDCOM 5.5</description>
    </event>
  </events>
</database>`,
		want: &Database{
			Events: &Events{
				Event: []Event{
					{
						Handle:      "_a5af0eb667015e355db",
						Change:      "1284030602",
						ID:          ptr("E0000"),
						Type:        ptr("Birth"),
						Dateval:     &Dateval{Val: "1987-08-29"},
						Place:       &Place{Hlink: "_08TJQCCFIX31BXPNXN"},
						Description: ptr("Birth of Warner, Sarah Suzanne"),
					},
					{
						Handle:      "_a5af0eb696917232725",
						Change:      "1284030602",
						ID:          ptr("E0001"),
						Type:        ptr("LVG"),
						Description: ptr("Custom FTW5 tag to specify LIVING not specified in GEDCOM 5.5"),
					},
				},
			},
		},
	},

	{
		name: "families",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <families>
    <family handle="_X8YJQC77ZZBLP5KB2" change="1185438865" id="F0046">
      <rel type="Married"/>
      <father hlink="_L8YJQCNEX0EVZ9TG2L"/>
      <mother hlink="_B9YJQCL6OS9VRKOIB"/>
      <eventref hlink="_a5af0edac9c3b0f9ff0" role="Family"/>
      <childref hlink="_V9YJQCQB9O2BBIF7H3"/>
      <childref hlink="_CAYJQCKOL49OF7XWB3"/>
      <childref hlink="_SS1KQCWWF9488Q330U"/>
      <citationref hlink="_c140d2908f969bc2007"/>
    </family>
  </families>
</database>`,
		want: &Database{
			Families: &Families{
				Family: []Family{
					{
						Handle: "_X8YJQC77ZZBLP5KB2",
						Change: "1185438865",
						ID:     ptr("F0046"),
						Rel:    &Rel{Type: "Married"},
						Father: &Father{Hlink: "_L8YJQCNEX0EVZ9TG2L"},
						Mother: &Mother{Hlink: "_B9YJQCL6OS9VRKOIB"},
						Eventref: []Eventref{
							{Hlink: "_a5af0edac9c3b0f9ff0", Role: ptr("Family")},
						},
						Childref: []Childref{
							{Hlink: "_V9YJQCQB9O2BBIF7H3"},
							{Hlink: "_CAYJQCKOL49OF7XWB3"},
							{Hlink: "_SS1KQCWWF9488Q330U"},
						},
						Citationref: []Citationref{
							{Hlink: "_c140d2908f969bc2007"},
						},
					},
				},
			},
		},
	},

	{
		name: "citations",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <citations>
    <citation handle="_c140d2362f25a92643b" change="1328025930" id="C0000">
      <dateval val="1855-06-25"/>
      <confidence>3</confidence>
      <sourceref hlink="_b39fe3f390e30bd2b99"/>
    </citation>
  </citations>
</database>`,
		want: &Database{
			Citations: &Citations{
				Citation: []Citation{
					{
						Handle:     "_c140d2362f25a92643b",
						Change:     "1328025930",
						ID:         ptr("C0000"),
						Dateval:    &Dateval{Val: "1855-06-25"},
						Confidence: "3",
						Sourceref:  &Sourceref{Hlink: "_b39fe3f390e30bd2b99"},
					},
				},
			},
		},
	},

	{
		name: "sources",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <sources>
    <source handle="_VUBKMQTA2XZG1V6QP8" change="1185438865" id="S0002">
      <stitle>World of the Wierd</stitle>
      <sauthor>John Jacob Jinglehiemerschmitt</sauthor>
      <sabbrev>WOTW</sabbrev>
      <noteref hlink="_ac3804a8405171ef666"/>
      <srcattribute type="Book Cover Type" value="Paperback"/>
      <reporef hlink="_a701e99f93e5434f6f3" callno="what-321-ever" medium="Photo"/>
      <reporef hlink="_a701ead12841521cd4d" callno="nothing-0" medium="Manuscript"/>
    </source>
  </sources>
</database>`,
		want: &Database{
			Sources: &Sources{
				Source: []Source{
					{
						Handle:  "_VUBKMQTA2XZG1V6QP8",
						Change:  "1185438865",
						ID:      ptr("S0002"),
						Stitle:  ptr("World of the Wierd"),
						Sauthor: ptr("John Jacob Jinglehiemerschmitt"),
						Sabbrev: ptr("WOTW"),
						Noteref: []Noteref{
							{Hlink: "_ac3804a8405171ef666"},
						},
						Srcattribute: []Srcattribute{
							{Type: "Book Cover Type", Value: "Paperback"},
						},
						Reporef: []Reporef{
							{Hlink: "_a701e99f93e5434f6f3", Callno: ptr("what-321-ever"), Medium: ptr("Photo")},
							{Hlink: "_a701ead12841521cd4d", Callno: ptr("nothing-0"), Medium: ptr("Manuscript")},
						},
					},
				},
			},
		},
	},

	{
		name: "places",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <places>
    <placeobj handle="_00BKQC7SA8C9NCGB0A" change="1234390424" id="P0852" type="City">
      <ptitle>Deltona, FL</ptitle>
      <pname value="Deltona"/>
      <coord long="-81.2636738" lat="28.9005446"/>
      <placeref hlink="_c96587264d513cb56c90efab74d"/>
    </placeobj>
    <placeobj handle="_023LQCRXHOWC0N2PG5" change="1234390522" id="P0853" type="City">
      <ptitle>Spirit Lake, IA</ptitle>
      <pname value="Spirit Lake"/>
      <coord long="-95.1022169" lat="43.4221843"/>
      <placeref hlink="_c9658726cea1b7809e4fb1fa292"/>
    </placeobj>
  </places>
</database>`,
		want: &Database{
			Places: &Places{
				Place: []Placeobj{
					{
						Handle: "_00BKQC7SA8C9NCGB0A",
						Change: "1234390424",
						ID:     ptr("P0852"),
						Type:   "City",
						Ptitle: ptr("Deltona, FL"),
						Pname: []Pname{
							{Value: "Deltona"},
						},
						Coord: &Coord{Long: "-81.2636738", Lat: "28.9005446"},
						Placeref: []Placeref{
							{Hlink: "_c96587264d513cb56c90efab74d"},
						},
					},
					{
						Handle: "_023LQCRXHOWC0N2PG5",
						Change: "1234390522",
						ID:     ptr("P0853"),
						Type:   "City",
						Ptitle: ptr("Spirit Lake, IA"),
						Pname: []Pname{
							{Value: "Spirit Lake"},
						},
						Coord: &Coord{Long: "-95.1022169", Lat: "43.4221843"},
						Placeref: []Placeref{
							{Hlink: "_c9658726cea1b7809e4fb1fa292"},
						},
					},
				},
			},
		},
	},

	{
		name: "objects",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <objects>
    <object handle="_238CGQ939HG18SS5MG" change="1328027353" id="O0010">
      <file src="1897_expeditionsmannschaft_rio_a.jpg" mime="image/jpeg" checksum="352c7ae13b8b642471ecae6fa78ce206" description="1897_expeditionsmannschaft_rio_a"/>
      <dateval val="1897"/>
      <tagref hlink="_bb80c2b235b0a1b3f49"/>
    </object>
    <object handle="_78V2GQX2FKNSYQ3OHE" change="1185438865" id="O0009">
      <file src="Gunnlaugur_Larusson_-_Yawn.jpg" mime="image/jpeg" checksum="6bae0888ffdbad79b2735a5ac17450aa" description="Yawn"/>
    </object>
  </objects>
</database>`,
		want: &Database{
			Objects: &Objects{
				Object: []Object{
					{
						Handle: "_238CGQ939HG18SS5MG",
						Change: "1328027353",
						ID:     ptr("O0010"),
						File: File{
							Src:         "1897_expeditionsmannschaft_rio_a.jpg",
							Mime:        "image/jpeg",
							Checksum:    ptr("352c7ae13b8b642471ecae6fa78ce206"),
							Description: "1897_expeditionsmannschaft_rio_a",
						},
						Dateval: &Dateval{Val: "1897"},
						Tagref: []Tagref{
							{Hlink: "_bb80c2b235b0a1b3f49"},
						},
					},
					{
						Handle: "_78V2GQX2FKNSYQ3OHE",
						Change: "1185438865",
						ID:     ptr("O0009"),
						File: File{
							Src:         "Gunnlaugur_Larusson_-_Yawn.jpg",
							Mime:        "image/jpeg",
							Checksum:    ptr("6bae0888ffdbad79b2735a5ac17450aa"),
							Description: "Yawn",
						},
					},
				},
			},
		},
	},

	{
		name: "repositories",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <repositories>
    <repository handle="_a701e99f93e5434f6f3" change="1328027438" id="R0002">
      <rname>New York Public Library</rname>
      <type>Library</type>
      <address>
        <street>5th Ave at 42 street</street>
        <city>New York</city>
        <state>New York</state>
        <country>USA</country>
        <postal>11111</postal>
        <citationref hlink="_c140e0925ac0adcf8c4"/>
      </address>
    </repository>
  </repositories>
</database>`,
		want: &Database{
			Repositories: &Repositories{
				Repository: []Repository{
					{
						ID:     ptr("R0002"),
						Handle: "_a701e99f93e5434f6f3",
						Change: "1328027438",
						Rname:  "New York Public Library",
						Type:   "Library",
						Address: []Address{
							{
								Street:  ptr("5th Ave at 42 street"),
								City:    ptr("New York"),
								State:   ptr("New York"),
								Country: ptr("USA"),
								Postal:  ptr("11111"),
								Citationref: []Citationref{
									{Hlink: "_c140e0925ac0adcf8c4"},
								},
							},
						},
						// Number: "-1",
						// Name:   "Surname, Name|Common Suffix (Nickname)",
						// Fmtstr: "surname, Name|common suffix (nickname)",
						// Active: ptr(true),
					},
				},
			},
		},
	},

	{
		name: "notes",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <notes>
    <note handle="_ac380498bac48eedee8" change="1185438865" id="N0001" type="Name Note">
      <text>Names can notes, too. This is a note for the alternate name of Louse Garner for Lewis Anderson Garner.</text>
      <style name="link" value="gramps://Person/handle/GNUJQCL9MD64AM56OH">
        <range start="80" end="101"/>
      </style>
    </note>
    <note handle="_b39feb55e1173f4a699" change="1234371685" id="N0010" type="Source text" format="1">
      <text>1855-06-25

    line 1    fac secunda Junij Baptiza-
    line 2    tus est Lewis Anderson
    line 3    filius legitimus Guillielmus
    line 4    Garner et Elisabetha
    line 5    Becks. Susceptores fuerent
    line 6    petrus Arts et Catharina
    line 7    van de Voorde</text>
      <style name="bold">
        <range start="0" end="10"/>
      </style>
      <style name="underline">
        <range start="0" end="10"/>
      </style>
    </note>
  </notes>
</database>`,
		want: &Database{
			Notes: &Notes{
				Note: []Note{
					{
						ID:     ptr("N0001"),
						Handle: "_ac380498bac48eedee8",
						Change: "1185438865",
						Type:   "Name Note",
						Text:   "Names can notes, too. This is a note for the alternate name of Louse Garner for Lewis Anderson Garner.",
						Style: []Style{
							{
								Name:  "link",
								Value: ptr("gramps://Person/handle/GNUJQCL9MD64AM56OH"),
								Range: []Range{
									{Start: 80, End: 101},
								},
							},
						},
					},
					{
						ID:     ptr("N0010"),
						Handle: "_b39feb55e1173f4a699",
						Change: "1234371685",
						Format: ptr(true),
						Type:   "Source text",
						Text: `1855-06-25

    line 1    fac secunda Junij Baptiza-
    line 2    tus est Lewis Anderson
    line 3    filius legitimus Guillielmus
    line 4    Garner et Elisabetha
    line 5    Becks. Susceptores fuerent
    line 6    petrus Arts et Catharina
    line 7    van de Voorde`,

						Style: []Style{
							{
								Name: "bold",
								Range: []Range{
									{Start: 0, End: 10},
								},
							},
							{
								Name: "underline",
								Range: []Range{
									{Start: 0, End: 10},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		name: "bookmarks",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <bookmarks>
    <bookmark target="person" hlink="_AWFKQCJELLUWDY2PD3"/>
    <bookmark target="person" hlink="_35WJQC1B7T7NPV8OLV"/>
  </bookmarks>
</database>`,
		want: &Database{
			Bookmarks: &Bookmarks{
				Bookmark: []Bookmark{
					{Target: "person", Hlink: "_AWFKQCJELLUWDY2PD3"},
					{Target: "person", Hlink: "_35WJQC1B7T7NPV8OLV"},
				},
			},
		},
	},

	{
		name: "namemaps",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <namemaps>
    <map type="group_as" key="Fernández" value="Fernandez"/>
  </namemaps>
</database>`,
		want: &Database{
			Namemaps: &Namemaps{
				Map: []Map{
					{
						Type:  "group_as",
						Key:   "Fernández",
						Value: "Fernandez",
					},
				},
			},
		},
	},
}

func TestWellFormedCases(t *testing.T) {
	for _, tc := range wellFormedCases {
		t.Run(tc.name, func(t *testing.T) {
			var db Database
			err := xml.Unmarshal([]byte(tc.input), &db)
			if err != nil {
				t.Fatalf("unexpected unmarshal error: %v", err)
			}

			if diff := cmp.Diff(tc.want, &db); diff != "" {
				t.Errorf("unmarshal mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// func TestWellFormedRoundtrip(t *testing.T) {
// 	for _, tc := range wellFormedCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			var db Database
// 			err := xml.Unmarshal([]byte(tc.input), &db)
// 			if err != nil {
// 				t.Fatalf("unexpected unmarshal error: %v", err)
// 			}

// 			out, err := xml.MarshalIndent(&db, "", "\t")
// 			if err != nil {
// 				t.Fatalf("unexpected marshal error: %v", err)
// 			}

// 			if diff := cmp.Diff([]byte(tc.input), out); diff != "" {
// 				t.Errorf("unmarshal mismatch (-want +got):\n%s", diff)
// 			}
// 		})
// 	}
// }
