package commands

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-gulfstream/gs/internal/goutil"

	"github.com/go-gulfstream/gs/internal/schema"
	"github.com/spf13/cobra"
)

func New() (*cobra.Command, error) {
	root := &cobra.Command{
		Use:   "gs",
		Short: "Standard Tooling for Go-Gulfstream Development",
	}

	root.AddCommand(initCommand())
	root.AddCommand(manifestCommand())
	root.AddCommand(applyCommand())
	root.AddCommand(addCommand())

	return root, nil
}

func loadManifestFromFile(projectPath string) (*schema.Manifest, error) {
	manifestFile := filepath.Join(projectPath, manifestFilename)
	data, err := ioutil.ReadFile(manifestFile)
	if err != nil {
		return nil, err
	}
	manifest, err := schema.DecodeManifest(data)
	if err != nil {
		return nil, err
	}
	schema.SanitizeManifest(manifest)
	if err := schema.ValidateManifest(manifest); err != nil {
		return nil, err
	}
	return manifest, nil
}

func writeManifestToFile(path string, manifest *schema.Manifest, force bool) error {
	data, err := schema.EncodeManifest(manifest)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(data)
	buf.WriteString("\n# available storage adapters:\n")
	for id, adapter := range schema.StorageAdapters {
		buf.WriteString(fmt.Sprintf("# id:%d, name: %s\n", id, adapter))
	}
	buf.WriteString("\n# available publisher adapters:\n")
	for id, adapter := range schema.PublisherAdapters {
		buf.WriteString(fmt.Sprintf("# id:%d, name: %s\n", id, adapter))
	}
	manifestFile := filepath.Join(path, manifestFilename)
	if _, err := os.Stat(manifestFile); err == nil && !force {
		return fmt.Errorf("manifest file already exists")
	}
	return ioutil.WriteFile(manifestFile, buf.Bytes(), 0755)
}

func filterDotFiles(files []fs.FileInfo) []fs.FileInfo {
	filtered := make([]fs.FileInfo, 0)
	for _, fi := range files {
		if strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		filtered = append(filtered, fi)
	}
	return filtered
}

func printManifest(m *schema.Manifest) {
	data, _ := schema.EncodeManifest(m)
	fmt.Printf("\nManifest:\n%s\n", string(data))
}

func countProjectFiles(projectPath string) int {
	files, err := ioutil.ReadDir(projectPath)
	if err != nil {
		return 0
	}
	return len(filterDotFiles(files))
}

func runGoTools(path string, packages []string) {
	if !goutil.GoInstall() {
		return
	}

	fmt.Println("===============================================")
	if len(packages) > 0 {
		fmt.Printf("go get:\n")
		for _, pkg := range packages {
			out, err := goutil.Get(path, pkg)
			if err != nil {
				fmt.Printf("%s - %s\n", redColor("[ERR]"), err)
				return
			}
			fmt.Printf("%s - %s %s\n", greenColor("[OK]"), pkg, string(out))
		}
	}

	fmt.Printf("go mod download:\n")
	out, err := goutil.Mod(path)
	if err != nil {
		fmt.Printf("%s - %s\n", redColor("[ERR]"), err)
		return
	}
	fmt.Printf("%s - %s\n", greenColor("[OK]"), string(out))

	fmt.Printf("go mod tidy:\n")
	out, err = goutil.Tidy(path)
	if err != nil {
		fmt.Printf("%s - %s\n", redColor("[ERR]"), err)
		return
	}
	fmt.Printf("%s - %s\n", greenColor("[OK]"), string(out))

	fmt.Printf("go test ./...:\n")
	out, err = goutil.Test(path)
	if err != nil {
		fmt.Printf("%s - %s\n", redColor("[ERR]"), err)
		return
	}
	fmt.Printf("%s - %s\n", greenColor("[OK]"), string(out))
}

func runGoTestAndGet(path string, packages []string) {
	if !goutil.GoInstall() {
		return
	}
	fmt.Println("===============================================")
	if len(packages) > 0 {
		fmt.Printf("go get:\n")
		for _, pkg := range packages {
			out, err := goutil.Get(path, pkg)
			if err != nil {
				fmt.Printf("%s - %s\n", redColor("[ERR]"), err)
				return
			}
			fmt.Printf("%s - %s %s\n", greenColor("[OK]"), pkg, string(out))
		}
	}
	fmt.Printf("go test ./...:\n")
	out, err := goutil.Test(path)
	if err != nil {
		fmt.Printf("%s - %s\n", redColor("[ERR]"), err)
		return
	}
	fmt.Printf("%s - %s\n", greenColor("[OK]"), string(out))
}
