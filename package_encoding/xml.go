package package_encoding

import (
	"encoding/xml"
	"fmt"
)

type Buku struct {
	Judul string `xml:"judul"`
	Penulis string `xml:"penulis"`
}

// DemoXML mendemokan encoding dan decoding XML
func DemoXML() {
	b := Buku{"Belajar Go", "Andi"}
	xmlBytes, err := xml.Marshal(b)
	if err != nil {
		fmt.Println("XML Marshal error:", err)
		return
	}
	fmt.Println("XML Encoded:", string(xmlBytes))

	var b2 Buku
	err = xml.Unmarshal(xmlBytes, &b2)
	if err != nil {
		fmt.Println("XML Unmarshal error:", err)
		return
	}
	fmt.Printf("XML Decoded: %+v\n", b2)
}
