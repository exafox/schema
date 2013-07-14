package schema

import (
	"encoding/json"
)

type DataType string

//Column data goes here
type Column struct {
	Name string
	Type DataType
}

// Source is a path reference to datafile
type Source string

//
type InfoDef struct{ Name, Description, Url string }

type Table struct {
	Sources []Source
	Columns map[string]Column
}

type Slice struct {
	Table      string
	Dimensions []string
	Metrics    []string
}

//implements https://github.com/cfpb/qu/wiki/Dataset-publishing-format
type Definition struct {
	Info   InfoDef
	Tables map[string]Table
	Slices map[string]Slice
}

// Makes an empty Definintion
func NewDefinition() *Definition {
	var definition = new(Definition)
	return definition
}

//Serialize a Definition to Json
func (d *Definition) ToJson() []byte {
	data, _ := json.Marshal(&d)
	return data
}

//Inflate a Definition from a json representation
func (d *Definition) FromJson(jsondata []byte) {
	json.Unmarshal(jsondata, &d)
}
