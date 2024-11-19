#!/bin/bash
set -e
go build -trimpath -ldflags "-w -s"
./simple-short-link-go
