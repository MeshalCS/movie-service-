// Alert: The code is in progress to implement requirements 


package main

import (
	"context"
	"fmt"
	"log"

	//pb "movie" // Import the generated protobuf code

	"google.golang.org/grpc"
)

func main() {
	address := "localhost:50051" // Address of the gRPC service

	// Connect to the gRPC service
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	

	
	// Alert: The code is in progress to implement requirements 
