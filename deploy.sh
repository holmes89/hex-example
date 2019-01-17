#!/usr/bin/env bash
TAG=$(git describe --abbrev=0 --tags)

GOOS=linux go build -o main cmd/ticketAPI/main.go
zip main.zip cmd/ticketAPI/main

aws s3 cp main.zip s3://hex-lambda/$TAG/main.zip

cd terraform/prod/

terraform apply -var "app_version=$TAG" -auto-approve

cd ../../
rm -rf main.zip
