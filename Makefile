REPO:=/go/src/github.com/aaron-vaz/gw

BUILD_DIR:=${PWD}/build
BINARY_DIR:=${BUILD_DIR}/binaries

NAME:=gw

GCFLAGS:=-gcflags "all=-trimpath=${PWD}"
ASMFLAGS:=-asmflags "all=-trimpath=${PWD}"
LDFLAGS:=-ldflags "-s -w -extldflags=-zrelro -extldflags=-znow"


clean:
	@-rm -rf ${BUILD_DIR}

setup_workspace:
	@mkdir -p ${BUILD_DIR} ${BINARY_DIR}

get_dependencies:
	@go get github.com/mitchellh/gox

test:
	@go test -cover ./cmd/${NAME}

build:
	@go build ${GCFLAGS} ${ASMFLAGS} ${LDFLAGS} -o ${BUILD_DIR}/${NAME} ./cmd/${NAME}

release: setup_workspace get_dependencies test
	@gox -verbose ${GCFLAGS} ${LDFLAGS} -os="windows linux darwin" -arch="amd64" -output="${BINARY_DIR}/${NAME}_{{.OS}}_{{.Arch}}" ./cmd/${NAME}

docker:
	@docker run --rm -v ${PWD}:${REPO} -w=${REPO} -e "GO111MODULE=on" golang make release

.PHONY: docker build release test clean setup_workspace get_dependencies
