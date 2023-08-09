package helper

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateOrderCode() string {
	uniqueId := uuid.New().String()
	code := strings.Replace(uniqueId, "-", "", -1)
	return code[:8]
}
