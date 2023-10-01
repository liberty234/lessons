package files

import (
	"fmt"
)

func GetFiles(file1, file2, file3 string) (file4 string, file5 string, file6 string) {
	file4 = fmt.Sprintf("File1: %s", file1)
	file5 = fmt.Sprintf("File2: %s", file2)
	file6 = fmt.Sprintf("File3: %s", file3)
	return

}
