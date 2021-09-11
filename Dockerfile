FROM golang:1.17.1-buster AS build-env

RUN apt-get update && \
    apt-get install -y \
    libopus-dev \
    build-essential \
    libssl-dev \
    pkg-config \
    curl \
    git

ENV LC_ALL=C.UTF-8

WORKDIR /go/src/github.com/

COPY  . .

RUN cd hi && pwd && go build

CMD [ "/bin/sh",  "-c", "/go/src/github.com/hi/hi" ]
