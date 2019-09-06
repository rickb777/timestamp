#!/bin/sh -e

VFILE=version.go

echo "// Generated automatically - do not edit!" > $VFILE
echo "package main" >> $VFILE
echo "" >> $VFILE
echo "const Version   = \"$(git describe --tags --always 2>/dev/null)\"" >> $VFILE
echo "const BuildDate = \"$(date '+%FT%T')\"" >> $VFILE
echo "const BuildYear = \"$(date '+%Y')\"" >> $VFILE
echo "const BuildUser = \"$USER\"" >> $VFILE

echo go install ...
go build -o timestamp .

echo go test ...
go test .

