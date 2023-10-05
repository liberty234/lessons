package jsonop

import "fmt"

func GetPara(name1 string) string {
	GitHubName := fmt.Sprintf("GitHub Name:%s", name1)
	return GitHubName

}
