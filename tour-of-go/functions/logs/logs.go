package logs

import "fmt"

func GetLogsTime(logIn, logOut string) (string, string) {
	li := fmt.Sprintf("Log In Time:%s", logIn)
	lo := fmt.Sprintf("Log Out Time:%s", logOut)
	return li, lo
}
