#!/bin/bash

filePath="/tmp/$(uuidgen).diff"
resultPath="/tmp/$(uuidgen).txt"

bash ./diff.sh > $filePath

./main -f $filePath > $resultPath

git commit -m "$(cat $resultPath)"

rm -rf $filePath
