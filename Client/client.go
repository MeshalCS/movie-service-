// Alert: This code is in progress to implement requirements 
// ********************************************************	

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
	if err != nil {
		log.Fatalf("Failed to connect to gRPC service: %v", err)
	}
	defer conn.Close()

	// Create a client using the connection
	client := pb.NewMovieServiceClient(conn)
	


}

//fun LatestMovies for the user want to see the last movies
func GetLatestMovies(client pb.MovieServiceClient) string {
	// Call the LatestMovies method
	latestMoviesResponse, err := client.LatestMovies(context.Background(), &pb.LatestMoviesRequest{})
	fmt.Println("Latest Movies:")
	for _, movie := range latestMoviesResponse.Movies {
		fmt.Println("-", movie.Name)
	}
}




// ********************************************************	
// Alert: This code is in progress to implement requirements 
