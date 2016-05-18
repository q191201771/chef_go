package mod_files

import (
	"fmt"
	"testing"
)

func addHello(in string) string {
	return in + string("\nHello\n")
}

func TestModFiles(t *testing.T) {
	files, err := ModFiles("/private/tmp", true, ".go", addHello)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("affected files:", files)
}
