package travis

import "testing"

func TestIs(t *testing.T) {
	if Is() == false {
		t.Fatal("Should return true")
	}
}
