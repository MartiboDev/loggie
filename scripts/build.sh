#!/bin/bash

# Build the backend Docker image
docker build -t loggie-backend .

# Change to the web directory
cd web

# Build the frontend Docker image
docker build -t loggie-frontend .

# Change back to the root directory
cd ..