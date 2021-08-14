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

	_ "net/http/pprof"

	"{{$.GoModules}}/pkg/{{$.StreamPkgName}}"

     metricsprometheus "github.com/go-gulfstream/gulfstream/pkg/metrics/prometheus"

	"github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/collectors"
    "github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-gulfstream/gulfstream/pkg/event"
	eventbuskafka "github.com/go-gulfstream/gulfstream/pkg/eventbus/kafka"
	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"

	"github.com/oklog/oklog/pkg/group"

	"{{$.GoModules}}/internal/config"

	"{{$.GoModules}}/internal/api"

	"{{$.GoModules}}/internal/projection"
	"{{$.GoModules}}/pkg/{{$.PackageName}}query"

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

	promReg := prometheus.NewRegistry()
    promReg.MustRegister(collectors.NewGoCollector())

	storage := projection.NewStorage()
	handler := projection.New(storage)
	controller := projection.NewController(handler)
    controllerWithInterceptor := gulfstream.WithEventHandlerInterceptor(controller,
		metricsprometheus.NewEventHandlerMetrics(promReg))

	var service {{$.PackageName}}query.Service
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
		internalListener, err := net.Listen("tcp", cfg.Internal.Addr)
		if err != nil {
			_ = logger.Log("transport", "internal/HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			_ = logger.Log("transport", "debug/HTTP", "addr", cfg.Internal.Addr)
			http.Handle("/metrics", promhttp.HandlerFor(promReg, promhttp.HandlerOpts{}))
			return http.Serve(internalListener, http.DefaultServeMux)
		}, func(error) {
			_ = internalListener.Close()
		})
	}
	{
		wait := make(chan struct{})
		subscriber := eventbuskafka.NewSubscriber(cfg.Kafka.Brokers,
			eventbuskafka.DefaultConfig(),
			eventbuskafka.WithSubscriberGroupName({{$.StreamPkgName}}.Name),
			eventbuskafka.WithSubscriberExitFunc(func() {
				_ = logger.Log("component", "subscriber", "method", "exit")
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
			_ = logger.Log("transport", "kafka", "brokers", cfg.Kafka)
			subscriber.Subscribe({{$.StreamPkgName}}.Name, controllerWithInterceptor)
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
			_ = logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			_ = logger.Log("transport", "HTTP", "addr", cfg.HTTP.Addr)
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

	_ = logger.Log("exit", g.Run())
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