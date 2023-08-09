package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {
	code := uuid.New()
	newCode := strings.Replace(code.String(), "-", "", -1)
	fmt.Println(newCode[:6])

	// Testing Order Code generation
}
