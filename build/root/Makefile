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



interface-location: build-interface-location docker-interface-location push-interface-location deploy-interface-location

build-interface-location:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/location/cmd/cmd"   app/interface/location/cmd/main.go
docker-interface-location:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/interface-location:${DOCKER_TAG} ./app/interface/location/cmd/
push-interface-location:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/interface-location:${DOCKER_TAG}
deploy-interface-location:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/interface-location:${DOCKER_TAG} --set global.name=interface-location --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.location --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=914e02fdda6011e997420a58ac140026  --set env.treeId=1412 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name interface-location-uat ./build/k8s


interface-locale: build-interface-locale docker-interface-locale push-interface-locale deploy-interface-locale

build-interface-locale:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/locale/cmd/cmd"   app/interface/locale/cmd/main.go
docker-interface-locale:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/interface-locale:${DOCKER_TAG} ./app/interface/locale/cmd/
push-interface-locale:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/interface-locale:${DOCKER_TAG}
deploy-interface-locale:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/interface-locale:${DOCKER_TAG} --set global.name=interface-locale --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.locale --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=8c534b50da6011e9884c0a58ac14008b  --set env.treeId=1411 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name interface-locale-uat ./build/k8s

interface-valcode: build-interface-valcode docker-interface-valcode push-interface-valcode deploy-interface-valcode

build-interface-valcode:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/valcode/cmd/cmd"   app/interface/valcode/cmd/main.go
docker-interface-valcode:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/interface-valcode:${DOCKER_TAG} ./app/interface/valcode/cmd/
push-interface-valcode:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/interface-valcode:${DOCKER_TAG}
deploy-interface-valcode:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/interface-valcode:${DOCKER_TAG} --set global.name=interface-valcode --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.valcode --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=bf3f77c9da6011e984660a58ac14011a  --set env.treeId=1418 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name interface-valcode-uat ./build/k8s

interface-passport-register: build-interface-passport-register docker-interface-passport-register push-interface-passport-register deploy-interface-passport-register

build-interface-passport-register:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/passport-register/cmd/cmd"   app/interface/passport-register/cmd/main.go
docker-interface-passport-register:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/interface-passport-register:${DOCKER_TAG} ./app/interface/passport-register/cmd/
push-interface-passport-register:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/interface-passport-register:${DOCKER_TAG}
deploy-interface-passport-register:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/interface-passport-register:${DOCKER_TAG} --set global.name=interface-passport-register --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.passport-register --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=9f5a4feeda6011e997420a58ac140026  --set env.treeId=1415 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name interface-passport-register-uat ./build/k8s

interface-passport-login: build-interface-passport-login docker-interface-passport-login push-interface-passport-login deploy-interface-passport-login

build-interface-passport-login:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/passport-login/cmd/cmd"   app/interface/passport-login/cmd/main.go
docker-interface-passport-login:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/interface-passport-login:${DOCKER_TAG} ./app/interface/passport-login/cmd/
push-interface-passport-login:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/interface-passport-login:${DOCKER_TAG}
deploy-interface-passport-login:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/interface-passport-login:${DOCKER_TAG} --set global.name=interface-passport-login --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.passport-login --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=9b06ae26da6011e9884c0a58ac14008b  --set env.treeId=1414 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name interface-passport-login-uat ./build/k8s

interface-passport-auth: build-interface-passport-auth docker-interface-passport-auth push-interface-passport-auth deploy-interface-passport-auth

build-interface-passport-auth:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/passport-auth/cmd/cmd"   app/interface/passport-auth/cmd/main.go
docker-interface-passport-auth:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/interface-passport-auth:${DOCKER_TAG} ./app/interface/passport-auth/cmd/
push-interface-passport-auth:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/interface-passport-auth:${DOCKER_TAG}
deploy-interface-passport-auth:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/interface-passport-auth:${DOCKER_TAG} --set global.name=interface-passport-auth --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.passport-auth --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=98a4e79bda6011e984660a58ac14011a  --set env.treeId=1413 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name interface-passport-auth-uat ./build/k8s

interface-file: build-interface-file docker-interface-file push-interface-file deploy-interface-file

build-interface-file:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/file/cmd/cmd"   app/interface/file/cmd/main.go
docker-interface-file:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/interface-file:${DOCKER_TAG} ./app/interface/file/cmd/
push-interface-file:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/interface-file:${DOCKER_TAG}
deploy-interface-file:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/interface-file:${DOCKER_TAG} --set global.name=interface-file --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.file --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=880b1154da6011e984660a58ac14011a  --set env.treeId=1410 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name interface-file-uat ./build/k8s

interface-feedback: build-interface-feedback docker-interface-feedback push-interface-feedback deploy-interface-feedback

