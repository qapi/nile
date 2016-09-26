#!/bin/bash

if [ $# -lt 1 ] ; then
    echo "Please provide parameter"
    exit 1
fi

dir="$1"
cd $dir
rm -rf bin
rm -rf deploy
mkdir deploy
cp index.js deploy/index.js

# compiling natively
docker run --rm -v "$PWD":/app -w /app qapi/gocker lambda

cp bin/aws-app deploy/main

cd deploy
zip -r lambda.zip main index.js

cp lambda.zip ~/Downloads/