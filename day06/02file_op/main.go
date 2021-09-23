package main

import (
	"io"
	"os"
)

func f2() {
	fileObj, err := os.OpenFile(`C:\Users\Aone\go\src\gitlab.renrenchongdian.com\studygo\day06\02file_op\sb.txt`, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer fileObj.Close()

	tmpFile, err := os.OpenFile(`C:\Users\Aone\go\src\gitlab.renrenchongdian.com\studygo\day06\02file_op\sb.tmp`, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer tmpFile.Close()

	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		return
	}
	tmpFile.Write(ret[:n])

	s := []byte{'C'}
	tmpFile.Write(s)

	x := [1024]byte{}
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		tmpFile.Write(x[:n])
	}

	fileObj.Close()
	tmpFile.Close()
	os.Rename(
		`C:\Users\Aone\go\src\gitlab.renrenchongdian.com\studygo\day06\02file_op\sb.tmp`,
		`C:\Users\Aone\go\src\gitlab.renrenchongdian.com\studygo\day06\02file_op\sb.txt`,
	)

	// fileObj.Seek(1, 0)
	// s := []byte{'C'}
	// fileObj.Write(s)
	// fmt.Println(string(ret[:n]))

}

func main() {
	f2()
}
