#!/bin/bash

set -xe

plugin_name=$1
PLUGIN_BUILD_PATH="/go/src/${plugin_name%.*}"

function usage() {
    cat <<EOF
To build a plugin:
      $0 <plugin_name>

EOF
}

mkdir -p $PLUGIN_BUILD_PATH

if [ -z "$plugin_name" ]; then
    usage
    exit 1
fi

yes | cp -r $PLUGIN_SOURCE_PATH/* $PLUGIN_BUILD_PATH || true

cd $PLUGIN_BUILD_PATH
# if plugin has go.mod
[ -f go.mod ] && [ ! -d vendor ] && go mod vendor
rm go.mod


go build -buildmode=plugin -o $plugin_name \
    && mv $plugin_name $PLUGIN_SOURCE_PATH
