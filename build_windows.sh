#!/bin/bash
set -e

GOOS=windows GOARCH=386 go build
