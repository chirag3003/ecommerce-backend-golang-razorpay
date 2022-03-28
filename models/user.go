package models

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type User struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string               `json:"name"`
	Email     string               `json:"email"`
	Password  string               `json:"password"`
	Address   []primitive.ObjectID `json:"address,omitempty"`
	CreatedAt int64                `json:"createdAt"`
	UpdatedAt int64                `json:"updatedAt"`
}

type UserResponse struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}

func (user *User) GetJWT() (string, error) {
	//Generating JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["iat"] = time.Now().Unix()

	//Getting encoded JWT token
	t, err := token.SignedString([]byte(os.Getenv("SECRET")))

	return t, err
}

func (user *User) SetCreatedAt() {
	user.CreatedAt = time.Now().Unix()
}

func (user *User) SetUpdatedAt() {
	user.UpdatedAt = time.Now().Unix()
}

func (user *User) SetCreateDefaults() {
	user.SetUpdatedAt()
	user.SetCreatedAt()
	user.Address = []primitive.ObjectID{}
}

func (user *User) CheckPass(pass string) bool {
	hash := user.Password

	//Comparing Hashed and received password is correct
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func (user *User) CreateHash(pass string) bool {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return false
	}
	user.Password = string(hash)
	return true
}

func (user *User) Response() *UserResponse {
	response := &UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
	return response
}
