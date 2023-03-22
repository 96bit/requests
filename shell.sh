#!/bin/zsh

git tag -d $1

git push origin :refs/tags/$1

git tag  $2

git push origin --tags
