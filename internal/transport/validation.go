package transport

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validatorInstance *validator.Validate
	syncOnce          sync.Once
)

func GetValidator() *validator.Validate {
	syncOnce.Do(func() {
		validatorInstance = validator.New()
	})

	return validatorInstance
}
