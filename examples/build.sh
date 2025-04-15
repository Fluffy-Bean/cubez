#!/usr/bin/env bash

go build -tags ballistic -o ballistic ballistic.go exampleapp.go
go build -tags cubedrop -o cubedrop cubedrop.go exampleapp.go
