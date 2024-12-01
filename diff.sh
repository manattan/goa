#!/bin/bash

git diff --cached | grep -v "^---" | grep -e '^+' -e '^-' | sed 's/^+++ b\//+++ \.\//'