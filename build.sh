#!/bin/sh -e
. ./env.sh

mkdir -p src/rickb777/timestamp/util
VFILE=src/rickb777/timestamp/util/version.go

echo "// Generated automatically" > $VFILE
echo "package util" >> $VFILE
echo "const HgTip     = \"$(hg heads . |grep changeset: |head -1 |cut -f3 -d:)\"" >> $VFILE
echo "const HgPath    = \"$(hg paths default)\"" >> $VFILE
echo "const HgBranch  = \"$(hg branch)\"" >> $VFILE
echo "const BuildDate = \"$(date '+%FT%T')\"" >> $VFILE
echo "const BuildYear = \"$(date '+%Y')\"" >> $VFILE
echo "const BuildUser = \"$USER\"" >> $VFILE

echo go install ...
go install rickb777/...

echo go test ...
go test rickb777/...

