//go:build gramps_schema180

package grampsxml

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var wellFormedCases180 = []wellFormedCase{
	{
		name: "dnatests",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.8.0/">
  <dnatests>
    <dnatest handle="_T1ABCDEF1234567890" change="1700000000" id="D0001">
      <person hlink="_P1ABCDEF1234567890"/>
      <account_name>tester@example.com</account_name>
      <provider>Ancestry</provider>
      <kit_id>KIT123456</kit_id>
      <test_type>Autosomal</test_type>
      <genome_build>GRCh37</genome_build>
      <dateval val="2023-01-15"/>
      <haplogroup>H1a1</haplogroup>
      <noteref hlink="_N1ABCDEF1234567890"/>
      <tagref hlink="_TAGABCDEF1234567890"/>
    </dnatest>
  </dnatests>
</database>`,
		want: &Database{
			DNATests: &DNATests{
				DNATest: []DNATest{
					{
						Handle:      "_T1ABCDEF1234567890",
						Change:      "1700000000",
						ID:          new("D0001"),
						Person:      &PersonLink{Hlink: "_P1ABCDEF1234567890"},
						AccountName: new("tester@example.com"),
						Provider:    new("Ancestry"),
						KitID:       new("KIT123456"),
						TestType:    new("Autosomal"),
						GenomeBuild: new("GRCh37"),
						Dateval:     &Dateval{Val: "2023-01-15"},
						Haplogroup:  new("H1a1"),
						Noteref:     []Noteref{{Hlink: "_N1ABCDEF1234567890"}},
						Tagref:      []Tagref{{Hlink: "_TAGABCDEF1234567890"}},
					},
				},
			},
		},
	},
	{
		name: "dnamatches",
		input: `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.8.0/">
  <dnamatches>
    <dnamatch handle="_M1ABCDEF1234567890" change="1700000001" id="DM0001">
      <subject_test hlink="_T1ABCDEF1234567890"/>
      <match_test hlink="_T2ABCDEF1234567890"/>
      <shared_cm val="142.3"/>
      <percent_shared val="0.0195"/>
      <segment_count val="7"/>
      <largest_segment_cm val="38.2"/>
      <predicted_relationship>3rd cousin</predicted_relationship>
      <predicted_generations val="4.1"/>
      <shared_ancestor confidence="1">
        <description>Great-great-grandmother</description>
        <person hlink="_P2ABCDEF1234567890"/>
        <noteref hlink="_N2ABCDEF1234567890"/>
      </shared_ancestor>
      <dna_segment chromosome="3" start_bp="45000000" end_bp="83000000" shared_cm="38.2" snp_count="1203" phase="2" ibd_state="1" start_rsid="rs123456" end_rsid="rs654321"/>
      <dna_segment chromosome="7" start_bp="12000000" end_bp="30000000" shared_cm="22.1" snp_count="678" phase="0"/>
      <citationref hlink="_C1ABCDEF1234567890"/>
      <tagref hlink="_TAGABCDEF1234567891"/>
    </dnamatch>
  </dnamatches>
</database>`,
		want: &Database{
			DNAMatches: &DNAMatches{
				DNAMatch: []DNAMatch{
					{
						Handle:                "_M1ABCDEF1234567890",
						Change:                "1700000001",
						ID:                    new("DM0001"),
						SubjectTest:           &SubjectTest{Hlink: "_T1ABCDEF1234567890"},
						MatchTest:             &MatchTest{Hlink: "_T2ABCDEF1234567890"},
						SharedCM:              &SharedCM{Val: 142.3},
						PercentShared:         &PercentShared{Val: 0.0195},
						SegmentCount:          &SegmentCount{Val: 7},
						LargestSegmentCM:      &LargestSegmentCM{Val: 38.2},
						PredictedRelationship: new("3rd cousin"),
						PredictedGenerations:  &PredictedGenerations{Val: 4.1},
						SharedAncestor: []SharedAncestor{
							{
								Confidence:  "1",
								Description: new("Great-great-grandmother"),
								Person:      &PersonLink{Hlink: "_P2ABCDEF1234567890"},
								Noteref:     []Noteref{{Hlink: "_N2ABCDEF1234567890"}},
							},
						},
						DNASegment: []DNASegment{
							{
								Chromosome: "3",
								StartBP:    45000000,
								EndBP:      83000000,
								SharedCM:   38.2,
								SNPCount:   1203,
								Phase:      2,
								IBDState:   new(1),
								StartRSID:  new("rs123456"),
								EndRSID:    new("rs654321"),
							},
							{
								Chromosome: "7",
								StartBP:    12000000,
								EndBP:      30000000,
								SharedCM:   22.1,
								SNPCount:   678,
								Phase:      0,
							},
						},
						Citationref: []Citationref{{Hlink: "_C1ABCDEF1234567890"}},
						Tagref:      []Tagref{{Hlink: "_TAGABCDEF1234567891"}},
					},
				},
			},
		},
	},
}

func TestWellFormedCases180(t *testing.T) {
	for _, tc := range wellFormedCases180 {
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
