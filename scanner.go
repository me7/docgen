package main

import (
	"html/template"
	"os"
)

// metadata of each document
type Doc struct {
	Group string
	Url   string
	Rev   string
	Date  int
}

// product contain docs of BOM, gerber, schematic, .. etc
type Product struct {
	Name string
	Docs []Doc
}

// Scanner is use for scan folder and generate index file based on template
type Scanner struct {
	tpl     *template.Template
	outfile string
	Data    []Product
}

func NewScanner() (sc *Scanner, err error) {
	sc = new(Scanner)
	sc.outfile = "index.htm"
	sc.tpl, err = template.ParseFiles("index.tpl")
	if err != nil {
		return nil, err
	}
	return sc, nil
}

// scan root folder to generate documents data
func (s *Scanner) Scan() (err error) {
	ps100_docs := []Doc{
		{"BOM", "url", "A", 1},
		{"Gerber", "url", "B", 2}}
	ps100 := Product{
		Name: "PS100 - 48721",
		Docs: ps100_docs,
	}
	nissin_docs := []Doc{
		{"BOM", "url", "A", 1},
		{"Gerber", "url", "B", 2}}
	nissin := Product{
		Name: "Nissin - 48847",
		Docs: nissin_docs,
	}
	s.Data = []Product{ps100, nissin}
	return nil
}

// Generate index.htm file
func (s *Scanner) GenIndex() (err error) {
	// create output file
	out, err := os.Create(s.outfile)
	if err != nil {
		return err
	}
	defer out.Close()

	//execute template
	err = s.tpl.Execute(out, s.Data)
	if err != nil {
		return err
	}
	return nil
}
