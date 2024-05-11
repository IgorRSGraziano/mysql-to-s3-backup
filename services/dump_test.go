package services

import "testing"

func Test_GenerateDump(t *testing.T) {
	dump := NewDump("echo", "test.txt")

	err := dump.Generate()

	if err != nil {
		t.Error("Expected nil, got", err)
	}
}

func Test_GenerateDumpWithArgs(t *testing.T) {
	dump := NewDump("echo args", "test.txt")

	err := dump.Generate()

	if err != nil {
		t.Error("Expected nil, got", err)
	}
}
