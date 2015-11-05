#!/bin/sh -e
. ./env.sh

for d in splitter timestamp; do
  VFILE=src/rickb777/$d/version.go

  echo "// Generated automatically" > $VFILE
  echo "package main" >> $VFILE
  echo "const HgTip     = \"$(hg heads . |grep changeset: |head -1 |cut -f3 -d:)\"" >> $VFILE
  echo "const HgPath    = \"$(hg paths default)\"" >> $VFILE
  echo "const HgBranch  = \"$(hg branch)\"" >> $VFILE
  echo "const BuildDate = \"$(date '+%FT%T')\"" >> $VFILE
  echo "const BuildYear = \"$(date '+%Y')\"" >> $VFILE
  echo "const BuildUser = \"$USER\"" >> $VFILE
done

echo go install ...
go install rickb777/...

echo go test ...
go test rickb777/...

