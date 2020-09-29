REPO:=/go/src/github.com/aaron-vaz/gw

BUILD_DIR:=${PWD}/build
BINARY_DIR:=${BUILD_DIR}/binaries

NAME:=gw

GCFLAGS:=-gcflags "all=-trimpath=${PWD}"
ASMFLAGS:=-asmflags "all=-trimpath=${PWD}"
LDFLAGS:=-ldflags "-s -w -extldflags=-zrelro -extldflags=-znow"

RELEASE_OS=windows linux darwin
RELEASE_ARCH=amd64

clean:
	@-rm -rf ${BUILD_DIR}

setup_workspace:
	@mkdir -p ${BUILD_DIR} ${BINARY_DIR}

test:
	@go test -cover ./cmd/${NAME}

build:
	@go build ${GCFLAGS} ${ASMFLAGS} ${LDFLAGS} -o ${BUILD_DIR}/${NAME} ./cmd/${NAME}

release: $(RELEASE_OS) setup_workspace test

$(RELEASE_OS):
	@GOOS=$@ GOARCH=$(RELEASE_ARCH) go build ${GCFLAGS} ${ASMFLAGS} ${LDFLAGS} -o ${BINARY_DIR}/${NAME}_$@_${RELEASE_ARCH}$(shell [ $@ = "windows" ] && echo ".exe" || echo "" ) ./cmd/${NAME}

docker:
	@docker run --rm -v ${PWD}:${REPO} -w=${REPO} -e "GO111MODULE=on" golang make release

.PHONY: docker build release test clean setup_workspace
