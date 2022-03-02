package adapter

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type FileAdapter struct {
	Server          string
	Chain           string
	Directory       string
	FilePattern     string
	BackupDirectory string
	Sender          chan string
}

// var ftpClient *FtpClient
var moveFlag = true

func (adapter FileAdapter) Start() {
	for {
		// files, _ := adapter.ScanDirectory()
		// fmt.Println(files)
		// if ftpClient != nil {

		// }

		adapter.ScanDirectory()

	}
}

func (adapter FileAdapter) ScanDirectory() ([]string, error) {
	fmt.Println("파일 검색중 ", adapter.Directory, "/", adapter.FilePattern)
	files, err := filepath.Glob(filepath.Join(adapter.Directory, "/", adapter.FilePattern))

	if err != nil {
		fmt.Println("error : ", err)
		return nil, err
	}

	sort.Strings(files)

	for _, v := range files {
		fmt.Println("filename : ", v)
		adapter.pushFile(v)
	}
	time.Sleep(time.Second * 3)
	return files, nil
}

func (adapter FileAdapter) pushFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error : ", "파일을 찾을 수 없습니다.", fileName)
		return
	}
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		adapter.Sender <- line
		// fmt.Println(line)
	}
	adapter.Sender <- "EOF"
	file.Close()
	if moveFlag {
		adapter.moveFile(fileName, adapter.BackupDirectory)
	}

}

// todo : 파일 처리 후 날짜별 백업 위치에 이동하는 기능이 필요
// 현재는 대충 옮김
func (adapter FileAdapter) moveFile(fileName, dstPath string) error {
	fileName1 := strings.Replace(fileName, "\\", "/", -1)
	destFileName := strings.Replace(dstPath, "\\", "/", -1) + "/" + path.Base(fileName1)

	err := os.Rename(fileName1, destFileName)
	if err != nil {
		return err
	}
	return nil
}
