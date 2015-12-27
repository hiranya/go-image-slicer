#!/bin/bash

env GOOS=windows GOARCH=386 go build -o ./binaries/go-image-slicer-win-386.exe go-image-slicer.go
env GOOS=linux GOARCH=386 go build -o ./binaries/go-image-slicer-linux-386 go-image-slicer.go
env GOOS=darwin GOARCH=386 go build -o ./binaries/go-image-slicer-mac-386 go-image-slicer.go
