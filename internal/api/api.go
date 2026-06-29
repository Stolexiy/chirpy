package api

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"
)

func HandleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func HandleValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Parameters is not valid")
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	type Body struct {
		CleanedBody string `json:"cleaned_body"`
	}

	respBody := Body{
		CleanedBody: cleanupBody(params.Body),
	}

	jsonResponse(w, 200, respBody)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type errorResp struct {
		Error string `json:"error"`
	}

	respBody := errorResp{
		Error: msg,
	}

	jsonResponse(w, code, respBody)
}

func respondValid(w http.ResponseWriter) {
	type valid struct {
		Valid bool `json:"valid"`
	}

	respBody := valid{
		Valid: true,
	}

	jsonResponse(w, 200, respBody)
}

func jsonResponse(w http.ResponseWriter, code int, respBody any) {
	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshaling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func cleanupBody(body string) string {
	blackLists := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}

	words := strings.Split(body, " ")
	for i, word := range words {
		if slices.Contains(blackLists, strings.ToLower(word)) {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}
