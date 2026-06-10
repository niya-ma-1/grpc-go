package main

import (
	// "context"
	"log"
	"net"

	// "time"

	pb "github.com/niya-ma-1/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	// "go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	// defer client.Disconnect(context.Background())

	// if err := client.Ping(ctx, readpref.Primary()); err != nil {
	// 	log.Fatalf("failed to ping MongoDB: %v", err)
	// }

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
