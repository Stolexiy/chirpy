package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	type UsersRequest struct {
		Email string `json:"email"`
	}

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, 400, "Parameters is not valid")
		return
	}

	var body *UsersRequest
	err = json.Unmarshal(rawBody, body)
	if err != nil {
		respondWithError(w, 400, "Something went wrong")
	}

}
