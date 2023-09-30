package logs

import "fmt"

func GetLogsTime(logIn, logOut string) (string, string) {
	logIn = fmt.Sprintf("Log In Time:%s", logIn)
	logOut = fmt.Sprintf("Log Out Time:%s", logOut)
	return logIn, logOut
}
