package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"golang.org/x/crypto/bcrypt"
)

// TODO: This shizzle

type User struct {
	UUID           uuid.UUID `json:"uuid"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	EmailConfirmed bool      `json:"email_confirmed"`
	PasswordHash   string    `json:"password_hash"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
	SuperAdmin     bool      `json:"super_admin"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (ac *AuthController) PutNewUserIntoDB(user *User) error {
	err := ac.DB.InsertOne("users", user)
	if err != nil {
		return err
	}
	return nil
}

func (ac *AuthController) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	singleResult := ac.DB.FindOne("users", bson.M{"email": email})
	if singleResult.Err() != nil && singleResult.Err() != mongo.ErrNoDocuments {
		return nil, singleResult.Err()
	}
	if singleResult.Err() == mongo.ErrNoDocuments {
		return nil, singleResult.Err()
	}
	err := singleResult.Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ac *AuthController) IsEmailAlreadyInUse(email string) bool {
	user, err := ac.GetUserByEmail(email)
	if err != nil {
		return false
	}
	if user.Email == email {
		return true
	}
	return false
}

