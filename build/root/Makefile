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


deploy-interface-account:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-account:${DOCKER_TAG} --set global.name=interface-account --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.account --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=57b94857da6011e984660a58ac14011a  --set env.treeId=1401 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-account-uat ./build/k8s

deploy-interface-valcode:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-valcode:${DOCKER_TAG} --set global.name=interface-valcode --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.valcode --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=bf3f77c9da6011e984660a58ac14011a  --set env.treeId=1418 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-valcode-uat ./build/k8s

deploy-interface-passport-register:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-passport-register:${DOCKER_TAG} --set global.name=interface-passport-register --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.passport-register --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=9f5a4feeda6011e997420a58ac140026  --set env.treeId=1415 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-passport-register-uat ./build/k8s

deploy-interface-passport-login:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-passport-login:${DOCKER_TAG} --set global.name=interface-passport-login --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.passport-login --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=9b06ae26da6011e9884c0a58ac14008b  --set env.treeId=1414 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-passport-login-uat ./build/k8s

deploy-interface-passport-auth:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-passport-auth:${DOCKER_TAG} --set global.name=interface-passport-auth --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.passport-auth --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=98a4e79bda6011e984660a58ac14011a  --set env.treeId=1413 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-passport-auth-uat ./build/k8s

deploy-interface-feedback:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-feedback:${DOCKER_TAG} --set global.name=interface-feedback --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.feedback --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=8459c49bda6011e997420a58ac140026  --set env.treeId=1409 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-feedback-uat ./build/k8s

deploy-interface-certification:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-certification:${DOCKER_TAG} --set global.name=interface-certification --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.certification --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=6512dc65da6011e997420a58ac140026  --set env.treeId=1403 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-certification-uat ./build/k8s

deploy-interface-topic:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-topic:${DOCKER_TAG} --set global.name=interface-topic --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.topic --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=ba6ea185da6011e997420a58ac140026  --set env.treeId=1417 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-topic-uat ./build/k8s

deploy-interface-discuss:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-discuss:${DOCKER_TAG} --set global.name=interface-discuss --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.discuss --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=69e78998da6011e984660a58ac14011a  --set env.treeId=1404 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-discuss-uat ./build/k8s

deploy-interface-article:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-article:${DOCKER_TAG} --set global.name=interface-article --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.article --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=610cae78da6011e9884c0a58ac14008b  --set env.treeId=1402 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-article-uat ./build/k8s

deploy-interface-like:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-like:${DOCKER_TAG} --set global.name=interface-like --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.like --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=503a91b9e3fa11e986590a58ac140094  --set env.treeId=1419 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-like-uat ./build/k8s

deploy-interface-fav:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-fav:${DOCKER_TAG} --set global.name=interface-fav --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.fav --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=76faaa4ada6011e984660a58ac14011a  --set env.treeId=1407 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-fav-uat ./build/k8s

deploy-interface-feed:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-feed:${DOCKER_TAG} --set global.name=interface-feed --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.feed --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=7f5850a8da6011e9884c0a58ac14008b  --set env.treeId=1408 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-feed-uat ./build/k8s

deploy-interface-comment:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-comment:${DOCKER_TAG} --set global.name=interface-comment --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.comment --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=0adcc057edfd11e986590a58ac140094  --set env.treeId=1421 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-comment-uat ./build/k8s

deploy-interface-search:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-search:${DOCKER_TAG} --set global.name=interface-search --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.search --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=2969e3a8edfd11e9b6a00a58ac14002b  --set env.treeId=1420 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-search-uat ./build/k8s

deploy-interface-recent:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-recent:${DOCKER_TAG} --set global.name=interface-recent --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.recent --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=b49ad0f4da6011e9884c0a58ac14008b  --set env.treeId=1416 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-recent-uat ./build/k8s

deploy-interface-dm:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/interface-dm:${DOCKER_TAG} --set global.name=interface-dm --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=interface.dm --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=6ee6e604da6011e9884c0a58ac14008b  --set env.treeId=1405 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc interface-dm-uat ./build/k8s


deploy-admin-config:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/admin-config:${DOCKER_TAG} --set global.name=admin-config --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=admin.config --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=25216dcada6011e997420a58ac140026  --set env.treeId=1101 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc admin-config-uat ./build/k8s



