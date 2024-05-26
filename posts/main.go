package main

import (
	"log"
	"net"
	posts "posts/posts_grpc"
	postspb "posts/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:11777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	posts_service, err := posts.NewServer()
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	postspb.RegisterPostsServiceServer(grpcServer, posts_service)

	log.Println("posts service started")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("posts service failed")
	}
}
