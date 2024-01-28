package main

import (
	"fmt"
	"net/http"

	"github.com/DatNgo-dev/rssaggregator/internal/database"
)

// func (apiCfg apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request, user database.User) {
// 	type parameters struct {
// 		Name string `json:"name"`
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}

// 	// %params is a pointer. Therefore this params structure we 'instantiated' will have its properties changed. Its stateful.
// 	// Golang passes value by value and not by referenced
// 	err := decoder.Decode(&params)

// 	if err != nil {
// 		respondWithError(w, 400, fmt.Sprintf("Error passing json: %v", err))
// 		return
// 	}

// 	newUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
// 		ID: uuid.New(),
// 		CreatedAt: time.Now().UTC(),
// 		UpdatedAt: time.Now().UTC(),
// 		Name: params.Name,
// 	})
// 	if err != nil {
// 		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
// 		return
// 	}

// 	respondWithJSON(w, 201, databaseUserToUser(newUser))
// }

func (apiCfg apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}