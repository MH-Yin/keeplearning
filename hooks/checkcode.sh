#!/bin/bash -e

# Copyright Greg Haskins All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

gopath=$GOPATH/src/github.com/MH-Yin/keeplearning
declare -a arr=(
    "$gopath/algorithm"
    "$gopath/context"
    "$gopath/pipeline"
    "$gopath/structure"
)

for i in "${arr[@]}"
do
    echo -e "\033[42m>>> Checking code under $i/\033[0m"

    echo "Checking with gofmt"
    OUTPUT="$(gofmt -l -s $i | grep -v testdata/ || true)"
    if [[ $OUTPUT ]]; then
        echo -e "\033[41mThe following files contain gofmt errors\033[0m"
        echo "$OUTPUT"
        echo "The gofmt command 'gofmt -l -s -w' must be run for these files"
        exit 1
    fi

    echo "Checking with go vet"
    OUTPUT="$(go vet -composites=false $i/)"
    if [[ $OUTPUT ]]; then
        echo -e "\033[41mThe following files contain go vet errors\033[0m"
        echo $OUTPUT
        exit 1
    fi
done
