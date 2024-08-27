package persistence

import (
	"fmt"
	"os"
	"server-v5-refactor-server/static/output"
)

func GetFileConnector(path string) (*os.File, error) {
	var fileConnector, err = os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(output.Error() + "Open file error: " + err.Error())
	} else {
		fmt.Println(output.Info() + "Open file success: " + path)
	}
	return fileConnector, err
}

func CloseFileConnector(fileConnector *os.File) {
	err := fileConnector.Close()
	if err != nil {
		fmt.Println(output.Error() + "Close file error: " + err.Error())
	}
}

func WriteToFile(fileConnector *os.File, str string) {
	_, err := fileConnector.WriteString(str)
	if err != nil {
		fmt.Println(output.Error() + "Write to file error: " + err.Error())
	}
}