build-interface-feedback:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/feedback/cmd/cmd"   app/interface/feedback/cmd/main.go
docker-interface-feedback:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/interface-feedback:${DOCKER_TAG} ./app/interface/feedback/cmd/
push-interface-feedback:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/interface-feedback:${DOCKER_TAG}
deploy-interface-feedback:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/interface-feedback:${DOCKER_TAG} --set global.name=interface-feedback --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.feedback --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=8459c49bda6011e997420a58ac140026  --set env.treeId=1409 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name interface-feedback-uat ./build/k8s

interface-editor: build-interface-editor docker-interface-editor push-interface-editor deploy-interface-editor

build-interface-editor:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/editor/cmd/cmd"   app/interface/editor/cmd/main.go
docker-interface-editor:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/interface-editor:${DOCKER_TAG} ./app/interface/editor/cmd/
push-interface-editor:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/interface-editor:${DOCKER_TAG}
deploy-interface-editor:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/interface-editor:${DOCKER_TAG} --set global.name=interface-editor --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.editor --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=72166bebda6011e997420a58ac140026  --set env.treeId=1406 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name interface-editor-uat ./build/k8s

interface-certification: build-interface-certification docker-interface-certification push-interface-certification deploy-interface-certification

build-interface-certification:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/certification/cmd/cmd"   app/interface/certification/cmd/main.go
docker-interface-certification:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/interface-certification:${DOCKER_TAG} ./app/interface/certification/cmd/
push-interface-certification:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/interface-certification:${DOCKER_TAG}
deploy-interface-certification:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/interface-certification:${DOCKER_TAG} --set global.name=interface-certification --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.certification --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=6512dc65da6011e997420a58ac140026  --set env.treeId=1403 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name interface-certification-uat ./build/k8s

interface-account: build-interface-account docker-interface-account push-interface-account deploy-interface-account

build-interface-account:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/interface/account/cmd/cmd"   app/interface/account/cmd/main.go
docker-interface-account:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/interface-account:${DOCKER_TAG} ./app/interface/account/cmd/
push-interface-account:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/interface-account:${DOCKER_TAG}
deploy-interface-account:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/interface-account:${DOCKER_TAG} --set global.name=interface-account --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.account --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=57b94857da6011e984660a58ac14011a  --set env.treeId=1401 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name interface-account-uat ./build/k8s


admin-config: build-admin-config docker-admin-config push-admin-config deploy-admin-config

build-admin-config:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/admin/config/cmd/cmd"   app/admin/config/cmd/main.go
docker-admin-config:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/admin-config:${DOCKER_TAG} ./app/admin/config/cmd/
push-admin-config:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/admin-config:${DOCKER_TAG}
deploy-admin-config:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/admin-config:${DOCKER_TAG} --set global.name=admin-config --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=admin.config --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=25216dcada6011e997420a58ac140026  --set env.treeId=1101 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name admin-config-uat ./build/k8s



service-msm: build-service-msm docker-service-msm push-service-msm deploy-service-msm

build-service-msm:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/service/msm/cmd/cmd"   app/service/msm/cmd/main.go
docker-service-msm:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/service-msm:${DOCKER_TAG} ./app/service/msm/cmd/
push-service-msm:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/service-msm:${DOCKER_TAG}
deploy-service-msm:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/service-msm:${DOCKER_TAG} --set global.name=service-msm --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.msm --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=3f7c84ebda6011e984660a58ac14011a  --set env.treeId=1303 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name service-msm-uat ./build/k8s

service-identify: build-service-identify docker-service-identify push-service-identify deploy-service-identify

build-service-identify:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/service/identify/cmd/cmd"   app/service/identify/cmd/main.go
docker-service-identify:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/service-identify:${DOCKER_TAG} ./app/service/identify/cmd/
push-service-identify:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/service-identify:${DOCKER_TAG}
deploy-service-identify:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/service-identify:${DOCKER_TAG} --set global.name=service-identify --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.identify --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=3af51098da6011e997420a58ac140026  --set env.treeId=1302 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name service-identify-uat ./build/k8s

service-discuss: build-service-discuss docker-service-discuss push-service-discuss deploy-service-discuss

build-service-discuss:
	go env -w GOPROXY=https://goproxy.cn,direct
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-w -s" -a -installsuffix cgo -o "app/service/discuss/cmd/cmd"   app/service/discuss/cmd/main.go
docker-service-discuss:
	docker build --no-cache --build-arg EXE=./cmd -f build/Dockerfile  -t registry.cn-hangzhou.aliyuncs.com/valerian/service-discuss:${DOCKER_TAG} ./app/service/discuss/cmd/
push-service-discuss:
	docker push registry.cn-hangzhou.aliyuncs.com/valerian/service-discuss:${DOCKER_TAG}
deploy-service-discuss:
	helm install --namespace uat --set global.image=registry.cn-hangzhou.aliyuncs.com/valerian/service-discuss:${DOCKER_TAG} --set global.name=service-discuss --set env.region=hz --set env.zone=hz001 env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.discuss --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=d388501cdd7511e986590a58ac140094  --set env.treeId=1306 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc --name service-discuss-uat ./build/k8s
