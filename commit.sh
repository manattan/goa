#!/bin/bash

filePath="/tmp/$(uuidgen).diff"
resultPath="/tmp/$(uuidgen).txt"

bash ~/sandbox/goa/diff.sh > $filePath

~/sandbox/goa/main -f $filePath > $resultPath

git commit -m "$(cat $resultPath)"

rm -rf $filePath
