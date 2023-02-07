package main

import (
	"context"
	"fmt"

	//pb "movie" // Import the generated protobuf code

	"google.golang.org/grpc"
)

func main() {
	address := "localhost:50051" // Address of the gRPC service

	// Connect to the gRPC service
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC service: %v", err)
		return
	}
	defer conn.Close()

	// Create a client using the connection
	client := pb.NewMovieServiceClient(conn)
	
	
	//*************************************
	// to test all methods:
	
	//getLatestMovies(client)
	//searchMovies(client, "Ex Machina")
	//addMovieToFavourites(client, 603)
	//removeMovieFromFavourites(client, 603)
	//*************************************
	


}

//fun LatestMovies for the user want to see the last movies
func getLatestMovies(client pb.MovieServiceClient) {
	res, err := client.GetLatestMovies(context.Background(), &empty.Empty{})
	if err != nil {
		fmt.Println("Error while getting latest movies.", err)
		return
	}

	fmt.Println("Latest Movies:")
	for _, movie := range res.Movies {
		fmt.Println("Name:", movie.Name)
		fmt.Println("Poster:", movie.Poster)
		fmt.Println("Description:", movie.Description)
		fmt.Println("Genres:", movie.Genres)
	}
}

//fun SearchMovies for the user want to Search the movies
func searchMovies(client pb.MovieServiceClient, query string) {
	res, err := client.SearchMovies(context.Background(), &SearchMovieRequest{Query: query})
	if err != nil {
		fmt.Println("Error while searching movies.", err)
		return
	}

	fmt.Println("Search Results:")
	for _, movie := range res.Movies {
		fmt.Println("Name:", movie.Name)
		fmt.Println("Poster:", movie.Poster)
		fmt.Println("Description:", movie.Description)
		fmt.Println("Genres:", movie.Genres)
	}
}

//fun AddMovieToFavorites for the user want to add the movie to Favorites
func addMovieToFavourites(client pb.MovieServiceClient, movieID int32) {
	_, err := client.AddMovieToFavourites(context.Background(), &AddMovieRequest{MovieId: movieID})
	if err != nil {
		fmt.Println("Error while adding movie to favourites.", err)
		return
	}

	fmt.Println("Movie added to favourites successfully.")
}

//fun removeMovieFromFavourites for the user want to remove the movie from Favorites
func removeMovieFromFavourites(client pb.MovieServiceClient, movieID int32) {
	_, err := client.RemoveMovieFromFavourites(context.Background(), &RemoveMovieRequest{MovieId: movieID})
	if err != nil {
		fmt.Println("Error while removing movie from favourites.", err)
		return
	}

	fmt.Println("Movie removed from favourites successful.")
}

