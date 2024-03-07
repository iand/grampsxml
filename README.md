# grampsxml

Go package to parse Gramps XML files.

[![Test Status](https://github.com/iand/grampsxml/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/iand/grampsxml/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/iand/grampsxml)](https://goreportcard.com/report/github.com/iand/grampsxml)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/iand/grampsxml)

## About

This package provides some Go types that can be used to unmarshal the native XML data exported by [Gramps](https://gramps-project.org/), an application for managing genealogical data.

Currently it supports [version 1.7.1](https://gramps-project.org/xml/1.7.1/) of the XML format.

## Status

This is the first draft of the XML type mapping and there are probably mistakes or oversights. Use with caution.

## Usage


```Go
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
```


## Getting Started

Simply run

	go get github.com/iand/gramps

Documentation is at [http://godoc.org/github.com/iand/gramps](http://godoc.org/github.com/iand/gramps)

## License

This is free and unencumbered software released into the public domain. For more
information, see <http://unlicense.org/> or the accompanying [`UNLICENSE`](UNLICENSE) file.
