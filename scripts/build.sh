#!/usr/bin/env bash

set -e

VERSION=$(git describe --tag  || echo -n "UNDEFINED")
BUILDTIME=$(date +%Y-%m-%d\ %H:%M:%S || echo -n "UNDEFINED")
COMMITHASH=$(git rev-parse --verify HEAD || echo -n "UNDEFINED")

GO_FLAGS=(-ldflags "-w -s -X 'github.com/Bilou4/godo/cmd.Version=$VERSION' -X 'github.com/Bilou4/godo/cmd.BuildTime=$BUILDTIME' -X 'github.com/Bilou4/godo/cmd.CommitHash=$COMMITHASH'" -trimpath)
ENABLE_CGO=0
PACKAGE_NAME=godo
OUTPUT_DIR=build

PLATFORMS=("windows/amd64" "linux/amd64")


for platform in "${PLATFORMS[@]}"
do
    # split platform at the '/' character and save the array value in platform_split
    IFS='/' read -ra platform_split <<< "$platform"
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    OUTPUT_NAME=$OUTPUT_DIR/$PACKAGE_NAME'-'$GOOS'-'$GOARCH
    if [ "$GOOS" = "windows" ]; then
        OUTPUT_NAME+='.exe'
    fi
    echo [+] Building $PACKAGE_NAME for "$GOOS" "$GOARCH"
    env CGO_ENABLED="$ENABLE_CGO" GOOS="$GOOS" GOARCH="$GOARCH" go build "${GO_FLAGS[@]}" -o "$OUTPUT_NAME" .
done