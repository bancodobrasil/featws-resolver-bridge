#!/bin/bash

docker-compose up -d mongo-express

go build -o resolver && ./resolver