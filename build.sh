#!/bin/bash

mkdir -p output

# Build for Linux
go build -o output/easywebstats .

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o output/easywebstats.exe .
