# go-docker-alpine
Building minimal docker containers for go applications. Original application code borrowed from this [post](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/), however this project uses [alpine](https://alpinelinux.org/) instead of `scratch`.

## Getting Started
Build the binary
```shell
$ docker run --rm -v "$PWD":/go/src/app -w /go/src/app golang:1.8-alpine go build
```

Build the container
```shell
$ docker build -t go-docker-apline .
$ docker run -it go-docker-alpine
```

## License
**UNLICENSED**
