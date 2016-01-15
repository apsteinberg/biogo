package gff

import (
	"os"
	"testing"
)

func TestReadAll(t *testing.T) {
	// read example.gff
	f, err := os.Open("example.gff")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		t.Error(err)
	}
	expectNum := 6
	if len(records) != expectNum {
		t.Errorf("Expect %d, got %d", expectNum, len(records))
	}

}
