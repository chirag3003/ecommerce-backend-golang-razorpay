package helpers

import (
	"fmt"
	"log"
	"os"
)

func VerifyENV() {
	variables := []string{
		"PORT",
		"MONGO_URI",
		"MONGO_DB",
		"SECRET",
		"S3_ACCESS_KEY",
		"S3_SECRET_KEY",
		"S3_REGION",
		"S3_BUCKET",
		"S3_ENDPOINT",
		"RAZORPAY_KEY",
		"RAZORPAY_SECRET",
		"RAZORPAY_WEBHOOK_SECRET",
	}

	err := false
	for _, field := range variables {
		if os.Getenv(field) == "" {
			err = true
			fmt.Println("Set a value for the environment variable: ", field)
		}
	}
	if err {
		log.Fatalln("Pls set the above mentioned environment variables")
	}
}
