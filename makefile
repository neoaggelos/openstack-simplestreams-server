all: docker

IMAGE=neoaggelos/openstack-simplestreams-server
VERSION=0.2.0

docker-build: *.go Dockerfile
	docker build -t $(IMAGE):$(VERSION) .

docker: docker-build
	docker push $(IMAGE):$(VERSION)
