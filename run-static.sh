#!/bin/bash
set -x

# remove docker image
docker rm -f static-compiler static-compiler

# build docker
cd static-compiler
docker build -f Dockerfile -t static-compiler .


cd ../src/plugin1
docker run --rm -v `pwd`:/plugin-source static-compiler plugin1.so


cd ../plugin2
docker run --rm -v `pwd`:/plugin-source static-compiler plugin2.so

cd ../..

pwd

ls src/plugin1 src/plugin2

mv {src/plugin1/plugin1.so,src/plugin2/plugin2.so} src/driver

# build driver
docker build -f Dockerfile -t static-driver .

# remove .so files
rm src/driver/*.so


# run docker
docker run static-driver