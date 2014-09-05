#!/bin/bash

export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export PATH=$GOPATH/bin:$PATH

echo $GOPATH

CASSANDRA_CONTAINER_IP=`docker inspect cass1 |grep IPA|cut -f4 -d'"'`
# docker run -ti --rm poklet/cassandra cqlsh $CASSANDRA_CONTAINER_IP -f ~/casper/golang-learn/init-cassandra.cql

