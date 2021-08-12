FROM golang:alpine AS builder
ARG BUILD_VERSION
LABEL stage=builder
RUN apk update && apk add --no-cache bash
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -ldflags="-w -s -X 'main.buildVersion=$BUILD_VERSION'" ./cmd/{{$.PackageName}}
FROM scratch
COPY --from=builder /app/{{$.PackageName}} /app/{{$.PackageName}}
ENTRYPOINT ["/app/{{$.PackageName}}"]