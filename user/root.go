package user

import "os/user"

func IsRoot() bool {
	user, err := user.Current()
	return err == nil && user.Gid == "0"
}
