package commands

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-gulfstream/gs/internal/pkgdev"

	"github.com/rogpeppe/go-internal/modfile"

	"github.com/spf13/cobra"
)

func modCommand() *cobra.Command {
	var flags addFlags
	command := &cobra.Command{
		Use:   "mod [PATH]",
		Short: "Check go modules dependencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateModCommandArgs(args); err != nil {
				return err
			}
			drawBanner()
			return runModCommand(args[0])
		},
	}
	command.Flags().BoolVarP(&flags.Apply, "apply", "a", false, "add and apply changes to the project")
	return command
}

func validateModCommandArgs(args []string) error {
	lenArgs := len(args)
	if lenArgs != 1 {
		return fmt.Errorf("invalid number of arguments. got %d, expected 1", lenArgs)
	}
	if _, err := os.Stat(args[0]); err != nil {
		return err
	}
	return nil
}

func runModCommand(path string) error {
	modfilename := filepath.Join(path, "go.mod")
	data, err := ioutil.ReadFile(modfilename)
	if err != nil {
		return err
	}
	mod, err := modfile.Parse(modfilename, data, nil)
	if err != nil {
		return err
	}
	for _, mf := range mod.Require {
		ver, _, err := pkgdev.Version(mf.Mod.Path)
		if err != nil {
			if errors.Is(err, pkgdev.ErrLoadPage) {
				return err
			}
			fmt.Printf("%s - %v\n", err, redColor("[ERR]"))
			continue
		} else {
			fmt.Printf("%s - %s@%s => %s\n",
				greenColor("[OK]"),
				mf.Mod.Path,
				mf.Mod.Version,
				versionFormat(mf.Mod.Version, ver))
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func versionFormat(curVer string, versions []string) string {
	results := make([]string, len(versions))
	var found bool
	for i, v := range versions {
		if v == curVer {
			results[i] = boldStyle(greenColor(v))
			found = true
		} else {
			results[i] = v
		}
	}
	format := strings.Join(results, ",")
	if !found {
		format = yellowColor("NOT FOUND => [" + format + "]")
	} else {
		format = "[" + format + "]"
	}
	return format
}
