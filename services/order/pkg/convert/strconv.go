package convert

import (
	"log"
	"strconv"
)

func MustConvStrToInt(idStr string) int64 {
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		log.Printf("Error convert %v to int64", idStr)
		return 0
	}

	return id
}
