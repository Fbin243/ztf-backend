package convert

import (
	"log/slog"
	"strconv"
)

func MustConvStrToInt(idStr string) int64 {
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		slog.Error("Error convert string to int64", "string", idStr)
		return 0
	}

	return id
}

func ConvIntToStr(id int64) string {
	return strconv.FormatInt(id, 10)
}
