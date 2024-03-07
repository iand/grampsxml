//go:build ignore

// run this using go run example.go

package main

import (
	"encoding/xml"
	"log"

	"github.com/iand/grampsxml"
)

func main() {
	input := `<?xml version="1.0" encoding="UTF-8"?>
<database xmlns="http://gramps-project.org/xml/1.7.1/">
  <people>
    <person handle="_076KQC7HG6P8BL5E35" change="1185438865" id="I0667">
      <gender>M</gender>
      <name type="Birth Name">
        <first>Edward</first>
        <surname>Потылицин</surname>
      </name>
    </person>
  </people>
</database>`

	var db grampsxml.Database
	err := xml.Unmarshal([]byte(input), &db)
	if err != nil {
		log.Fatalf("unexpected unmarshal error: %v", err)
	}

	for _, p := range db.People.Person {
		for _, name := range p.Name {
			log.Printf("%s: %s\n", *name.Type, *name.First)
		}
	}
}
