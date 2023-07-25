package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DeleteObject(s3Config aws.Config) error {
	// create a new session using the config above and profile
	goSession, err := session.NewSessionWithOptions(session.Options{
		Config:  s3Config,
		Profile: "filebase",
	})

	// check if the session was created correctly.
	if err != nil {
		return err
	}

	// create a s3 client session
	s3Client := s3.New(goSession)

	// create put object input
	deleteObjectInput := &s3.DeleteObjectInput{
		Bucket: aws.String("bucket-name"),
		Key:    aws.String("object-name"),
	}

	// get file
	_, err = s3Client.DeleteObject(deleteObjectInput)

	return err
}