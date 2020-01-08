package main

import (
	"fmt"
	"unicode"

	"beneficiofacil.gopkg.net/site/util"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	mystring, err := util.ToUTF8("Damião", true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T: %#v", mystring, mystring)
}

func normalize(strangerThing string) string {
	b := make([]byte, len(strangerThing))
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	_, _, e := t.Transform(b, []byte(strangerThing), true)
	if e != nil {
		panic(e)
	}
	return string(b)
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
