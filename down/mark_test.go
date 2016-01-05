package down

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestMark(t *testing.T) {
	data, err := ioutil.ReadFile("fixture/sample.md")
	if err != nil {
		t.Fatal(err)
	}
	buf := &bytes.Buffer{}
	err = Mark(buf, bytes.NewReader(data), true, true, "prettity")
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("fixture/sample.html", buf.Bytes(), 0600)
}
