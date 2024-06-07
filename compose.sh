#!/bin/bash

# Run the build.sh script
./scripts/build.sh

# Compose docker containers
docker compose -p loggie-stack -f ./docker-compose.yml up -d