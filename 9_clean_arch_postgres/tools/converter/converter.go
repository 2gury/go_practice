package converter

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

func AnyToBytesBuffer(i interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(i)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

func ReadBytes(r io.Reader) ([]byte) {
	bytes, _ := ioutil.ReadAll(r)
	return bytes
}

func AnyBytesToString(i interface{}) *strings.Reader {
	anyJson, _ := AnyToBytesBuffer(i)
	return strings.NewReader(anyJson.String())
}