package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
IP             => '210.1.1.136',
ID             => 'uswas',
PASSWD         => 'US_meap9',
REMOTE_DIR     => '/APP/WAS/UBMWP/LOG/log4j/SND/',
LOCAL_DIR      => '/c/LogApp/LogSource/RCV',
FILE_PATTERN   => '*'
*/

// 'UBIUPMAQ' =>
// {
// 		 IP             => '210.1.1.138',
// 		 ID             => 'tstwas',
// 		 PASSWD         => 'tstwas!!',
// 		 REMOTE_DIR     => '/APP/WAS/UBMWQ/LOG/log4j/SND/',
// 		 LOCAL_DIR      => '/c/LogApp/LogSource/RCV',
// 		 FILE_PATTERN   => '*'
// }

func TestConfig(t *testing.T) {
	dat, _ := ioutil.ReadFile("./config.json")

	var c ServerConfig
	err := json.Unmarshal(dat, &c)
	if err != nil {
		panic(err)
	}

	// assert.NotEqual(t, ftpClient, nil)

}

func TestConnect(t *testing.T) {
	assert := assert.New(t)
	ftpClient := &Server{
		ServerIP:    "210.1.1.138",
		UserName:    "tstwas",
		Password:    "tstwas!!",
		RemoteRoot:  "/APP/WAS/UBMWQ",
		AnalogRoot:  "C:/analogo",
		FilePattern: "*",
	}

	client, _ := ftpClient.connect()

	// assert.NotEmpty(t, files)
	assert.NotEqual(t, client, nil)
}

func TestGetFileList(t *testing.T) {
	// assert := assert.New(t)
	ftpClient := &Server{
		ServerIP:   "210.1.1.138",
		UserName:   "tstwas",
		Password:   "tstwas!!",
		RemoteRoot: "/APP/WAS/UBMWQ/LOG/log4j/",
		AnalogRoot: "C:/analogo",
		RemoteDirs: []string{
			"C10",
		},
		FilePattern: "*",
	}
	client, err := ftpClient.connect()
	if err != nil {
		panic(err)

	}

	files, _ := ftpClient.getFileList("M60", client)

	t.Logf("%d\n", len(files))

	assert.NotEqual(t, len(files), 0)
}

func TestDownloadFile(t *testing.T) {
	// assert := assert.New(t)
	ftpClient := &Server{
		ServerIP:   "210.1.1.138",
		UserName:   "tstwas",
		Password:   "tstwas!!",
		RemoteRoot: "/APP/WAS/UBMWQ/LOG/log4j",
		RemoteDirs: []string{
			"M47",
		},
		AnalogRoot:  "C:/analogo",
		FilePattern: "",
	}
	client, err := ftpClient.connect()
	if err != nil {
		panic(err)

	}
	defer client.Close()

	// /APP/WAS/UBMWQ/LOG/log4j/M47/m47.log
	// "C:/analogo/m47.log
	err = downloadFile(ftpClient.RemoteRoot+"/M47", ftpClient.AnalogRoot, "m47.log", client)

	if err != nil {
		panic(err)
	}

}
func TestGetFileExtension(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(getFileExtension("M60"), "")
	assert.Equal(getFileExtension("."), "")
	assert.Equal(getFileExtension(""), "")
	assert.Equal(getFileExtension("M60.log"), "log")
	assert.Equal(getFileExtension("M60.teststring.log"), "log")
	assert.Equal(getFileExtension("m60.log2021-11-22-15-00b"), "log2021-11-22-15-00b")

}

// func TestMoveFile(t *testing.T) {

// 	fmt.Println(os.Rename("c:/ago/temp/c10.log2021-11-23-14-37b", "c:/ago/rcv/c10.log2021-11-23-14-37b"))

// }

func TestReadFromConfigFile(t *testing.T) {

	config := Config("./config.json")

	fmt.Printf("%v\n", config)

	s1 := config["UBMWQ"]
	assert := assert.New(t)

	fmt.Printf("ServerIP : %v\n", s1.ServerIP)
	fmt.Printf("UserName : %v\n", s1.UserName)
	fmt.Printf("Password : %v\n", s1.Password)
	fmt.Printf("RemoteDirRoot : %v\n", s1.RemoteRoot)
	fmt.Printf("RemoteDirs : %v\n", s1.RemoteDirs)
	fmt.Printf("AnalogRoot : %v\n", s1.AnalogRoot)
	fmt.Printf("FilePattern : %v\n", s1.FilePattern)

	assert.Equal(s1.ServerIP, "210.1.1.138")
}

func TestFtpRun(t *testing.T) {

	flag.Set("test.timeout", "30m0s")
	config := Config("./config.json")

	ftpClient := config["UBMWQ"]
	localDirCheck(ftpClient)
	dirs := ftpClient.RemoteDirs
	var wait sync.WaitGroup
	// done := make(chan bool)
	wait.Add(len(dirs))
	for _, dir := range dirs {
		go func(dir string) {
			defer wait.Done()
			ftpClient.getFileFromDir(dir)
		}(dir)
	}

	wait.Wait()
}

func TestFtpRun2(t *testing.T) {

	config := Config("./config.json")

	// ftpClient := config["UBMWQ"]
	// var serverWait sync.WaitGroup
	// serverWait.Add(len(config))

	serverChan := make(chan string, 10)

	for key := range config {
		serverChan <- key
	}

	for serverName := range serverChan {
		go func(serverName string) {
			ftpClient := config[serverName]
			localDirCheck(ftpClient)
			dirs := ftpClient.RemoteDirs

			// done := make(chan bool)

			var wait sync.WaitGroup
			wait.Add(1)
			// 서버내에서 디렉토리 전송시 하나만 죽어도 그 서버는 다시 시작해라
			for _, dir := range dirs {

				go func(dir string) {
					defer wait.Done()
					ftpClient.getFileFromDir(dir)
				}(dir)

			}
			wait.Wait()
			serverChan <- serverName

		}(serverName)
		time.Sleep(time.Second * 3)
	}
}
