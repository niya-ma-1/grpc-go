package main

import (
	"go.mongodb.org/mongo-driver/v2/bson"

	pb "github.com/niya-ma-1/grpc-go/blog/proto"
)

type BlogItem struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	AuthorID string        `bson:"author_id"`
	Title    string        `bson:"title"`
	Content  string        `bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}
