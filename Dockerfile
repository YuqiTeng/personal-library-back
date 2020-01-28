FROM golang:1.13-alpine3.10

EXPOSE 8080

WORKDIR /go/src/app
COPY . .

RUN apk add --no-cache git mercurial

RUN go get -d -v ./...
RUN go install -v ./...

RUN apk del git mercurial
CMD ["app"]