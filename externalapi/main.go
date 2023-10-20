package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	S3Bucket      = "crowdsvt"
	URLExpiration = 3 * time.Minute
)

type Response events.APIGatewayProxyResponse

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const filenameLength = 20

func getFilename() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, filenameLength)

	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func HandleRequest(ctx context.Context) (Response, error) {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	filename := getFilename()

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(S3Bucket),
		Key:    &filename,
	})
	urlStr, err := req.Presign(URLExpiration)

	if err != nil {
		return Response{
			StatusCode: 500,
			Body:       fmt.Sprintf("Failed to sign request: %v", err),
		}, nil
	}

	return Response{
		StatusCode: 200,
		Body:       urlStr,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
