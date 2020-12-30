# micro reverse proxy

[
![commit](https://img.shields.io/github/last-commit/0ex/micro-containers?label=commit)
![stars](https://img.shields.io/github/stars/0ex/micro-containers)
](https://github.com/0ex/micro-containers)
[
![pulls](https://img.shields.io/docker/pulls/0ex0/micro-reverse-proxy?label=pulls)
](https://hub.docker.com/r/0ex0/micro-reverse-proxy)

## usage

    docker run -p 800:80 ghcr.io/0ex/micro-reverse-proxy https://0ex.github.io
    
    curl -v http://localhost:800

## build

    make build

## develop

    go run main.go https://0ex.github.io
