package schema

import (
	"encoding/csv"
	"fmt"
	"strings"
	"testing"
)

type Tests []struct {
	s    string
	want bool
}

func TestIsDate(t *testing.T) {
	// fmt.Println("Testing date discovery")
	var tests = Tests{
		{"07-09-1982", true},
		{"7/9/1982", true},
		{"1982-09-07", true},
	}

	for _, c := range tests {
		got := IsDate(c.s)
		if got != c.want {
			t.Errorf("IsDate(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}

func TestIsTime(t *testing.T) {
	// fmt.Println("Testing time discovery")
	var tests = Tests{
		{"01/02/2006 03:04:05 PM -0700", true},
		{"01/02/2006 03:04:05 PM -0700", true},
	}

	for _, c := range tests {
		got := IsTime(c.s)
		if got != c.want {
			t.Errorf("IsTime(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}

var testdata = `
Joined,"First Name",Second Name,Age
2009-10-03,John,Smith,67
2010-03-15,Jill,Tailor,54
`

func TestCsvDiscovery(t *testing.T) {
	r := csv.NewReader(strings.NewReader(testdata))
	row, err := r.Read()
	fields := make([]*Field, 0, 256)
	for pos := range row {
		f := NewField(row[pos], pos)
		fields = append(fields, f)
	}

	for err == nil {
		row, err = r.Read()
		for pos := range row {
			fields[pos].Observe(row[pos])
		}
	}
	//print ending header
	fmt.Println("\nFinished observing")
	for _, field := range fields {
		fmt.Println(field)
	}

}
