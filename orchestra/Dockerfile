FROM golang:latest AS builder

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go get -v github.com/wasperen/go-cluster-demo/cmd/cacher

CMD /go/bin/cacher

FROM scratch
COPY --from=builder /go/bin/cacher /
ENTRYPOINT ["/cacher"]