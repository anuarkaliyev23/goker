#!/bin/bash

platforms=("darwin/amd64" "darwin/arm64" "freebsd/amd64" "freebsd/arm" "linux/amd64" "linux/arm" "linux/arm64" "windows/amd64")
packageName="goker"

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}

	env GOOS=$GOOS GOARCH=$GOARCH go build -o "bin/${packageName}-${GOOS}-${GOARCH}"
done

