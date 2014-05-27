package form

import (
	"net/url"
	"testing"
)

type Str struct {
	A string
	B int64
	C float64
	D bool
	E string  `form:"vare"`
	F int64   `form:"varf"`
	G float64 `form:"varg"`
	H bool    `form:"varh"`
}

var unmarshalTests = []struct {
	in  string
	out Str
}{
	{"http://localhost/test/?a=asdf&b=42&c=42.42&d=true&e=ffff&vare=asdf&varf=42&varg=42.42&varh=true", Str{"asdf", 42, 42.42, true, "asdf", 42, 42.42, true}},
}

func TestUnmarshal(t *testing.T) {
	for i, tt := range unmarshalTests {
		uri, _ := url.Parse(tt.in)
		query, _ := url.ParseQuery(uri.RawQuery)

		s := Str{}
		Unmarshal(query, &s)

		if s != tt.out {
			t.Errorf("%d. Unmarshal(%s) => %+v returned, expected %+v", i, uri.RawQuery, s, tt.out)
		}
	}
}
