BASE_DIR := $(shell pwd)
BUILD_DIR := $(BASE_DIR)/build
VERSION ?= $(shell git describe --exact-match 2> /dev/null || \
                 git describe --match=$(git rev-parse --short=8 HEAD) \
		 --always --dirty --abbrev=8)

version:
	@echo $(VERSION)

all: build

build: osbchecker mockbroker

prebuild:
	mkdir -p $(BUILD_DIR)

.PHONY: osbchecker mockbroker

osbchecker: prebuild
	go test -c -o $(BUILD_DIR)/osbchecker.test github.com/openservicebrokerapi/osb-checker/test

mockbroker: prebuild
	go build -o $(BUILD_DIR)/osbchecker.mockbroker github.com/openservicebrokerapi/osb-checker/autogenerated/cmd/open-service-broker-server

docker: build
	cp $(BUILD_DIR)/osbchecker.test ./test/osbchecker.test
	cp $(BUILD_DIR)/osbchecker.mockbroker ./test/osbchecker.mockbroker
	docker build test/ -t openservicebrokerapi/osbchecker-test:$(VERSION)
	docker build autogenerated/cmd/open-service-broker-server/ -t openservicebrokerapi/osbchecker-mockbroker:$(VERSION)

clean:
	rm -rf $(BUILD_DIR) ./test/osbchecker.test ./test/osbchecker.mockbroker
