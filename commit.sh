#!/bin/bash

filePath="/tmp/$(uuidgen).diff"
resultPath="/tmp/$(uuidgen).txt"

bash ./diff.sh > $filePath

go run main.go -f $filePath > $resultPath

git commit -m "$(cat $resultPath)"

rm -rf $filePath
