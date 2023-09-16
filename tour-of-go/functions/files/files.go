package files

import "fmt"

func GetFiles(file1, file2, file3 string) (string, string, string) {
	eb := fmt.Sprintf("File1: %s", file1)
	wb := fmt.Sprintf("File2: %s", file2)
	mb := fmt.Sprintf("File3: %s", file3)
	return eb, wb, mb

}
