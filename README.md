# OpenStack Simplestreams Server

Run a server that exposes a simplestreams endpoint for the Ubuntu images of an OpenStack Glance instance.

## Install

```bash
openstack user create product-streams --domain default --project service --password-prompt
openstack role add --user product-streams --project service reader
```

Create `openrc` file with `OS_*` credentials for the productstreams user. Run server locally:

```bash
# with docker-compose
docker-compose up -d

# or with docker
docker create --name openstack_simplestreams_server --restart always --network host --env-file ./openrc --restart=always neoaggelos/openstack-simplestreams-server:0.3.0
docker start openstack_simplestreams_server
```

## Release

```bash
make docker IMAGE=neoaggelos/openstack-simplestreams-version VERSION=0.3.0
```

## Development

```bash
go run . -listen 127.0.0.1:8080
```
