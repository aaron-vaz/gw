REPO:=/go/src/github.com/aaron-vaz/gw

BUILD_DIR:=${PWD}/build
BINARY_DIR:=${BUILD_DIR}/binaries

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS:=-ldflags "-s -w"

clean:
	@-rm -rf ${BUILD_DIR}

setup_workspace:
	@mkdir -p ${BUILD_DIR} ${BINARY_DIR}

get_dependencies:
	@go get github.com/mitchellh/gox

test:
	@go test -cover ./cmd/gw

build: setup_workspace get_dependencies test
	@gox -verbose ${LDFLAGS} -os="windows linux darwin" -arch="amd64" -output="${BINARY_DIR}/gw_{{.OS}}_{{.Arch}}" ./cmd/gw

docker:
	@docker run --rm -v ${PWD}:${REPO} -w=${REPO} -e "GO111MODULE=on" golang make build

.PHONY: docker build test clean setup_workspace get_dependencies