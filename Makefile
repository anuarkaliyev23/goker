test:
	go clean -testcache
	go test ./... 

build:
	go clean -cache
	go build

commit-check: test build

hooks:
	chmod +x .hooks/pre-commit
	git config core.hooksPath .hooks/
