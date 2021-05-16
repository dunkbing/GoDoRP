package data

import (
	"encoding/json"
	"io"
)

// serializes the given interface into a string based JSON format
func ToJson(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJson(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
