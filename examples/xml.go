package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Nama    string   `xml:"nama"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Nama, p.Origin)
}

func main() {
	coffe := &Plant{Id: 27, Nama: ""}
	coffe.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffe, " ", " ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Nama: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffe, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(out))
}
