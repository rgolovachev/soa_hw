FROM golang:1.22.0

RUN mkdir /auth
COPY . /auth
WORKDIR /auth

ENTRYPOINT [ "go", "run", "main.go" ]