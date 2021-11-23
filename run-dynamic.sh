#!/bin/bash
set -x

# remove docker image
docker rm -f dynamic-compiler dynamic-compiler
# build compiler
cd dynamic-compiler
docker build -f Dockerfile -t dynamic-compiler .

cd ../src/plugin1
docker run --rm -v `pwd`:/plugin-source  dynamic-compiler plugin1.so


cd ../plugin2
docker run --rm -v `pwd`:/plugin-source  dynamic-compiler plugin2.so

cd ../..

pwd

ls src/plugin1 src/plugin2

mv {src/plugin1/plugin1.so,src/plugin2/plugin2.so} src/driver


# build driver
docker build -f Dockerfile -t dynamic-driver .

# remove .so files
rm src/driver/*.so

#run docker
docker run dynamic-driver
