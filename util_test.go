package swiftpass

import (
	"encoding/xml"
	"testing"
)

func assertEqual(t testing.TB, excepct, acquired string) {
	if excepct == acquired {
		return
	}
	t.Fatalf("Excepct: %s, acquired: %s\n", excepct, acquired)
}

func TestXmlObject2Map(t *testing.T) {
	type Object struct {
		XMLName xml.Name
		F1      string
		F2      string `xml:"f2"`
		F3      int
		F4      float64
		F5      string `xml:",omitempty"`
		F6      string `xml:"-"`
		f7      string
	}
	obj := Object{
		XMLName: xml.Name{
			Local: "obj",
		},
		F1: "1",
		F2: "2",
		F3: 3,
		F4: 4.1,
		F5: "5",
		F6: "6",
		f7: "7",
	}
	m := XmlObject2Map(obj)
	assertEqual(t, "1", m["F1"])
	assertEqual(t, "2", m["f2"])
	assertEqual(t, "3", m["F3"])
	assertEqual(t, "4.1", m["F4"])
	assertEqual(t, "5", m["F5"])
	assertEqual(t, "", m["F6"])
	assertEqual(t, "", m["f7"])
}
