package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DatNgo-dev/rssaggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg apiConfig) handlerGetAllFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := apiCfg.DB.GetAllFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feedFollow: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollow))
}

func (apiCfg apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error passing json: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create a feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDString := chi.URLParam(r, "feedFollowID")

	feedFollowID, err :=  uuid.Parse(feedFollowIDString)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID: feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}