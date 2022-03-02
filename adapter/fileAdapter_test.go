package adapter

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScanDirectory(t *testing.T) {
	ch := make(chan string, 100)

	fileAdapter := &FileAdapter{
		Server:          "운영계1",
		Chain:           "C10",
		Directory:       "C:/Temp",
		FilePattern:     "*.log",
		BackupDirectory: "C:/Temp/backup",
		Sender:          ch,
	}

	fileAdapter.ScanDirectory()

}

func TestPushFile(t *testing.T) {
	assert := assert.New(t)

	ch := make(chan string, 100)

	fileAdapter := &FileAdapter{
		Server:          "운영계1",
		Chain:           "C10",
		Directory:       "C:/Temp",
		FilePattern:     "*.log",
		BackupDirectory: "C:/Temp/backup",
		Sender:          ch,
	}

	f, _ := os.Create("./pushFileTest.txt")
	f.Sync()

	f.WriteString("push file test\n")
	f.Close()
	fileAdapter.pushFile("./pushFileTest.txt")

	time.Sleep(time.Second * 3)

	text := <-ch

	os.Remove(fileAdapter.BackupDirectory + "/pushFileTest.txt")
	assert.Equal(text, "push file test")
	fmt.Println("받은 내용 : ", text)

}

func TestMovefile(t *testing.T) {
	fileAdapter := &FileAdapter{}
	fileAdapter.moveFile("C:\\Temp\\test.log", "C:\\Temp\\backup\\")

}

func TestFileGlob(t *testing.T) {
	files, err := filepath.Glob(filepath.Join("C:/analogo/rcv", "/*"))

	fmt.Println(files, err)

}
