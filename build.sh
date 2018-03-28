#!/bin/bash

gofmt -s -w *.go

go test

go install
