#!/bin/bash

for f in *.go; do
    mkdir ${f%.go}
    mv $f ${f%.go}
done

