FROM golang:1.22.0

RUN mkdir /auth
COPY . /auth
WORKDIR /auth

RUN go get -d -v ./...

ENTRYPOINT [ "go", "run", "main.go" ]