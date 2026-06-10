package main

import (
	"context"
	"log"

	pb "github.com/niya-ma-1/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---update blog was invoked---")

	newBlog := &pb.Blog{
		Id: id, 
		AuthorId: "Not Niya",
		Title: "A new title",
		Content: "Content of the first blog with some additions",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating %v\n", err)
	}

	log.Println("Blog was updated!")
}