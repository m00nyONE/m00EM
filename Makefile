OUT := m00EM
PKG := github.com/m00nyONE/m00EM
VERSION := $(shell git describe --always --long --dirty)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

all: run

server:
	go build -v -o ${OUT} -gcflags=all=-trimpath=$(GOPATH)/src -asmflags=all=-trimpath=${GOPATH}/src -ldflags="-X main.Version=${VERSION}" ${PKG}

test:
	@go test -short ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

static: vet lint
	go build -v -o ${OUT}-v${VERSION} -tags netgo -gcflags=all=-trimpath=$(GOPATH)/src -asmflags=all=-trimpath=$GOPATH/src -ldflags="-extldflags \"-static\" -w -s -X main.Version=${VERSION}" ${PKG}

run: server
	./${OUT} -version

clean:
	-@rm ${OUT} ${OUT}-v*

.PHONY: run server static vet lint
