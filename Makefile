APPS=twemproxy-exporter
APPS_VERSION=$(shell cat VERSION | tr -d '[:blank:]')
DOCKER_APPS=$(shell which docker)

build:
	$(DOCKER_APPS) build -t $(APPS):$(APPS_VERSION) . 

run:
	$(DOCKER_APPS) run $(APPS):$(APPS_VERSION)
