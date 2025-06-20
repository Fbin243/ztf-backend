package utils

import (
	"fmt"
	"strconv"
)

func ConvertStringToUInt(fromValue string) (uint, error) {
	toValue, err := strconv.ParseUint(fromValue, 10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, fmt.Errorf("failed to convert string to uint: %w", err)
	}

	return uint(toValue), nil
}
