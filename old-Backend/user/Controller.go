package user

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/go-chi/chi"
	"github.com/threpio/mongoDBAtlasHackathon/backend/db"
	"github.com/threpio/mongoDBAtlasHackathon/backend/util"
	swizzle "github.com/threpio/mongoDBAtlasHackathon/backend/error-constants"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"net/mail"
	"time"
)

type AuthController struct {
	Router *chi.Router
	DB     *db.DB
}

func (ac *AuthController) GenerateSessionToken(email string) (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	token := util.Base64EncodeString(string(b))



	return token, nil
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

	if ac.IsEmailAlreadyInUse(util.Base64EncodeString(attempt.Email)) {
		return "", errors.New(swizzle.EmailExists)
	}

	// Create the User
	var user User
	user.UUID, _ = uuid.New()
	base64EncodedName := util.Base64EncodeString(attempt.Name)
	base64EncodedEmail := util.Base64EncodeString(attempt.Email)
	base64EncodedPassword := util.Base64EncodeString(attempt.Password)

	user.Name = base64EncodedName
	user.Email = base64EncodedEmail
	user.PasswordHash, _ = hashPassword(base64EncodedPassword)
	user.EmailConfirmed = false
	user.SuperAdmin = false
	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.UpdatedAt = time.Now().Format(time.RFC3339)
	err = ac.PutNewUserIntoDB(&user)
	// Send Confirmation email?
	//
	if err != nil {
		return "", err
	}
	// TODO: Session Token
	return "", nil
}
