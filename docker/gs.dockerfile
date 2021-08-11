FROM golang:alpine AS builder
LABEL stage=builder
ARG BUILD_VERSION
RUN apk update && apk add --no-cache bash
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -ldflags="-w -s -X 'main.buildVersion=$BUILD_VERSION'" ./cmd/gs
FROM scratch
COPY --from=builder /app/gs /app/gs
ENTRYPOINT ["/app/gs"]