package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GenIndent(level int) string {
	var res string
	for i := 0; i < level; i++ {
		res += "   "
	}
	return res
}

func DefaultCacheDir() string {
	defaultCacheDir, err := os.UserCacheDir()
	if err != nil {
		defaultCacheDir = ""
	} else {
		defaultCacheDir = filepath.Join(defaultCacheDir, "CodeQuery")
	}
	return defaultCacheDir
}

func NormalizeSlashes(str string) string {
	return strings.ReplaceAll(str, `\`, `\\`)
}

var NameToIdentifierRegexp = regexp.MustCompile("[^a-zA-Z0-9]")

func NameToIdentifier(str string) string {
	return NameToIdentifierRegexp.ReplaceAllString(str, "_")
}

// Unquote returns unquoted version of s, if there are any quotes.
func Unquote(s string) string {
	if len(s) >= 2 && (s[0] == '\'' || s[0] == '"') {
		return s[1 : len(s)-1]
	}
	return s
}
