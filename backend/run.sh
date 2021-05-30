#!/bin/bash

cd scripts/codegen
cbgen -n startpage -v dev
mv *.go ../../
cd ../../
cd cmd/startpage
go build
mv startpage ../../
cd ../../
./startpage -s ../frontend/build "$@"