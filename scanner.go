package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// metadata of each document
type Doc struct {
	Rev  string
	Name string
	Date time.Time
}

// product contain docs of BOM, gerber, schematic, .. etc
type DocKey struct {
	Customer string
	Product  string
	DocType  string
}

// Scanner is use for scan folder and generate index file based on template
type Scanner struct {
	tpl      *template.Template
	outfile  string
	document map[DocKey]Doc
}

func NewScanner() (sc *Scanner, err error) {
	sc = new(Scanner)
	sc.outfile = "index.htm"
	sc.document = make(map[DocKey]Doc)
	sc.tpl, err = template.ParseFiles("index.tpl")
	if err != nil {
		return nil, err
	}
	return sc, nil
}

func (s *Scanner) visit(path string, info os.FileInfo, err error) error {
	x := strings.Split(path, string(os.PathSeparator))
	switch len(x) {
	case 5:
		x := strings.Split(path, string(os.PathSeparator))
		customer, product, doctype, rev := x[0], x[1], x[2], x[3]
		s.document[DocKey{customer, product, doctype}] = Doc{rev, info.Name(), info.ModTime()}
	}
	return nil
}

// scan root folder to generate documents data
func (s *Scanner) Scan() (err error) {
	filepath.Walk(".", s.visit)
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

	cust := []string{}
	for k, _ := range s.document {
		cust = append(cust, k.Customer)
	}
	fmt.Printf("%v", cust)
	//products []string
	//doctype  []string

	//execute template
	data := struct {
		LastUpdate string
		Doc        map[DocKey]Doc
	}{
		time.Now().Format(time.RFC850),
		s.document,
	}

	err = s.tpl.Execute(out, data)
	if err != nil {
		return err
	}
	return nil
}
