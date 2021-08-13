{{if $.StreamPublisher.IsKafka -}}
 kafka:
    brokers:
       - kafka: 9092
{{end -}}
http:
    addr: 127.0.0.1:8088
grpc:
    addr: 127.0.0.1:9098
internal:
    addr: 127.0.0.1:7072