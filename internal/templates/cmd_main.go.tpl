package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
    "syscall"

	"{{$.GoModules}}/pkg/{{$.StreamPkgName}}"

	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
	"github.com/go-kit/log"
	"{{$.GoModules}}/internal/config"
    "{{$.GoModules}}/internal/stream"

	{{if $.StreamStorage.IsPostgres}}
	    storagepostgres "github.com/go-gulfstream/gulfstream/pkg/storage/postgres"
	   "github.com/jackc/pgx/v4/pgxpool"
	{{else if $.StreamStorage.IsRedis}}
	    "github.com/go-redis/redis/v8"
	    storageredis "github.com/go-gulfstream/gulfstream/pkg/storage/redis"
	{{end}}

	{{if $.StreamPublisher.IsKafka }}
        eventbuskafka "github.com/go-gulfstream/gulfstream/pkg/eventbus/kafka"
	{{end}}

	{{if $.CommandBus.IsGRPC -}}
	    commandbusgrpc "github.com/go-gulfstream/gulfstream/pkg/commandbus/grpc"
	    "google.golang.org/grpc"
	    "net"
    {{else if $.CommandBus.IsNATS -}}
    {{else if $.CommandBus.IsHTTP -}}
    {{end}}
)

func main() {
    cfg := loadConfig()
    ctx := context.Background()

    _ = cfg
	_ = ctx

   	var logger log.Logger
   	{
   		logger = log.NewLogfmtLogger(os.Stderr)
   		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
   		logger = log.With(logger, "caller", log.DefaultCaller)
   		logger = log.WithPrefix(logger, "stream", {{$.StreamPkgName}}.Name)
   	}

    {{if $.StreamStorage.IsDefault -}}
        storage := gulfstream.NewStorage({{$.StreamPkgName}}.Name, newEmptyStream)
    {{else if $.StreamStorage.IsPostgres -}}
        pool, err := pgxpool.Connect(ctx, "")
        if err != nil {
           _ = logger.Log("db", "postgres", "err", err)
           os.Exit(1)
        }
        defer pool.Close()
        {{if $.StreamStorage.EnableJournal -}}
        storage := storagepostgres.New(pool, {{$.StreamPkgName}}.Name, newEmptyStream, storagepostgres.WithJournal())
        {{else -}}
        storage := storagepostgres.New(pool, {{$.StreamPkgName}}.Name, newEmptyStream)
        {{end -}}
    {{else if $.StreamStorage.IsRedis -}}
        rdb := redis.NewClient(&redis.Options{Addr: ""})
        if err := rdb.Ping(ctx).Err(); err != nil && err != redis.Nil {
        	_ = logger.Log("db", "redis", "method", "ping", "err", err)
            os.Exit(1)
        }
        defer rdb.Close()
        storage := storageredis.New(rdb, {{$.StreamPkgName}}.Name, newEmptyStream)
    {{end -}}
    {{if $.StreamPublisher.IsKafka -}}
        publisher := eventbuskafka.NewPublisher([]string{}, eventbuskafka.DefaultConfig())
    {{else if $.StreamPublisher.IsConnector -}}
        publisher := gulfstream.NewConnectorPublisher()
    {{end -}}

    controller := gulfstream.NewMutator(storage, publisher)
    commandMutations := stream.NewCommandMutation( /* deps */ )
    eventMutations := stream.NewEventMutation( /* deps */ )

    stream.MakeCommandControllers(commandMutations, controller)
    stream.MakeEventControllers(eventMutations, controller)

    var g group.Group
    {{if $.CommandBus.IsGRPC -}}
    {
                grpcListener, err := net.ListenTCP("tcp", nil)
        		if err != nil {
        			_ = logger.Log("transport", "gRPC", "during", "Listen", "err", err)
        			os.Exit(1)
        		}
        		grpcServer := grpc.NewServer(/* interceptors */)
        		g.Add(func() error {
        			_ = logger.Log("transport", "gRPC", "addr", "")
        			commandBus := commandbusgrpc.NewServer(controller,
        				commandbusgrpc.WithServerErrorHandler(
        					func(err error) {
        						_ = logger.Log("transport", "commandbus/GRPC", "err", err)
        					}))
        			commandBus.Register(grpcServer)
        			return grpcServer.Serve(grpcListener)
        		}, func(error) {
        			grpcServer.GracefulStop()
        			_ = grpcListener.Close()
        		})
    }
    {{else if $.CommandBus.IsNATS -}}
    {

    }
    {{else if $.CommandBus.IsHTTP -}}
    {

    }
    {{end -}}
    {
    		cancelInterrupt := make(chan struct{})
    		g.Add(func() error {
    			c := make(chan os.Signal, 1)
    			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
    			select {
    			case sig := <-c:
    				return fmt.Errorf("received signal %s", sig)
    			case <-cancelInterrupt:
    				return nil
    			}
    		}, func(err error) {
    			close(cancelInterrupt)
    		})
    }

    logger.Log("exit", g.Run())
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
	return gulfstream.Blank({{$.StreamPkgName}}.Name, stream.New())
}