deploy-service-msm:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-msm:${DOCKER_TAG} --set global.name=service-msm --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.msm --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=3f7c84ebda6011e984660a58ac14011a  --set env.treeId=1303 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-msm-uat ./build/k8s

deploy-service-identify:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-identify:${DOCKER_TAG} --set global.name=service-identify --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.identify --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=3af51098da6011e997420a58ac140026  --set env.treeId=1302 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-identify-uat ./build/k8s



deploy-service-account:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-account:${DOCKER_TAG} --set global.name=service-account --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.account --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=379d0e0ada6011e9884c0a58ac14008b  --set env.treeId=1301 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-account-uat ./build/k8s


deploy-service-discuss:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-discuss:${DOCKER_TAG} --set global.name=service-discuss --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.discuss --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=d388501cdd7511e986590a58ac140094  --set env.treeId=1306 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-discuss-uat ./build/k8s

deploy-service-topic:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-topic:${DOCKER_TAG} --set global.name=service-topic --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.topic --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=49c6fcc6da6011e997420a58ac140026  --set env.treeId=1305 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-topic-uat ./build/k8s

deploy-service-feed:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-feed:${DOCKER_TAG} --set global.name=service-feed --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.feed --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=42385238dfe711e986590a58ac140094  --set env.treeId=1307 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-feed-uat ./build/k8s

deploy-service-fav:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-fav:${DOCKER_TAG} --set global.name=service-fav --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.fav --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=b86b18ade3f911e9b6a00a58ac14002b  --set env.treeId=1308 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-fav-uat ./build/k8s


deploy-service-like:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-like:${DOCKER_TAG} --set global.name=service-like --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.like --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=abdbe3cfe3f911e986590a58ac140094  --set env.treeId=1309 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-like-uat ./build/k8s

deploy-service-relation:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-relation:${DOCKER_TAG} --set global.name=service-relation --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.relation --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=45745073da6011e9884c0a58ac14008b  --set env.treeId=1304 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-relation-uat ./build/k8s


deploy-service-article:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-article:${DOCKER_TAG} --set global.name=service-article --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.article --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=68fefc9de93811e986590a58ac140094  --set env.treeId=1310 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-article-uat ./build/k8s

deploy-service-recent:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-recent:${DOCKER_TAG} --set global.name=service-recent --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.recent --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=0aa2605ae94511e9b6a00a58ac14002b  --set env.treeId=1311 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-recent-uat ./build/k8s


deploy-service-topic-feed:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-topic-feed:${DOCKER_TAG} --set global.name=service-topic-feed --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.topic-feed --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=08b7ff7eea0811e986590a58ac140094  --set env.treeId=1312 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-topic-feed-uat ./build/k8s

deploy-service-account-feed:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-account-feed:${DOCKER_TAG} --set global.name=service-account-feed --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.account-feed --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=0fdd0111ea0811e9b6a00a58ac14002b  --set env.treeId=1313 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-account-feed-uat ./build/k8s

deploy-service-message:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-message:${DOCKER_TAG} --set global.name=service-message --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.message --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=16e51e5cea0811e98a140a58ac14011e  --set env.treeId=1314 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-message-uat ./build/k8s

deploy-service-comment:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-comment:${DOCKER_TAG} --set global.name=service-comment --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.comment --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=b596198aedfc11e986590a58ac140094  --set env.treeId=1315 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-comment-uat ./build/k8s

deploy-service-search:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-search:${DOCKER_TAG} --set global.name=service-search --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.search --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=17e3ee7fefff11e986590a58ac140094  --set env.treeId=1316 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-search-uat ./build/k8s

deploy-service-certification:
	helm upgrade --install --namespace uat --set global.image=registry-vpc.cn-hangzhou.aliyuncs.com/valerian/service-certification:${DOCKER_TAG} --set global.name=service-certification --set env.region=hz --set env.zone=hz001 --set env.siteUrl=https://uat.flywk.com  --set env.deployEnv=uat --set env.appId=service.certification --set env.trace=jaeger-uat-agent:6831 --set env.confPath=/go/bin --set env.confEnv=uat --set env.confToken=c23a51c9f69e11e9a3610a58ac140155  --set env.treeId=1317 --set env.confVersion=0.1 --set env.confHost=config.pronote.loc service-certification-uat ./build/k8s
