all: docker

IMAGE=neoaggelos/openstack-simplestreams-server
VERSION=0.4.0

docker-build: *.go Dockerfile
	docker build -t $(IMAGE):$(VERSION) . --network=host

docker: docker-build
	docker push $(IMAGE):$(VERSION)
