FROM golang:1.22.0

RUN mkdir /posts
COPY . /posts
WORKDIR /posts

ENTRYPOINT [ "go", "run", "main.go" ]