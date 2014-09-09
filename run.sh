#!/bin/bash

export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export PATH=$GOPATH/bin:$PATH

go $1 $2