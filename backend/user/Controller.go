package user

import (
	"github.com/go-chi/chi"
	"github.com/threpio/mongoDBAtlasHackathon/backend/db"
	swizzle "github.com/threpio/mongoDBAtlasHackathon/backend/error-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"net/mail"
	"errors"
	"time"
)

type AuthController struct {
	Router *chi.Router
	DB *db.DB
}

func (ac *AuthController) GenerateSessionToken(email string) (string, error) {
	// TODO: Generate Session Token

	return "", nil
}

func (ac *AuthController) ValidateSessionToken(token string) bool {
	// TODO: Validate Session Token
	return false
}

//LoginUser logs in a user and returns a SessionToken
func (ac *AuthController) LoginUser(attempt AuthAttempt) (string, error) {
	// Check that email is an actual email
	_, err := mail.ParseAddress(attempt.Email)
	if err != nil {
		return "", errors.New(swizzle.EmailInvalid)
	}
	// Find User with that email and retrieve hashed password
	// TODO:

}

//RegisterUser registers a user and returns a SessionToken
func (ac *AuthController) RegisterUser(attempt NewUserAttempt) (string, error) {
	// Check that email is an actual email
	_, err := mail.ParseAddress(attempt.Email)
	if err != nil {
		return "", errors.New(swizzle.EmailInvalid)
	}
	// Check that password is at least 8 characters long
	if len(attempt.Password) < 8 {
		return "", errors.New(swizzle.PasswordTooShort)
	}
	// Check that the email is not already registered
	if ac.IsEmailAlreadyInUse(attempt.Email) {
		return "", errors.New(swizzle.EmailExists)
	}

	// Create the User
	var user User
	user.UUID, _ = uuid.New()
	user.Name = attempt.Name
	user.Email = attempt.Email
	user.PasswordHash, _ = hashPassword(attempt.Password)
	user.EmailConfirmed = false
	user.SuperAdmin = false
	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.UpdatedAt = time.Now().Format(time.RFC3339)
	err = ac.PutNewUserIntoDB(&user)

	// Send Confirmation email?
	//

	return "", errors.New(swizzle.EmailInvalid)
}