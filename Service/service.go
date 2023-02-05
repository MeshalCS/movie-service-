package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	//pb "movie_service/porto/protobuf.proto"
)

const (
	port = ":50051"
)

type movieServer struct{
	pb.UnimplementedMovieServiceServer
	db     *database.DB
	// client *api.Client
}


// Function for Latest Movies:
func (s *MovieService) LatestMovies(ctx context.Context, in *pb.LatestMoviesRequest) (*pb.LatestMoviesResponse, error) {
	movies, err := s.client.GetLatestMovies(in.Page)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get latest movies: %v", err)
	}

	res := &pb.LatestMoviesResponse{}
	for _, movie := range movies {
		res.Movie = append(res.Movie, &pb.Movie{
			id:          int32(movie.id),
			Title:       movie.Title,
			poster_path:  movie.poster_path,
			description:    movie.description,
			genres:      movie.genres,
		})
	}

	return res, nil
}

// Function for Search Movies:
func (s *MovieService) SearchMovies(ctx context.Context, in *pb.SearchMoviesRequest) (*pb.SearchMoviesResponse, error) {
	movies, err := s.client.SearchMovies(in.Query, in.Page)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to search for movies: %v", err)
	}

	res := &pb.SearchMoviesResponse{}
	for _, movie := range movies {
		res.Movie = append(res.Movie, &pb.Movie{
			id:          int32(movie.id),
			Title:       movie.Title,
			poster_path:  movie.poster_path,
			description: movie.description,
			genres:      movie.genres,
		})
	}

	return res, nil
}




// Function Add Movie from Favorites:
func (s *server) AddMovieToFavorites(ctx context.Context, in *pb.AddMovieToFavoritesRequest) (*pb.AddMovieToFavoritesResponse, error) {
	err := s.db.AddMovieToFavorites(in.MovieId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add movie to favorites: %v", err)
	}

	return &pb.AddMovieToFavoritesResponse{Success: true}, nil
}


// Function Remove Movie from Favorites:
func (s *server) RemoveMovieFromFavorites(ctx context.Context, in *pb.RemoveMovieFromFavoritesRequest) (*pb.RemoveMovieFromFavoritesResponse, error) {
	err := s.db.RemoveMovieFromFavorites(in.MovieId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to remove movie from favorites: %v", err)
	}

	return &pb.RemoveMovieFromFavoritesResponse{Success: true}, nil
}

// Function for Movie Details:
func (s *MovieService) MovieDetails(ctx context.Context, req *MovieDetailsRequest) (*MovieDetailsResponse, error) {
	// I will retrieve the latest movies from the The Movie DB API, when I finish the implementation code
	movie, err := getMovieDetails(req.MovieID)
	if err != nil {
		return nil, err
	}

	// Build and return the response
	//in progress to implement
	resp := &MovieDetailsResponse{
		Movie: movie,
	}
	return resp, nil
}
