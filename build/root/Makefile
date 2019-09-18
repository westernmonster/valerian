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



app-locale: build-app-locale docker-build-app-locale docker-push-app-locale docker-deploy-app-locale

build-app-locale:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/locale/cmd/cmd"   app/interface/locale/cmd/main.go
docker-build-app-locale:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/app-locale:${DOCKER_TAG} ./app/interface/locale/cmd/
docker-push-app-locale:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/app-locale:${DOCKER_TAG}
docker-deploy-app-locale:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/app-locale:${DOCKER_TAG} --set global.name=locale --set env.region=hz --set env.zone=hz001 --set env.deployEnv=uat --set env.appId=app.locale --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin/conf --set env.confEnv=uat --set env.confToken=740ac786d93c11e988f22c4d549dd5ee  --set env.treeId=1 --set env.confVersion=1 --set env.confHost=config.pronote.loc --name locale ./build/k8s


admin-config: build-admin-config docker-build-admin-config docker-push-admin-config docker-deploy-admin-config

build-admin-config:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/admin/config/cmd/cmd"   app/admin/config/cmd/main.go
docker-build-admin-config:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/admin-config:${DOCKER_TAG} ./app/admin/config/cmd/
docker-push-admin-config:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/admin-config:${DOCKER_TAG}
docker-deploy-admin-config:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/admin-config:${DOCKER_TAG} --set global.name=admin-config --set env.region=hz --set env.zone=hz001 --set env.deployEnv=uat --set env.appId=admin.config --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=b2dc5b32da4a11e9908b00163e147fe0  --set env.treeId=1 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name admin-config-uat ./build/k8s

