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
# We don't need make's built-in rules.
MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

# Constants used throughout.
.EXPORT_ALL_VARIABLES:
OUT_DIR ?= _output
BIN_DIR := $(OUT_DIR)/bin

.PHONY: build update clean
all: check bazel-update

build: init check bazel-build

ci-build: check bazel-build

simple-build:
	bazel build --watchfs -- //tools/... -//vendor/...
ifeq ($(WHAT),)
bazel-build:
	bazel build --config=office --watchfs //app/...  //library/...
else
bazel-build:
	bazel build --config=ci -- //$(WHAT)/...
endif

build-keep-going:
	bazel build --config=ci -k //app/... //build/... //library/...
	cat bazel-out/stable-status.txt
clean:
	bazel clean --expunge
	rm -rf _output
update: init bazel-update

bazel-update:
	./build/update-bazel.sh
prow-update:
	./build/update-prow.sh
test:
	@if [ "$(WHAT)" !=  "" ]; \
         then \
	 cd $(WHAT) && make ; \
	 else \
	 echo "Please input the WHAT" ;\
	 fi

bazel-test:
	@if [ "$(WHAT)" !=  "" ]; \
         then \
	 bazel test --watchfs -- //$(WHAT)/... ; \
	 else \
	 echo "Please input the WHAT" ;\
	 fi
check:
	@./build/check.sh

init:
	@if [ ! -f .git/hooks/pre-commit ] ; \
	then \
	echo "make all" >> .git/hooks/pre-commit; \
	sudo chmod +x .git/hooks/pre-commit; \
	fi
ci-bazel-build:
	bazel build --config=ci -- //app/...
ci-bazel-build-b:
	bazel build --config=ci -- //app/interface/...


ci-bazel-build-d:
	bazel build --config=ci -- //app/service/... //library/...
