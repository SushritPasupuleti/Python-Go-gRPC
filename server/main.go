package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "server/pb/query"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewResponseServiceClient(conn)

	response, err := c.GetResponse(context.Background(), &pb.RequestBody{Message: "Hello World!"})

	if err != nil {
		log.Fatalf("Error when calling GetResponse: %s", err)
	}

	log.Printf("Response from server: %s with status: %d", response.Message, response.Status)
}
