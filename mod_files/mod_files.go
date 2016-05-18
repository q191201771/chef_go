package mod_files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/// TODO
/// dir support `~/`
/// link

type ModFunc func(in string) (out string)

func ModFiles(dir string, recursive bool, fileSuffix string, modFunc ModFunc) (affectedFiles []string, err error) {
	if len(dir) <= 0 {
		err = fmt.Errorf("dir [%s] expected len > 0", dir)
		return
	}
	dir = strings.TrimSuffix(dir, "/")
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if (!recursive && dir != filepath.Dir(path)) ||
			info.IsDir() ||
			(len(fileSuffix) > 0 && !strings.HasSuffix(path, fileSuffix)) {
			return nil
		}
		in, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		out := modFunc(string(in))
		err = ioutil.WriteFile(path, []byte(out), info.Mode())
		affectedFiles = append(affectedFiles, path)
		return err
	})
	return
}
