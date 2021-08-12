package schema

var gomodules = []string{
	"github.com/golang/mock/mockgen@v1.6.0",
	"github.com/go-kit/kit@v0.11.0",
	"github.com/gorilla/mux@v1.8.0",
	"github.com/google/uuid@v1.3.0",
	"github.com/go-yaml/yaml@v2",
	"github.com/oklog/run@v1.1.0",
	"github.com/go-gulfstream/tmpevents@latest",
	"github.com/go-gulfstream/gulfstream@latest",
	"github.com/go-kit/log@v0.1.0",
	"github.com/jackc/pgx/v4@v4.11.0",
	"github.com/go-redis/redis/v8@v8.11.0",
	"google.golang.org/grpc@v1.38.0",
}

func SetDefaultGoModules(manifest *Manifest) {
	manifest.GoGetPackages = append(manifest.GoGetPackages, gomodules...)
}

func SetGoModules(manifest *Manifest, modules []string) {
	manifest.GoGetPackages = append(manifest.GoGetPackages, modules...)
}
