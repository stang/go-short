# README

## Abstract

This project is a simple URL shortner.

## Pre-requisites

    # install dependencies
    dep ensure

## Development

    # start redis server with persistent storage
    docker run --rm --name redis -p 6379:6379 -v $(pwd)/data:/data -d redis:latest redis-server --appendonly yes

    # start dev server
    go run *.go

## Technical implementation

The APIs are wrote in `Go`, and the data is stored in `redis`.
The frontend is a static page that is bundled and served by the binary.

## TODO

[ ] add logging
[ ] add tests
[ ] add documentation
[ ] use config file
[ ] validate URL format
[ ] handle key colisions
[ ] add basic stats
