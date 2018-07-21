#! /bin/bash

local() {
    go test ./...
    python2 /usr/lib/google-cloud-sdk/bin/dev_appserver.py ./appeng/app.yaml
}

remote() {
    go test ./...
    yes Y | gcloud app deploy ./appeng/app.yaml
}

if [[ $(type -t $1) == function ]] 
then 
    $1
else
    echo "Error: ${1} is not a function"
    echo ""
    echo "The following functions do exist:"
    echo "local"
    echo "remote"
    echo ""
    echo "example usage: ./deploy.sh local"
fi

