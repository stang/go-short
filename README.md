# README

## Abstract

`go-short` a simple URL shortner.

## Pre-requisites

    # install dependencies
    dep ensure

## Development (local)

    # start redis server with persistent storage
    docker run --rm --name redis -p 6379:6379 -v $(pwd)/data:/data -d redis:latest redis-server --appendonly yes

    # start dev server
    go run *.go

## Deploy demo (heroku)

    # assuming you already have all the heroku-cli setup
    # see: https://devcenter.heroku.com/articles/getting-started-with-go#introduction

    # create free redis instance (data not persisted)
    heroku addons:create heroku-redis:hobby-dev

    # this creates a redis instance and expose it's config in the REDIS_URL environment variable
    heroku config

    # (optional) needs to be done only once, if you want a custom domain name
    heroku domains:add go.stang.sh

    # deploy
    git push heroku master

## Technical implementation

The APIs are wrote in `Go`, and the data is stored in `redis`.
The frontend is a static page that is bundled and served by the binary.

## TODO

* [ ] add logging
* [ ] add tests
* [ ] add documentation
* [ ] validate URL format
* [ ] handle key colisions
* [ ] add basic stats
