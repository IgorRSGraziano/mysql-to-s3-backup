package services

import (
	"os/exec"
	"strings"
)

type Dump struct {
	GenerateCmd string
	FilePath    string
}

func NewDump(generateCmd, filePath string) *Dump {
	return &Dump{
		GenerateCmd: generateCmd,
		FilePath:    filePath,
	}
}

func (d *Dump) Generate() error {
	preparedCmd, args := prepareCommand(d.GenerateCmd)
	cmd := exec.Command(preparedCmd, args...).Run()
	return cmd
}

func prepareCommand(cmd string) (exec string, args []string) {
	parts := strings.Split(cmd, " ")
	return parts[0], parts[1:]
}
