#!/bin/sh
set -e
rm -rf manpages
mkdir manpages
go run cmd/git-stack/main.go man | gzip -c -9 >manpages/git-stack.1.gz