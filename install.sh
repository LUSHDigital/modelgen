#!/usr/bin/env bash

# base directory in which we will extract the files
base="src/github.com/LUSHDigital/modelgen"

# get the latest tarball url from github's api
url=$(curl -s https://api.github.com/repos/LUSHDigital/modelgen/releases/latest \
| grep "tarball_url" \
| cut -d : -f 2,3 \
| tr -d \"\,\ )

# get the tagged version
version=$(curl -s https://api.github.com/repos/LUSHDigital/modelgen/releases/latest \
| grep "tag_name" \
| cut -d : -f 2,3 \
| tr -d \"\,\ )

# download the tarball
curl -s -L ${url} --output release.tar.gz

# set the gopath
export GOPATH=$(pwd)
export PATH="$GOPATH/bin:$PATH"

# create the base folder
mkdir -p ${base}

# extract everything into it
tar -xvzf release.tar.gz -C ${base} --strip-components=1

# move into it
cd ${base}

# build
go build -ldflags "-X main.version=$version" -o /usr/local/bin/modelgen
go build -ldflags "-X main.version=$version" -o /usr/local/bin/modelgen

# move back to the root
cd -

# cleanup the files
rm release.tar.gz
rm -rf src
