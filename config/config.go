package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func SaveToLocalPath(filename string, obj interface{}) error {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filename = filepath.Join(path, filename)
	return Save(filename, obj)
}

func LoadFromLocalPath(filename string, obj interface{}) error {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filename = filepath.Join(path, filename)
	return Load(filename, obj)
}

func Save(filename string, obj interface{}) error {
	raw, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, raw, 0644)
	return err
}

func Load(filename string, obj interface{}) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &obj)
	return err
}
