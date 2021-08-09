package schema

var gomodules = []string{
	"github.com/golang/mock/mockgen@v1.6.0",
}

func SetDefaultGoModules(manifest *Manifest) {
	manifest.GoGetPackages = append(manifest.GoGetPackages, gomodules...)
}

func SetGoModules(manifest *Manifest, modules []string) {
	manifest.GoGetPackages = append(manifest.GoGetPackages, modules...)
}
