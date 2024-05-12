package services

import (
	"testing"
)

func Test_GenerateDump(t *testing.T) {
	dump := NewDump("echo test")

	err := dump.GenerateDumpFile()

	if err != nil {
		t.Error("Expected nil, got", err)
	}

	err = dump.DeleteDumpFile()

	if err != nil {
		t.Error("Expected nil, got", err)
	}
}
