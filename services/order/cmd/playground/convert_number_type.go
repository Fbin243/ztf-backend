package playground

import (
	"log"

	"github.com/shopspring/decimal"
)

func ConvertNumberType() {
	amount := int64(9_223_372_036_854_775_807) // Max int64
	value := 0.10
	result := int64(float64(amount) * value) // May give imprecise result

	log.Printf("==> ConvertNumberType")
	log.Printf("float64 amount: %f", float64(amount))
	log.Printf("float64 amount * value: %f", float64(amount)*value) // Failed at this line
	log.Printf("Converted amount: %v", result)
}

func ConvertNumberTypeWithDecimal() {
	amount := decimal.NewFromInt(9_223_372_036_854_775_807)
	value := decimal.NewFromFloat(0.10)
	result := amount.Mul(value)

	log.Printf("==> ConvertNumberTypeWithDecimal")
	log.Printf("Converted amount: %v", result.IntPart())
}

/**
 * ➜  ztf-backend git:(main) ✗ go run services/order/cmd/main.go convert-number-type
2025/07/04 17:38:39 Converted amount: 922337203685477632
*/
