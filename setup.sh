#!/bin/bash

cd plugin
go build -buildmode=plugin -o plugin1.so plugin.go
go build -buildmode=plugin -o plugin2.so plugin.go
cp {plugin1.so,plugin2.so} ../driver
cd ../driver
go build main.go
./main.go