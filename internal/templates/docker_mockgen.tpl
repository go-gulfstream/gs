FROM golang:{{$.GoVersion}} AS builder
ENV MOCKGEN_VERSION "1.6.0"
RUN apt-get update -yqq && \
  apt-get install -yqq curl git unzip
RUN curl -sfLo mockgen.zip "https://github.com/golang/mock/archive/v${MOCKGEN_VERSION}.zip" && \
  unzip -q -d mockgen mockgen.zip && \
  cd ./mockgen/mock-${MOCKGEN_VERSION} && go mod download && GO111MODULE=on go build -o /go/bin/mockgen ./mockgen
FROM golang:{{$.GoVersion}}
COPY --from=builder /go/bin/mockgen /usr/local/bin/mockgen
ENTRYPOINT ["/usr/local/bin/mockgen"]