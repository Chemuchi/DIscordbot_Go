#!/usr/bin/env bash
set -e

IMAGE_NAME="discord-bot"
CONTAINER_NAME="discord-bot"

echo "Building Docker image..."
docker build -t $IMAGE_NAME .

echo "Stopping existing container"
docker stop $CONTAINER_NAME 2>/dev/null || true

echo "Removing existing container"
docker rm $CONTAINER_NAME 2>/dev/null || true

echo "Starting Container"
docker run -d --name $CONTAINER_NAME --restart unless-stopped --env-file .env $IMAGE_NAME