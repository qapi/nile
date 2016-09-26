#!/bin/bash

if [ $# -lt 1 ] ; then
    echo "Please provide parameter"
    exit 1
fi

dir="$1"
cd $dir

rm -rf vendor

# building vendor dependencies
docker run --rm -it -v "$PWD":/app -w /app qapi/gocker vendor