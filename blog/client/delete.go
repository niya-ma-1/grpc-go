package main

import (
	"context"
	"log"

	pb "github.com/niya-ma-1/grpc-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog was invoked---")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error while deleting")
	}

	log.Println("Blog was deleted")
}
