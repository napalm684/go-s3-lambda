package storage

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/napalm684/mytest/domain"
	"github.com/pkg/errors"
)

// S3Repository - AWS S3 Storage Repository
type S3Repository struct {
	service s3.S3
}

// NewS3Repository - constructs new S3 repository instance
func NewS3Repository() (*S3Repository, error) {
	service, err := getS3Service()

	if err != nil {
		return nil, errors.Wrap(err, "Unable to get an S3 service instance")
	}

	return &S3Repository{
		service: *service,
	}, nil
}

// GetObject - get object in bucket with key as bytes
func (repo *S3Repository) GetObject(request domain.Event) ([]byte, error) {
	var s3Event = request.(events.S3Event)
	var key = s3Event.Records[0].S3.Object.Key
	var bucket = s3Event.Records[0].S3.Bucket.Name

	log.Printf("Getting key/bucket: key: %v bucket: %v", key, bucket)

	r, err := repo.service.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"Unable to retrieve %v from bucket %v",
			key,
			bucket))
	}

	defer r.Body.Close() //TODO https://mijailovic.net/2017/05/09/error-handling-patterns-in-go/

	return ioutil.ReadAll(r.Body)
}

// getS3Service - Get an S3 service instance from AWS SDK
func getS3Service() (*s3.S3, error) {
	sess, err := session.NewSession()
	return s3.New(sess), err
}
