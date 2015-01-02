#!/usr/bin/env bash

if [ ! -f goget.sh ]; then
  echo 'go get must be run within its container folder' 1>&2
  exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

# config your go get
go get -u github.com/leowenyang/goELF

export GOPATH="$OLDGOPATH"

echo 'finished'
  
