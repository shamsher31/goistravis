// Package travis check if your code is running on Travis
package travis // import "github.com/shamsher31/goistravis"

import (
	"os"
	"strings"
)

func get(t string) bool {
	for _, val := range os.Environ() {
		s := strings.Split(val, "=")
		if strings.Contains(s[0], t) {
			return true
		}
	}
	return false
}

// Is returns whether env contains Travis config
func Is() bool {
	if get("TRAVIS") && get("CI") {
		return true
	}
	return false
}
