package services

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadFileFromCtx(c *fiber.Ctx) (string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		return "", err
	}
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	filePath := fmt.Sprintf("public/%s.%s", filename, fileExt)
	err = c.SaveFile(file, filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
