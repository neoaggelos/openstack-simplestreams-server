FROM golang:1.18 AS builder
WORKDIR /src
COPY . /src
RUN CGO_ENABLED=0 go build -a -ldflags '-s -w -extldflags "-static"' -o /openstack-simplestreams-server /src

FROM alpine
COPY --from=builder /openstack-simplestreams-server /openstack-simplestreams-server
ENTRYPOINT ["/openstack-simplestreams-server"]
