package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/secsy/goftp"
)

type ServerConfig struct {
	Servers map[string]Server `json:"servers"`
}

type Server struct {
	ServerIP    string   `json:"Server"`
	UserName    string   `json:"UserName"`
	Password    string   `json:"Password"`
	RemoteRoot  string   `json:"RemoteRoot"`
	RemoteDirs  []string `json:"RemoteDirs"`
	AnalogRoot  string   `json:"AnalogRoot"`
	FilePattern string   `json:"FilePattern"`
}

func (f *Server) ftpConfig() goftp.Config {

	config := goftp.Config{
		User:               f.UserName,
		Password:           f.Password,
		ConnectionsPerHost: 10,
		Timeout:            10 * time.Second,
		Logger:             os.Stderr,
	}
	return config
}

func Config(propertiesFilePath string) map[string]Server {
	dat, err := ioutil.ReadFile(propertiesFilePath)

	if err != nil {
		panic(err)
	}
	c := make(map[string]Server)

	err = json.Unmarshal(dat, &c)
	if err != nil {
		panic(err)
	}
	return c

}

func (f *Server) connect() (*goftp.Client, error) {

	client, err := goftp.DialConfig(f.ftpConfig(), f.ServerIP)
	if err != nil {
		panic(err)
	}

	return client, nil
}

func (f *Server) getFileFromDir(dirName string) ([]string, []string, error) {
	// for

	client, err := f.connect()

	if err != nil {
		// panic(err)
		return nil, nil, err
	}

	var successlist []string
	var errorlist []string

	tmpDir := f.AnalogRoot + "/temp"
	rcvDir := f.AnalogRoot + "/rcv"
	errDir := f.AnalogRoot + "/error"
	remoteDir := f.RemoteRoot + "/" + dirName
	loop := true

	if files, err := filepath.Glob(filepath.Join(tmpDir, "/*.log*")); err != nil {
		fmt.Println(err)
		return nil, nil, err
	} else {
		for _, file := range files {
			filename := strings.Replace(file, "\\", "/", -1)
			dstFileName := strings.Replace(rcvDir, "\\", "/", -1) + "/" + path.Base(filename)
			if err := os.Rename(file, dstFileName); err != nil {
				fmt.Println(err)
				return nil, nil, err
			}
		}
	}

	// downloadChan := make(chan string, 3)
	// downloadChan <- "UBMWQ"

	for loop {
		fileNames, err := f.getFileList(dirName, client)
		if err != nil {
			// panic(err)
			return successlist, errorlist, err
		}

		if len(fileNames) <= 0 {
			fmt.Println(".")
			time.Sleep(time.Second * 6)
			continue

		}
		for _, fileName := range fileNames {
			err := downloadFile(remoteDir, tmpDir, fileName, client)

			if err != nil {
				errString := fmt.Sprintf("%s: %v\n", fileName, err)
				errorlist = append(errorlist, errString)
			} else {
				srcPath := tmpDir + "/" + fileName
				dstPath := rcvDir + "/" + fileName
				err = os.Rename(srcPath, dstPath)

				if err != nil {
					fmt.Println("파일 이동 오류 ", dstPath)
					os.Remove(dstPath)
					if err = os.Rename(srcPath, dstPath); err != nil {
						os.Rename(srcPath, errDir)
					}
				}

				if err != nil {
					errString := fmt.Sprintf("%s: %v\n", fileName, err)
					errorlist = append(errorlist, errString)
				} else {

					successlist = append(successlist, fileName)
				}
			}
		}

		// loop = false
		fmt.Println("error list : ", errorlist)
		fmt.Println("success list : ", successlist)
		successlist = successlist[:0]
		errorlist = errorlist[:0]

		time.Sleep(time.Second * 6)
	}

	return successlist, errorlist, nil

}

func ensureDir(dirName string) error {
	// https://stackoverflow.com/questions/28448543/how-to-create-nested-directories-using-mkdir-in-golang
	err := os.MkdirAll(dirName, os.ModeDir)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func downloadFile(remoteDir, localDir, fileName string, client *goftp.Client) error {

	downloadFilePath := remoteDir + "/" + fileName //  strings.Join([]string{f.RemoteRoot, dir, fileName}, "/")
	tempFilePath := localDir + "/" + fileName      // strings.Join([]string{f.RemoteRoot, dir, fileName}, "/")

	fmt.Printf("downloaded fileName: %s\n", downloadFilePath)

	fp, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}

	err = client.Retrieve(downloadFilePath, fp)
	defer client.Delete(downloadFilePath)
	defer fp.Close()
	if err != nil {
		return err // fmt.Errorf("일단 오류")
	}

	return nil
}

func (f *Server) getFileList(dirName string, client *goftp.Client) ([]string, error) {
	lastChar := f.RemoteRoot[len(f.RemoteRoot)-1:]
	dirDelim := ""

	if lastChar != "/" && lastChar != "\\" {
		dirDelim = "/"
	}

	// fmt.Printf("디렉토리 ReadDir %s\n", f.RemoteDir+dirDelim+dirName+dirDelim+f.FilePattern)
	files, err := client.ReadDir(f.RemoteRoot + dirDelim + dirName + dirDelim + f.FilePattern)

	fileNames := []string{}

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		// fmt.Println(file.Name())
		if getFileExtension(file.Name()) != "log" {
			fileNames = append(fileNames, file.Name())
		}
	}
	sort.Slice(fileNames, func(i, j int) bool {
		return fileNames[i] < fileNames[j]
	})

	// fmt.Println(fileNames)
	return fileNames, nil
}

func getFileExtension(fileName string) string {
	lastindex := strings.LastIndex(fileName, ".")

	if lastindex == -1 || len(fileName) <= 1 {
		return ""

	}
	return fileName[strings.LastIndex(fileName, ".")+1:]
}

func localDirCheck(ftpClient Server) {

	if err := ensureDir(ftpClient.AnalogRoot); err != nil {
		panic(err)
	}
	if err := ensureDir(ftpClient.AnalogRoot + "/rcv"); err != nil {
		panic(err)
	}
	if err := ensureDir(ftpClient.AnalogRoot + "/temp"); err != nil {
		panic(err)
	}
	if err := ensureDir(ftpClient.AnalogRoot + "/error"); err != nil {
		panic(err)
	}
}

func main() {

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
	// serverWait.Wait()

}
