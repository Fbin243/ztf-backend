package errors

import "fmt"

var (
	ErrorNotFound       = fmt.Errorf("not found")
	ErrorNoRowsAffected = fmt.Errorf("no rows affected")
)
