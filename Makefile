BIN_FOLDER ?= bin/

test:
	go clean -testcache
	go test ./... 

build:
	go clean -cache
	go build

build-multiplatform:
	mkdir -p ${BIN_FOLDER}
	./build.sh

commit-check: test build

hooks:
	chmod +x .hooks/pre-push
	chmod +x .hooks/prepare-commit-msg
	git config core.hooksPath .hooks/


tag:
	git tag $(cat VERSION)
	git push origin --tags
