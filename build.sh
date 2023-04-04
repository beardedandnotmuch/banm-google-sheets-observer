#!/bin/bash

TAG=${TAG:-latest}

while [ $# -gt 0 ]; do
    case "$1" in
        -t|--tag)
            shift
            TAG=$1
            shift
            ;;
        --no-push)
            shift
            NOPUSH=1
            ;;
        *)
            break
            ;;
    esac
done

api() {
    DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker build -f Dockerfile.multistage -t "registry.beardedandnotmuch.com/banm-google-sheets-observer:$TAG" .
    [ -z $NOPUSH ] && docker push "registry.beardedandnotmuch.com/banm-google-sheets-observer:$TAG"
}

case $1 in
    api)
        api
        ;;
    *)
        api
        ;;
esac
