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

COPY  . .

# CMD [ "/bin/sh",  "-c", "cargo run" ]
