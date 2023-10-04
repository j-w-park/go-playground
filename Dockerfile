FROM golang

WORKDIR /app

RUN apt-get update
RUN apt-get install -y git
RUN git config --global --add safe.directory $PWD
