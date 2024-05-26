FROM golang:1.22.0

RUN mkdir /stat
COPY . /stat
WORKDIR /stat

ENTRYPOINT [ "go", "run", "main.go" ]