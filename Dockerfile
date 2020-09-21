FROM golang:1.15-buster AS build

ENV CGO_ENABLED=0
ENV GO111MODULE="on"
ENV GOFLAGS="-mod=vendor"
WORKDIR /go/src/github.com/object88/namespaced-helm-install

ADD . .
RUN go build -o /usr/local/bin/namespaced-helm-install -tags "netgo" -ldflags "-extldflags \"-static\" -s -w" ./main.go

# FROM scratch
FROM debian:buster-slim
CMD ["/usr/local/bin/namespaced-helm-install"]

ADD testdata /opt/
COPY --from=build /usr/local/bin/namespaced-helm-install /usr/local/bin/