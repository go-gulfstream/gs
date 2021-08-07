package source

import (
	"bytes"
	"go/printer"
	"io/ioutil"

	"github.com/dave/dst/decorator"

	"github.com/dave/dst"
)

var modifiedFiles = map[string]*dst.File{}

func FromFile(filename string) (*dst.File, error) {
	file, found := modifiedFiles[filename]
	if found {
		return file, nil
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	file, err = decorator.Parse(data)
	if err != nil {
		return nil, err
	}
	modifiedFiles[filename] = file
	return file, nil
}

func FlushToDisk() error {
	buf := &bytes.Buffer{}
	for filename, src := range modifiedFiles {
		buf.Reset()

		fset, asrc, err := decorator.RestoreFile(src)
		if err != nil {
			return err
		}

		if err := printer.Fprint(buf, fset, asrc); err != nil {
			return err
		}

		data, err := Format(filename, buf.Bytes())
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(filename, data, 0755); err != nil {
			return err
		}
	}
	return nil
}
