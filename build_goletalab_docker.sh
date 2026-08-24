#!/usr/bin/env bash
set -euo pipefail
IMAGE_NAME=${1:-goletalab-ants}
DOCKER_BUILDKIT=1 docker build --platform linux/amd64 -f goletalab.Dockerfile -t "$IMAGE_NAME" .
