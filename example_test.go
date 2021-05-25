package androidbinary_test

import (
	"encoding/xml"
	"fmt"
	"os"

	"androidbinary"
	"androidbinary/apk"
)

func ExampleNewXMLFile() {
	f, _ := os.Open("testdata/AndroidManifest.xml")
	xmlFile, err := androidbinary.NewXMLFile(f)
	if err != nil {
		panic(err)
	}

	var v apk.Manifest
	dec := xml.NewDecoder(xmlFile.Reader())
	dec.Decode(&v)

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "\t")
	enc.Encode(v)

	// Output:

}

func ExampleNewTableFile() {
	f, err := os.Open("testdata/resources.arsc")
	if err != nil {
		panic(err)
	}
	tableFile, err := androidbinary.NewTableFile(f)
	if err != nil {
		panic(err)
	}

	val, err := tableFile.GetResource(0x7f040000, &androidbinary.ResTableConfig{})
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	// Output:
	// FireworksMeasure
}
