# OpenStack Simplestreams Server

Run a server that exposes a simplestreams endpoint for the Ubuntu images of an OpenStack Glance instance.

## Install

```bash
openstack user create productstreams --domain default --project service --password-prompt
openstack role add --user product-streams --project service reader
```

Create `openrc` file with `OS_*` credentials for the productstreams user. Run server locally:

```bash
docker-compose up -d
```

## Release

```bash
make docker IMAGE=neoaggelos/openstack-simplestreams-version VERSION=0.1.0
```

## Development

```bash
go run . -listen 127.0.0.1:8080
```
