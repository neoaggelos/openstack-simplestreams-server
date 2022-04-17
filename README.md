# OpenStack Simplestreams Server

Run a server that exposes a simplestreams endpoint for the Ubuntu images of an OpenStack Glance instance.

## Install

```bash
openstack user create productstreams --domain default --project service --password-prompt
```

Create `openrc` file with `OS_*` credentials for the productstreams user. Run server locally:

```bash



## Development

```bash
go run . -listen 127.0.0.1:8080
```
