package helpers

import (
	"fmt"
	"log"
	"os"
)

func VerifyENV() {
	variables := make(map[string]string)
	variables["MONGO_URI"] = os.Getenv("MONGO_URI")
	variables["SECRET"] = os.Getenv("SECRET")
	variables["S3_ACCESS_KEY"] = os.Getenv("S3_ACCESS_KEY")
	variables["S3_SECRET_KEY"] = os.Getenv("S3_SECRET_KEY")
	err := false
	for field := range variables {
		if variables[field] == "" {
			err = true
			fmt.Println("Set a value for the environment variable: ", field)
		}
	}
	if err {
		log.Fatalln("Pls set the above mentioned environment variables")
	}
}
