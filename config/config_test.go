package config

import (
	"os"
	"testing"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestSave(t *testing.T) {
	defer func() {
		os.Remove("test.json")
	}()
	err := Save("test.json", user{
		Name: "chef",
		Age:  22,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoad(t *testing.T) {
	defer func() {
		os.Remove("test.json")
	}()
	err := Save("test.json", user{
		Name: "chef",
		Age:  22,
	})
	if err != nil {
		t.Fatal(err)
	}
	var u user
	err = Load("test.json", &u)
	if err != nil {
		t.Fatal(err)
	}
	if u.Name != "chef" || u.Age != 22 {
		t.Fatal(u)
	}
}
