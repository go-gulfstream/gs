package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"{{$.GoModules}}/pkg/{{$.StreamPkgName}}"

	"github.com/go-gulfstream/gulfstream/pkg/event"
	eventbuskafka "github.com/go-gulfstream/gulfstream/pkg/eventbus/kafka"

	"github.com/oklog/oklog/pkg/group"

	"{{$.GoModules}}/internal/config"

	"{{$.GoModules}}/internal/api"

	"{{$.GoModules}}/internal/projection"

	"github.com/go-kit/log"
)

func main() {
	cfg := loadConfig()
	ctx := context.Background()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
		logger = log.WithPrefix(logger, "projection", {{$.StreamPkgName}}.Name)
	}

	storage := projection.NewStorage()
	handler := projection.New(storage)
	controller := projection.NewController(handler)

	var service api.Service
	{
		service = api.NewService(storage)
		service = api.LoggingMiddleware(logger)(service)
	}

	httpHandler := api.MakeHTTPHandler(service, logger)

	var g group.Group
	{
		// The debug listener mounts the http.DefaultServeMux, and serves up
		// stuff like the Prometheus metrics route, the Go debug and profiling
		// routes, and so on.
		debugListener, err := net.Listen("tcp", cfg.Internal.Addr)
		if err != nil {
			logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "debug/HTTP", "addr", cfg.Internal.Addr)
			return http.Serve(debugListener, http.DefaultServeMux)
		}, func(error) {
			_ = debugListener.Close()
		})
	}
	{
		wait := make(chan struct{})
		subscriber := eventbuskafka.NewSubscriber(cfg.Kafka.Brokers,
			eventbuskafka.DefaultConfig(),
			eventbuskafka.WithSubscriberGroupName({{$.StreamPkgName}}.Name),
			eventbuskafka.WithSubscriberExitFunc(func() {
				logger.Log("component", "subscriber", "method", "exit")
			}),
			eventbuskafka.WithSubscriberErrorHandler(
				func(e *event.Event, err error) {
					logger.Log(
						"transport", "kafka",
						"method", "errorHandler",
						"event", e,
						"err", err)
				}))
		g.Add(func() error {
			logger.Log("transport", "kafka", "brokers", cfg.Kafka)
			subscriber.Subscribe({{$.StreamPkgName}}.Name, controller)
			if err := subscriber.Listen(ctx); err != nil {
				return err
			}
			<-wait
			return nil
		}, func(err error) {
			close(wait)
			_ = subscriber.Close()
		})
	}
	{
		httpListener, err := net.Listen("tcp", cfg.HTTP.Addr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", cfg.HTTP.Addr)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			_ = httpListener.Close()
		})
	}
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

func loadConfig() *config.Projection {
	filename := flag.String("config", os.Getenv("CONFIG_FILE"), "path to configuration file")
	flag.Parse()

	if len(*filename) == 0 {
		fmt.Println("configuration file not found")
		os.Exit(1)
	}

	cfg, err := config.ParseProjection(*filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return cfg
}