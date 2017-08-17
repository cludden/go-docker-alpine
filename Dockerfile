FROM golang:1.8-alpine

WORKDIR /go/src/app
COPY . .

ENV URL=https://google.com
CMD ["./app"]
