#!/bin/bash

go install .

rm -rf test-directory
mkdir test-directory


pushd test-directory
mraziz init ../test-template
popd