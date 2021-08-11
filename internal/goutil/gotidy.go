package goutil

import (
	"os"
	"os/exec"
)

func Tidy(path string) ([]byte, error) {
	oldPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = os.Chdir(oldPath)
	}()
	if err := os.Chdir(path); err != nil {
		return nil, err
	}
	return exec.Command(bin(), "mod", "tidy").CombinedOutput()
}
