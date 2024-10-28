package user

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/shaply/ProximityChat/Backend/types"
)

func TestUserServiceHandlers(test *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	test.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "",
			Password:  "password",
		}
		marshalled, _ := json.Marshal(&payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(ctx context.Context, id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(ctx context.Context, user *types.User) error {
	return nil
}
