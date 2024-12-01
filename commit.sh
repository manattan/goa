#!/bin/bash

filePath="/tmp/$(uuidgen).diff"
resultPath="/tmp/$(uuidgen).txt"

git diff --cached | grep -v "^---" | grep -e '^+' -e '^-' | sed 's/^+++ b\//+++ \.\//' > $filePath

~/sandbox/goa/main -f $filePath > $resultPath

git commit -m "$(cat $resultPath)"

rm -rf $filePath
rm -rf $resultPath
