#!/bin/sh

dep ensure -v &&
go build src/main.go
