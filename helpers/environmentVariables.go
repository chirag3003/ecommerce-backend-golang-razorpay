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
