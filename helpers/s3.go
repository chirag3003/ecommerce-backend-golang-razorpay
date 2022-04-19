package helpers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

func ConnectS3() *session.Session {
	s3AccessKey := os.Getenv("S3_ACCESS_KEY")
	s3Secret := os.Getenv("S3_SECRET_KEY")
	s3Region := os.Getenv("S3_REGION")
	sess, err := session.NewSession(

		&aws.Config{
			Endpoint: aws.String(os.Getenv("S3_ENDPOINT")),
			Region:   aws.String(s3Region),
			Credentials: credentials.NewStaticCredentials(
				s3AccessKey,
				s3Secret,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}
