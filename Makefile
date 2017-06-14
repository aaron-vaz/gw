VERSION?=$(shell git describe --tags)
COMMIT:=$(shell git rev-parse HEAD)

BUILD_DIR:=$(shell pwd)/build
BINARY_DIR:=${BUILD_DIR}/binaries

# set environment variables
export GOPATH:=${BUILD_DIR}/gopath
export PATH:=$(PATH):${GOPATH}/bin
export OLDPWD:=$(shell pwd)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS:=-ldflags "-s -w -X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT}"

# Build the project
all: clean setup_workspace get_dependencies test build

clean:
	-rm -rf ${BUILD_DIR}

setup_workspace:
	mkdir -p ${BUILD_DIR} ${TEST_REPORT_DIR} ${BINARY_DIR}

get_dependencies:
	go get github.com/mitchellh/gox

build:
	cd ${BINARY_DIR}; \
	gox -verbose ${LDFLAGS} ../../ ; \
	cd - >/dev/null

test:
	go test -v --cover ; \
	cd - >/dev/null

.PHONY: build test clean setup_workspace get_dependencies