package jsonexraw

import (
	"bytes"

	json "github.com/yaegashi/jsonex.go/v1"
)

func Marshal(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(v)
	out := bytes.TrimRight(buf.Bytes(), "\n")
	return out, err
}

func MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent(prefix, indent)
	err := enc.Encode(v)
	out := bytes.TrimRight(buf.Bytes(), "\n")
	return out, err
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
