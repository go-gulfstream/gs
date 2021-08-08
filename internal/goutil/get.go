package goutil

import (
	"os"
	"os/exec"
)

func Get(path string, pkg string) ([]byte, error) {
	oldPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	defer func() {
		os.Chdir(oldPath)
	}()
	if err := os.Chdir(path); err != nil {
		return nil, err
	}
	return exec.Command(bin(), "get", pkg).CombinedOutput()
}
