internal_http:
    addr: 127.0.0.1:7070
{{if $.StreamStorage.IsPostgres -}}
postgres:
    addr: postgres://user:password@127.0.0.1:5432/postgres
{{else if $.StreamStorage.IsRedis -}}
redis:
    addr: 127.0.0.1:6379
{{end -}}
{{if $.CommandBus.IsGRPC -}}
commandbus_grpc:
    addr: 127.0.0.1:9091
{{else if $.CommandBus.IsNATS -}}
commandbus_nats:
    addr: nats://127.0.0.1
{{else if $.CommandBus.IsHTTP -}}
commandbus_http:
    addr: 127.0.0.1:7092
{{end -}}
 kafka:
    brokers:
       - kafka:9092

