package helper

import (
	"fmt"
	"strings"

	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hosseinmirzapur/parsian-backend/config"
)

func UploadCtxFile(c *fiber.Ctx) (string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		return "", err
	}
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	filePath := fmt.Sprintf("public/%s.%s", filename, fileExt)

	// Save FIle locally
	err = c.SaveFile(file, "public")
	if err != nil {
		return "", err
	}
	return filePath[7:], nil
}

func UploadToAWS(c *fiber.Ctx) (string, error) {
	// load AWS config & Create uploader
	client := config.GetClient()
	uploader := manager.NewUploader(client)

	// Create Filename
	file, err := c.FormFile("image")
	if err != nil {
		return "", err
	}
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	filePath := fmt.Sprintf("%s.%s", filename, fileExt)

	// Upload File
	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("parsian"),
		Key:    aws.String(filePath),
		Body:   f,
	})

	if err != nil {
		return "", fmt.Errorf("error uploading file: %w", err)
	}

	return result.Location, nil
}
