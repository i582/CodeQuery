package namegen

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/VKCOM/noverify/src/meta"
)

func FileFunction(filename string) string {
	hash := md5.Sum([]byte(filename))
	return "src$" + hex.EncodeToString(hash[:]) + "$" + filename
}

func DefaultConstructor(class string) string {
	return Method(class, "__construct (default autogenerated)")
}

func Method(class, method string) string {
	return class + "::" + method
}

func ClassFQN(state *meta.ClassParseState, class string) string {
	return state.Namespace + `\` + class
}

func FunctionFQN(state *meta.ClassParseState, name string) string {
	if state.Namespace != "" {
		return state.Namespace + `\` + name
	}

	return `\` + name
}

func FromStubs(name string) bool {
	return strings.Contains(name, "phpstorm-stubs")
}