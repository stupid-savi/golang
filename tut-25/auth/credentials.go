package auth

import (
	"github.com/fatih/color"
)

// Note:-  name of a variable, function or struct if starts with Capital Letter then it is Public else private
// private means we can't use it somewhere else

func LoginWithCredentials(username string, password string) {
	if len(username) > 0 && len(password) > 0 {
		color.Yellow("Successfuly Login")
	} else {
		color.Red("Login Failed")
	}
}
