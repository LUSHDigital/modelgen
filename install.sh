#!/usr/bin/env bash

git clone -b v1 https://github.com/LUSHDigital/modelgen.git modelgen_src
go build -o /usr/local/bin/modelgen ./modelgen_src
rm -rf modelgen_src
