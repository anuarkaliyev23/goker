BIN_FOLDER ?= bin/

test:
	go clean -testcache
	go test ./... 

build:
	go clean -cache
	go build

build-multiplatform:
	./build.sh

commit-check: test build

hooks:
	chmod +x .hooks/pre-push
	chmod +x .hooks/prepare-commit-msg
	git config core.hooksPath .hooks/
