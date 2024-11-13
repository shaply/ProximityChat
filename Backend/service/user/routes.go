package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/shaply/ProximityChat/Backend/config"
	"github.com/shaply/ProximityChat/Backend/service/auth"
	"github.com/shaply/ProximityChat/Backend/types"
	"github.com/shaply/ProximityChat/Backend/utils"
)

type Handler struct {
	store types.UserStore // By using interface, it's much easier to test
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %+v", errors))
		return
	}

	// check if the user exists
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	payload.Email = utils.FixEmail(payload.Email)
	user, err := h.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("email or password is incorrect"))
		return
	}

	fmt.Printf("User: %+v\n", user)

	// check if the password is correct
	if !auth.CheckPasswordHash(payload.Password, user.Password) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("email or password is incorrect"))
		return
	}

	// generate a JWT token
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWTToken(secret, user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Placeholder
	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// fmt.Print("Payload: ", payload)

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %+v", errors))
		return
	}

	// fmt.Print("Payload is valid")

	// check if the user exists
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// fmt.Print("Beginning to check if user exists")

	payload.Email = utils.FixEmail(payload.Email)
	_, err := h.store.GetUserByEmail(ctx, payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// fmt.Print("User does not exist")

	// if not, create the new user
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(ctx, &types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// fmt.Print("Registered user")

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user created"})
}
