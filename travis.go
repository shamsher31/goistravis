// Package travis check if your code is running on Travis
package travis // import "github.com/shamsher31/goistravis"

import (
	"os"
	"strings"
)

func get(t string) bool {
	for _, val := range os.Environ() {
		s := strings.Split(val, "=")
		if s[0] == t {
			return true
		}
	}
	return false
}

func Is() bool {
	if get("TRAVIS") && get("CI") {
		return true
	}
	return false
}
