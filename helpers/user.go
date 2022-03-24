package helpers

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"os"
)

func validateSignedMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(os.Getenv("SECRET")), nil
}

type UserJWT struct {
	ID    string
	Email string
}

func VerifyJWT(tokenString string) (bool, *UserJWT) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	_, err := jwt.ParseWithClaims(tokenString, claims, validateSignedMethod)
	if err != nil {
		return false, nil
	}
	email := claims["email"].(string)
	id := claims["id"].(string)

	return true, &UserJWT{
		Email: email,
		ID:    id,
	}
}

func ValidateUserRegisterInput(data *models.User) Errors {
	err := Errors{}
	if !govalidator.IsEmail(data.Email) {
		err.Add("email", "Email is not valid")
	}
	err.CheckLen(data.Password, 5, "password")
	if !err.IsValid() {
		return err
	}
	return nil
}

func ParseUser(ctx *fiber.Ctx) *models.User {
	return ctx.Locals("user").(*models.User)
}
