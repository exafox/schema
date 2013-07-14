package schema

import (
	"bytes"
	"os"
	"testing"
)

//Test that Definitions can be initialized, serialized, and desreialized
func TestDefinitionJson(t *testing.T) {

	var definition = &Definition{
		Info: InfoDef{"Name", "Description", "URL"},
		Tables: map[string]Table{
			"The Table": Table{
				Sources: []Source{
					"source1",
					"source2",
				},

				Columns: map[string]Column{
					"LOL Name": {"LOLname", "YAYtype"},
				},
			},
		},

		Slices: map[string]Slice{
			"whats_a_slice": Slice{
				Table:      "the_table",
				Dimensions: []string{"LolName1"},
				Metrics:    []string{"LolName2"},
			},
		},
	}

	jsondata := definition.ToJson()

	definition2 := NewDefinition()
	definition2.FromJson(jsondata)
	jsondata2 := definition2.ToJson()
	os.Stdout.Write(jsondata)

	if !bytes.Equal(jsondata, jsondata2) {
		t.Errorf("FromJson or ToJson not bit accurate")
	}

}
