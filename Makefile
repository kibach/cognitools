# Hooks
CLEAN_TARGETS = clean-pkged
BUILD_RELEASE_DEP_TARGETS = bin/pkger
PRE_BUILD_RELEASE_TARGETS = webui clean-pkged $(patsubst cmd/%,cmd/%/pkged.go,$(wildcard cmd/*))
BUILD_DEBUG_DEP_TARGETS = bin/pkger
PRE_BUILD_DEBUG_TARGETS = $(patsubst cmd/%,cmd/%/pkged.go,$(wildcard cmd/*))

include main.mk

bin/pkger: go.mod
	@mkdir -p bin
	go build -o bin/pkger github.com/markbates/pkger/cmd/pkger

cmd/%/pkged.go: bin/pkger
	bin/pkger -o cmd/$*

.PHONY: clean-pkged
clean-pkged:
	rm -rf cmd/*/pkged.go

.PHONY: generate
generate:
	go generate -x ./...

webui:
	yarn --cwd ./web/ui/ install
	yarn --cwd ./web/ui/ build
