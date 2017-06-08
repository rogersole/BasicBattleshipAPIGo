package handler

import (
	"testing"
	"encoding/json"
	"bytes"
)

func TestJsonDecoder(t *testing.T) {
	body := []byte(`[[4, 8],[3, 2],[8, 4],[0,9]]`)
	r := bytes.NewReader(body)
	_, err := json.Marshal(r)
	if err != nil {
		t.Error(err)
	}
}
