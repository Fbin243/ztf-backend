package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func ConvertStringToUUID(fromValue string) (uuid.UUID, error) {
	u, err := uuid.Parse(fromValue)
	if err != nil {
		fmt.Println("Error:", err)
		return uuid.Nil, fmt.Errorf("failed to convert string to uuid: %w", err)
	}
	return u, nil
}
