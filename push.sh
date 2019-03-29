#!/bin/bash
function tag_and_push {
    if [ -n "$1" ] && [ -n "$IMAGE_NAME" ]; then
        echo "Pushing docker image to hub tagged as $IMAGE_NAME:$1"
        docker build -t $IMAGE_NAME:$1 -t $IMAGE_NAME -f Dockerfile .
        docker push $IMAGE_NAME
    fi
}
VERSION_TAG=v.$TRAVIS_BUILD_NUMBER
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
    tag_and_push $VERSION_TAG
