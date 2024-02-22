test:
	go clean -testcache
	go test ./... 

build:
	go build
