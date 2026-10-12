package lab

import "testing"

func TestNormalizeUser(t *testing.T) {
	if got := NormalizeUser("alice"); got != "alice" {
		t.Fatalf("NormalizeUser() = %q", got)
	}
}
