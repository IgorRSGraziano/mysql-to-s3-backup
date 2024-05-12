package services

import (
	"os"
	"testing"
)

func Test_GenerateDumpWithArgs(t *testing.T) {
	dump := NewDump("echo test", os.TempDir())

	err := dump.Generate()

	if err != nil {
		t.Error("Expected nil, got", err)
	}
}
