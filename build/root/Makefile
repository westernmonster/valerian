# Don't allow an implicit 'all' rule. This is not a user-facing file.
ifeq ($(MAKECMDGOALS),)
    $(error This Makefile requires an explicit rule to be specified)
endif

ifeq ($(DBG_MAKEFILE),1)
    $(warning ***** starting Makefile.generated_files for goal(s) "$(MAKECMDGOALS)")
    $(warning ***** $(shell date))
endif

GOBIN  := $(go env GOBIN)
ifeq ($(GOBIN),)
   GOBIN := ~/go/bin
endif

SHELL := /bin/bash
ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"

# Git stuff management
GIT_VERSION := $(shell git --no-pager describe --tags --always --dirty)
GIT_COMMIT  := $(shell git rev-parse --verify HEAD)
GIT_DATE    := $(firstword $(shell git --no-pager show --date=iso-strict --format="%ad" --name-only))
BUILD_DATE=$(shell date -u '+%Y%m%d')
DOCKER_TAG=${BUILD_DATE}-${GIT_VERSION}




.PHONY: build
# all: check bazel-update

build: init

clean:
	rm -rf _output
init:
	echo "hello"


service-msm: build-service-msm docker-build-service-msm

build-service-msm:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/service/msm/cmd/cmd"   app/service/msm/cmd/main.go
docker-build-service-msm:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/service-msm:${DOCKER_TAG} ./app/service/msm/cmd/
docker-push-service-msm:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/service-msm:${DOCKER_TAG}
local-deploy-service-msm:
	echo "hello"



app-locale: build-app-locale docker-build-app-locale docker-push-app-locale

build-app-locale:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/locale/cmd/cmd"   app/interface/locale/cmd/main.go
docker-build-app-locale:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/app-locale:${DOCKER_TAG} ./app/interface/locale/cmd/
docker-push-app-locale:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/app-locale:${DOCKER_TAG}

