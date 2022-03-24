package models

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

func (user *User) GetJWT() (string, error) {
	//Generating JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Email
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	//Getting encoded JWT token
	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	return t, err
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
