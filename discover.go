package schema

import (
	"strconv"
	"time"
)

//Determines if a value can represent a boolean
func IsBoolean(v string) bool {
	_, err := strconv.ParseBool(v)
	return err == nil
}

//Determines if a value can represent an integer
func IsInteger(v string) bool {
	_, err := strconv.ParseInt(v, 10, 64)
	return err == nil
}

//Determines if a value matches any of the given date formats
func dateTimeMatchesFormats(v string, formats []string) bool {

	var is_dt = false

	for _, format := range formats {
		_, err := time.Parse(format, v)
		is_dt = is_dt || (err == nil)
	}
	return is_dt

}

//Determines if the value is a date
func IsDate(v string) bool {

	// fmt.Println("IsDate")
	var formats = []string{
		"01-02-2006",
		"01/02/2006",
		"2006-01-02",
		"01/02/2006",
		"1/2/2006",
	}
	return dateTimeMatchesFormats(v, formats)
}

//Determines if the value is a timestamp
func IsTime(v string) bool {

	// fmt.Println("IsTime")
	var formats = []string{
		"01/02/2006 03:04:05 PM -0700",
		"01/02/2006 03:04:05 PM -0700",
	}
	return dateTimeMatchesFormats(v, formats)

}

// Used to represent schema details
type Field struct {
	Name      string
	Position  int
	IsBoolean bool
	IsInteger bool
	IsDate    bool
	IsTime    bool
}

//Sets some default values on new Fields
func NewField(n string, p int) *Field {
	f := Field{
		Name:      n,
		Position:  p,
		IsBoolean: true,
		IsInteger: true,
		IsDate:    true,
		IsTime:    true,
	}
	return &f
}

//Send a series of values to this method to inflate Field with 
func (f *Field) Observe(v string) {
	// fmt.Println("Observing", v)
	f.IsBoolean = f.IsBoolean && IsBoolean(v)
	f.IsInteger = f.IsInteger && IsInteger(v)
	f.IsDate = f.IsDate && IsDate(v)
	f.IsTime = f.IsTime && IsTime(v)
}
