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

	// Call the LatestMovies method
	latestMoviesResponse, err := client.LatestMovies(context.Background(), &pb.LatestMoviesRequest{})
	if err != nil {
		log.Fatalf("Failed to get latest movies: %v", err)
	}
	fmt.Println("Latest Movies:")
	for _, movie := range latestMoviesResponse.Movies {
		fmt.Println("-", movie.Name)
	}

	// Call the SearchMovies method
	searchMoviesResponse, err := client.SearchMovies(context.Background(), &pb.SearchMoviesRequest{
		Query: "Matrix",
	})
	if err != nil {
		log.Fatalf("Failed to search movies: %v", err)
	}
	fmt.Println("\nSearch Results:")
	for _, movie := range searchMoviesResponse.Movies {
		fmt.Println("-", movie.Name)
	}

	// Call the AddMovieToFavorites method
	addMovieToFavoritesResponse, err := client.AddMovieToFavorites(context.Background(), &pb.AddMovieToFavoritesRequest{
		MovieId: 123,
	})
	if err != nil {
		log.Fatalf("Failed to add movie to favorites: %v", err)
	}
	if addMovieToFavoritesResponse.Success {
		fmt.Println("\nMovie added to favorites successfully.")
	} else {
		fmt.Println("\nFailed to add movie to favorites.")
	}

	// Call the RemoveMovieFromFavorites method
	removeMovieFromFavoritesResponse, err := client.RemoveMovieFromFavorites(context.Background(), &pb.RemoveMovieFromFavoritesRequest{
		MovieId: 123,
	})
	if err != nil {
		log.Fatalf("Failed to remove movie from favorites: %v", err)
	}
	if removeMovieFromFavoritesResponse.Success {
		fmt.Println("\nMovie removed from favorites successfully.")
	} else {
		fmt.Println("\nFailed to remove movie from favorites.")
	}

	// Call the MovieDetails method
	movieDetailsResponse, err := client.MovieDetails(context.Background(), &pb.MovieDetailsRequest{
		MovieId: 123,
	})
