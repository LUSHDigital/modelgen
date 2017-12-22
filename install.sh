#!/usr/bin/env bash
version=v1
git clone -b $version https://github.com/LUSHDigital/modelgen.git modelgen_src
go build -ldflags "-X main.version=$version" -o /usr/local/bin/modelgen ./modelgen_src
rm -rf modelgen_src
