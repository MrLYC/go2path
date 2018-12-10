VERSION = 0.0.1

ROOTDIR = $(shell pwd)
APPNAME = go2path
APPPATH = github.com/mrlyc/go2path
GODIR = ${GOPATH}
SRCDIR = ${GODIR}/src/${APPPATH}
TARGET = bin/${APPNAME}

GOENV = GOPATH=${GODIR} GO15VENDOREXPERIMENT=1

GO = ${GOENV} go

LDFLAGS = 

.PHONY: build
build: ${SRCDIR}
	${GO} build -i -ldflags="${LDFLAGS}" -o ${TARGET} ${APPPATH}

${SRCDIR}:
	mkdir -p bin
	mkdir -p `dirname "${SRCDIR}"`
	ln -s ${ROOTDIR} ${SRCDIR}

.PHONY: init
init: ${SRCDIR}

.PHONY: go-env
go-env:
	@go env
