#!/bin/sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run cmd/git-stack/main.go completion "$sh" >"completions/git-stack.$sh"
done