#!/bin/bash

cd plugin
go build -buildmode=plugin -o plugin.so plugin.go
cp plugin.so ../driver
cd ../driver
go build main.go