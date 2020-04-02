package util

import (
	"fmt"
	"io"
	"os"
)

func FileIsExist(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func WriteToFile(fileName string, bytes []byte) (err error) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	n, _ := f.Seek(0, io.SeekEnd)
	_, err = f.WriteAt(bytes, n)
	fmt.Println("write succeed!")
	defer f.Close()

	return
}