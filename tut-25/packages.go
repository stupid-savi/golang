// go mod init github.com/stupid-savi/golang
// go get github.com/fatih/color
// go mod tidy

package main

import (
	"github.com/fatih/color"

	"github.com/stupid-savi/golang/auth"
	"github.com/stupid-savi/golang/user"
)

func main() {
	user := user.User{
		Username: "go@gmil.com",
		Password: "golang",
	}
	auth.LoginWithCredentials(user.Username, user.Password)
	session := auth.GetUserSession()
	color.Green(session)

}
