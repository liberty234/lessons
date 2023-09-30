package files

import "fmt"

func GetFiles(file1, file2, file3 string) (string, string, string) {
	file1 = fmt.Sprintf("File1: %s", file1)
	file2 = fmt.Sprintf("File2: %s", file2)
	file3 = fmt.Sprintf("File3: %s", file3)
	return file1, file2, file3

}
