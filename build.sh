#!/bin/env sh

cd shared-lib && \
    cmake . && \
    make && \
    echo "build success :)\n"