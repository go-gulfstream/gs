# Builder
FROM golang:alpine AS builder
LABEL stage=builder
RUN apk update && apk add --no-cache bash
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -ldflags="-w -s" ./cmd/{{$.Project.Name}}

# Service
FROM scratch
COPY --from=builder /app/{{$.Project.Name}} /app/{{$.Project.Name}}
ENTRYPOINT ["/app/{{$.Project.Name}}"]