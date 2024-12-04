#!/bin/sh -e

DATE=$(date '+%F')

if [ -d .git ]; then
  VERSION=$(git describe --tags --always --dirty 2>/dev/null)
else
  VERSION=dev
fi

echo go build ...
go build -o timestamp -ldflags "-s -X main.version=$VERSION -X main.date=$DATE" .

echo go test ...
go test .

