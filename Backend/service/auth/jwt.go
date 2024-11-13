package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shaply/ProximityChat/Backend/config"
	"github.com/shaply/ProximityChat/Backend/types"
	"github.com/shaply/ProximityChat/Backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateJWTToken(secret []byte, userID primitive.ObjectID) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	// fmt.Printf("Expiration: %v\n", expiration)
	// fmt.Printf("Time: %v\n", time.Now().Add(expiration).Unix())
	// fmt.Printf("userID: %v\n", userID.Hex())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID.Hex(),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type contextKey string

const UserKey contextKey = "userID"

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetTokenFromRequest(r)

		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		// Need to make it not get int and get ObjectID
		userID, err := primitive.ObjectIDFromHex(str)
		if err != nil {
			log.Printf("failed to convert userID to ObjectID: %v", err)
			permissionDenied(w)
			return
		}

		ctx1, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()
		u, err := store.GetUserByID(ctx1, userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		// Add the user to the context
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		// Call the function if the token is valid
		handlerFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func GetUserIDFromContext(ctx context.Context) primitive.ObjectID {
	userIDHex, ok := ctx.Value(UserKey).(string)
	if !ok {
		return primitive.NilObjectID
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return primitive.NilObjectID
	}

	return userID
}
