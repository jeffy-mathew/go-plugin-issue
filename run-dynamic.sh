#!/bin/bash
set -x

docker build -f Compiler-Dockerfile --build-arg BUILD_FILE_PATH=build-dynamic.sh -t dynamic-compiler .

cd src/plugin1
docker run --rm -v `pwd`:/plugin-source  dynamic-compiler plugin1.so


cd ../plugin2
docker run --rm -v `pwd`:/plugin-source  dynamic-compiler plugin2.so


cd ../plugin3
docker run --rm -v `pwd`:/plugin-source  dynamic-compiler plugin3.so

cd ../..

ls src/plugin1 src/plugin2 src/plugin3

mv {src/plugin1/plugin1.so,src/plugin2/plugin2.so,src/plugin3/plugin3.so} src/driver


# build driver
docker build -f Dockerfile -t dynamic-driver .

# remove .so files
rm src/driver/*.so

#run docker
docker run dynamic-driver
