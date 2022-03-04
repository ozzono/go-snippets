#!/bin/bash

# pull(){
#     repo=$1
#     here=$2
#     echo "cloning $repo"
#     git clone <repo_path>/$repo
#     echo "cloned $repo"
#     cd $here/$repo
#     echo "fetching packages"
#     go mod tidy
#     go mod download
#     echo "fetched packages"
#     cd $here
# }

go_files=$(ls *.go)

for f in *.go; do
    mkdir ${f%.go}
    mv $f ${f%.go}
done

