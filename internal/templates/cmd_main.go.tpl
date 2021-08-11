package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"{{$.GoModules}}/pkg/{{$.StreamPkgName}}"

	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
	"github.com/go-kit/log"
	"{{$.GoModules}}/internal/config"

	{{if $.StreamStorage.IsPostgres}}
	    storagepostgres "github.com/go-gulfstream/gulfstream/pkg/storage/postgres"
	   "github.com/jackc/pgx/v4/pgxpool"
	{{else if $.StreamStorage.IsRedis}}
	    "github.com/go-redis/redis/v8"
	    storageredis "github.com/go-gulfstream/gulfstream/pkg/storage/redis"
	{{end}}
)

func main() {
    cfg := loadConfig()
    ctx := context.Background()

   	var logger log.Logger
   	{
   		logger = log.NewLogfmtLogger(os.Stderr)
   		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
   		logger = log.With(logger, "caller", log.DefaultCaller)
   		logger = log.WithPrefix(logger, "stream", {{$.StreamPkgName}}.Name)
   	}

    {{if $.StreamStorage.IsDefault}}
        storage := gulfstream.NewStorage({{$.StreamPkgName}}.Name, newEmptyStream)
    {{else if $.StreamStorage.IsPostgres}}
        pool, err := pgxpool.Connect(ctx, "")
        if err != nil {
           _ = logger.Log("db", "postgres", "err", err)
           os.Exit(1)
        }
        defer pool.Close()
        {{if $.StreamStorage.EnableJournal}}
        storage := storagepostgres.New(pool, {{$.StreamPkgName}}.Name, newEmptyStream, storagepostgres.WithJournal())
        {{else -}}
        storage := storagepostgres.New(pool, {{$.StreamPkgName}}.Name, newEmptyStream)
        {{end}}
    {{else if $.StreamStorage.IsRedis}}
        rdb := redis.NewClient(&redis.Options{Addr: ""})
        if err := rdb.Ping(ctx).Err(); err != nil && err != redis.Nil {
        	_ = logger.Log("db", "redis", "method", "ping", "err", err)
            os.Exit(1)
        }
        defer rdb.Close()
        storage := storageredis.New(rdb, {{$.StreamPkgName}}.Name, newEmptyStream)
    {{end}}

   	_ = logger
    _ = ctx
    _ = cfg
    _ = storage
}

func loadConfig() *config.Stream {
	filename := flag.String("config", os.Getenv("CONFIG_FILE"), "path to configuration file")
	flag.Parse()

	if len(*filename) == 0 {
		fmt.Println("configuration file not found")
		os.Exit(1)
	}

	cfg, err := config.ParseStream(*filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return cfg
}

func newEmptyStream() *gulfstream.Stream {
	return nil
}