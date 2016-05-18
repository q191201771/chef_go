package email

import "testing"

func TestSend(t *testing.T) {
	err := Send(
		[]string{"cbmprl@foxmail.com"},
		"email_subject_test",
		`<strong style="color:red">hello<strong> world`,
	)
	if err != nil {
		t.Fatal(err)
	}
}
