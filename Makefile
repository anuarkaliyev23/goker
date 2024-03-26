test:
	go clean -testcache
	go test ./... 

build:
	go clean -cache
	go build
