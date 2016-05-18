package uuid

import (
	"strings"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := UUID()
	vec := strings.Split(uuid, "-")
	if len(vec) != 5 || len(vec[0]) != 8 || len(vec[1]) != 4 || len(vec[2]) != 4 || len(vec[3]) != 4 || len(vec[4]) != 12 {
		t.Fatal(vec)
	}
}
