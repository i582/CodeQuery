package conf

import (
	"github.com/i582/CodeQuery/pkg/querylang/errors"
)

type Config struct {
	ErrorHandlerFunc func(e *errors.Error)
}
