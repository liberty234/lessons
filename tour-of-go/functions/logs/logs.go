package logs

import (
	"fmt"
)

func GetLogsTime(logIn, logOut string) (log1, log2 string) {
	log1 = fmt.Sprintf("Log In Time:%s", logIn)
	log2 = fmt.Sprintf("Log Out Time:%s", logOut)
	return
}
