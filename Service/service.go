package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"


	"google.golang.org/grpc"

	//pb "movie_service/porto/protobuf.proto"
)

const (
	port = ":50051"
	apiKey = "ba72b135c4f80393492d0bf1d60d08fc"
	baseURL = "https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s"

)

type movieServer struct{
	pb.UnimplementedMovieServiceServer
	db     *database.DB
	favourites map[int64]*Movie
	// client *api.Client
}


// Function for Latest Movies:
// I wrote it but there is bugs, so I removed as you say before do what you can do it,
//because I don’t wnat to write something I know it’s not correct

// Function for Search Movies:
func (s *MovieService) SearchMovies(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	url := fmt.Sprintf(baseURL, apiKey, req.Query)

	resp, err := http.Get(url)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to fetch data from TheMovieDB API")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to read response from TheMovieDB API")
	}

	var searchResult struct {
		Results []*Movie `json:"results"`
	}
	if err := json.Unmarshal(body, &searchResult); err != nil {
		return nil, status.Error(codes.Internal, "failed to parse response from TheMovieDB API")
	}

	movies := make([]*Movie, 0, len(searchResult.Results))
	for _, movie := range searchResult.Results {
		movies = append(movies, movie)
	}

	return &SearchResponse{Movies: movies}, nil
}





// Function Add Movie from Favorites:
func (s *MovieService) AddMovieToFavorites(ctx context.Context, in *pb.AddMovieToFavoritesRequest) (*pb.AddMovieToFavoritesResponse, error) {
	err := s.db.AddMovieToFavorites(in.MovieId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add movie to favorites: %v", err)
	}

	return &pb.AddMovieToFavoritesResponse{Success: true}, nil
}


// Function Remove Movie from Favorites:
func (s *MovieService) RemoveMovieFromFavourites(ctx context.Context, req *RemoveMovieFromFavoritesRequest) (*empty.Empty, error) {
    userID := req.GetUserId()
    movieID := req.GetMovieId()

    // First, check if the user already has the movie in their favorites list
    favourite, err := s.db.GetFavouriteMovie(userID, movieID)
    if err != nil {
        return nil, err
    }
    if favourite == nil {
        return nil, status.Errorf(codes.NotFound, "Movie not found in favorites")
    }

    // If the movie is in the favorites list, remove it
    if err := s.db.RemoveFavouriteMovie(userID, movieID); err != nil {
        return nil, err
    }

    return &empty.Empty{}, nil
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
