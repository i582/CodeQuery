package runtime

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/i582/CodeQuery/pkg/utils"
)

func fatal(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

func BoolCast(expr interface{}) bool {
	switch exp := expr.(type) {
	case bool:
		return exp
	default:
		typ := reflect.TypeOf(exp)
		fatal("Type '%v' is not convertible to 'bool'", typ)
	}

	return false
}

func IntCast(expr interface{}) int64 {
	switch exp := expr.(type) {
	case int64:
		return exp
	default:
		typ := reflect.TypeOf(exp)
		fatal("Type '%v' is not convertible to 'int'", typ)
	}

	return 0
}

func Smaller(left, right interface{}) bool {
	if _, ok := left.(int64); !ok {
		fatal("Type '%v' is not convertible to 'int64'", reflect.TypeOf(left))
	}

	if _, ok := right.(int64); !ok {
		fatal("Type '%v' is not convertible to 'int64'", reflect.TypeOf(right))
	}

	return left.(int64) < right.(int64)
}

func Equal(left, right interface{}) bool {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		fatal("Type '%v' is not compatible with '%v'", reflect.TypeOf(left), reflect.TypeOf(right))
	}

	if _, ok := left.(bool); ok {
		return left.(bool) == right.(bool)
	}

	if _, ok := left.(int64); ok {
		return left.(int64) == right.(int64)
	}

	if _, ok := left.(string); ok {
		return left.(string) == utils.Unquote(right.(string))
	}

	return false
}

func StringContains(where string, what interface{}) bool {
	if _, ok := what.(string); !ok {
		fatal("The '%v' type is incompatible with 'string'", reflect.TypeOf(what))
	}

	str := utils.Unquote(what.(string))

	return strings.Contains(where, str)
}
