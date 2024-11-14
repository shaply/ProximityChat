// This package contains utility functions that are used across the application.

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is missing")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// Makes error reporting consistent.
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

// Make email storage and getting consistent
func FixEmail(email string) string {
	return strings.ToLower(email)
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")
	tokenPath := mux.Vars(r)["JWTToken"]

	if tokenAuth != "" {
		fmt.Println("Token from tokenAuth: ", tokenAuth)
		return tokenAuth
	}

	if tokenQuery != "" {
		fmt.Println("Token from tokenQuery: ", tokenQuery)
		return tokenQuery
	}

	if tokenPath != "" {
		fmt.Println("Token from tokenPath: ", tokenPath)
		return tokenPath
	}

	return ""
}
