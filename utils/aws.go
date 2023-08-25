package utils

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/hosseinmirzapur/parsian-backend/config"
)

func UploadToAWS(filepath string) (string, error) {
	// Load AWS Config
	client := config.GetClient()
	uploader := manager.NewUploader(client)

	// Build filename
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	uniqueId := uuid.New()
	fileMidName := strings.Replace(uniqueId.String(), "-", "", -1)
	fileFullname := fmt.Sprintf("parsian-%s.xlsx", fileMidName)

	// Upload File to AWS
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("parsian"),
		Key:    aws.String(fileFullname),
		Body:   file,
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file, %s", err.Error())
	}

	return result.Location, nil

}