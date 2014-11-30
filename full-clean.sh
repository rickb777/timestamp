#!/bin/sh -e
. ./env.sh

make -C assets clean
rm -rf bin pkg roster3.zip logs *.log
rm -rf src/github.com src/launchpad.net src/code.google.com src/golang.org src/mock
tar zxvf dependencies.tar.gz
go clean ./...
