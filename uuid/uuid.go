package uuid

import (
	"log"
	"os/exec"
	"strings"
)

func UUID() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Panic(err)
	}
	return strings.TrimSuffix(strings.TrimSuffix(string(out), "\n"), "\r\n")
}
