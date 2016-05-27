package env

import (
	"os"
	"path/filepath"
)

func SetBaseDirLocal() error {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	err = os.Chdir(path)
	return err
}
