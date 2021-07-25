package goutil

import (
	"os/exec"
	"regexp"
)

const DefaultVersion = "1.17"
const pattern = `go(\d+\.\d+)`

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(pattern)
}

func GoInstall() bool {
	_, err := exec.Command(bin(), "version").CombinedOutput()
	if err != nil {
		return false
	}
	return true
}

func Version() string {
	out, err := exec.Command(bin(), "version").CombinedOutput()
	if err != nil {
		return DefaultVersion
	}
	return parse(string(out), DefaultVersion)
}

func parse(s string, def string) string {
	if len(s) == 0 {
		return def
	}
	matches := re.FindAllStringSubmatch(s, 1)
	if len(matches) != 1 {
		return def
	}
	return matches[0][1]
}
