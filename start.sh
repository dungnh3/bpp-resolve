#!/bin/sh

echo "starting.."

echo "build docker image"
make build-docker-image

echo "Migrate down database"
make
echo "done"