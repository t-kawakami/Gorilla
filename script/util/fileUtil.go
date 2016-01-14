package util
import (
	"os"
	"bufio"
	"io/ioutil"
)

func WriteFileAppend(text string, fileName string) {
	writeFile, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, os.ModePerm)

	writer := bufio.NewWriter(writeFile)
	writer.WriteString(text)
	writer.WriteString("\r\n")
	writer.Flush()
}

func WriteFile(text string, fileName string, flag int) {
	if flag == os.O_APPEND {
		WriteFileAppend(text, fileName)
	} else {
		ioutil.WriteFile(fileName, []byte(text), os.ModePerm)
	}
}

func DeleteFile(fileName string) {
	os.Truncate(fileName, 0)
}

func ReadFile(fileName string) string {
	contents, _ := ioutil.ReadFile(fileName)
	return string(contents)
}