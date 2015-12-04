package main

import (
	"encoding/csv"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var output *csv.Writer

func main() {
	flag.Parse()
	root := flag.Arg(0)
	output = csv.NewWriter(os.Stdout)
	defer output.Flush()
	filepath.Walk(root, visit)
}

type testsuite struct {
	XMLName   xml.Name   `xml:"testsuite"`
	Testcases []testcase `xml:"testcase"`
}

type testcase struct {
	Name      string `xml:"name,attr"`
	Classname string `xml:"classname,attr"`
	Time      string `xml:"time,attr"`
}

func visit(path string, f os.FileInfo, err error) error {

	if f.IsDir() {
		return nil
	}

	if strings.HasPrefix(f.Name(), "TEST") && strings.HasSuffix(f.Name(), ".xml") {

		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not read file: %v\n", err)
			return err
		}

		t := testsuite{}

		errXML := xml.Unmarshal(bytes, &t)
		if errXML != nil {
			fmt.Fprintf(os.Stderr, "Can not unmarshal file: %v\n", errXML)
			return err
		}

		for _, tc := range t.Testcases {
			output.Write([]string{tc.Classname, tc.Name, tc.Time})
		}
	} else {
		fmt.Fprintf(os.Stderr, "Not a xml file: %s\n", path)
	}

	return nil
}
