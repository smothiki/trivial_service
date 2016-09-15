SHORT_NAME ?= trivial_service
# dockerized development environment variables
REPO_PATH := github.com/smothiki/${SHORT_NAME}
DEV_ENV_IMAGE := golang:1.6-alpine#blang/golang-alpine
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_PREFIX := docker run --rm -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR} #-e GOVENDOREXPERIMENT=1
DEV_ENV_CMD := ${DEV_ENV_PREFIX} ${DEV_ENV_IMAGE}
DEV_GLIDE_CMD := ${DEV_ENV_PREFIX} quay.io/deis/go-dev:0.17.0


# SemVer with build information is defined in the SemVer 2 spec, but Docker
# doesn't allow +, so we use -.
BINARY_DEST_DIR := rootfs/usr/bin
# Common flags passed into Go's linker.
LDFLAGS := "-s -w -X main.version=${VERSION}"
# Docker Root FS
BINDIR := ./rootfs

GOTEST := go test --race

bootstrap:
	${DEV_GLIDE_CMD} glide install

glideup:
	${DEV_GLIDE_CMD} glide up

build-proxy:
	${DEV_ENV_CMD} go build -ldflags ${LDFLAGS} -o ${BINARY_DEST_DIR}/proxy proxy/proxy.go

build-www:
	${DEV_ENV_CMD} go build -ldflags ${LDFLAGS} -o ${BINARY_DEST_DIR}/www www/www.go

build-backend:
	${DEV_ENV_CMD} go build -ldflags ${LDFLAGS} -o ${BINARY_DEST_DIR}/backend backend/backend.go

test: test-style test-unit

test-style:
	${DEV_ENV_CMD} lint

test-unit:
	${DEV_ENV_CMD} sh -c '${GOTEST} $$(glide nv)'

test-cover:
	${DEV_ENV_CMD} test-cover.sh

docker-build: build-proxy build-www build-backend
	docker build --rm -t quantum rootfs

deploy:
	kubectl create -f manifests/backend_service.yaml
	kubectl create -f manifests/frontend_service.yaml
	kubectl create -f manifests/backend_rs.yaml
	kubectl create -f manifests/frontend_rs.yaml
	kubectl create -f manifests/proxy_rs.yaml
