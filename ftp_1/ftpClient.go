package main

import (
	"fmt"
	"log"

	"time"

	"github.com/jlaffaye/ftp"
)

type FtpClient struct {
	Server          string
	UserName        string
	Password        string
	RemoteDir       string
	LocalDir        string
	FilePattern     string
	BackupDirectory string
}

func main() {
	f := &FtpClient{
		Server:          "210.1.1.138:21",
		UserName:        "tstwas",
		Password:        "tstwas!!",
		RemoteDir:       "/APP/WAS/UBMWQ/LOG/log4j/SND",
		LocalDir:        "C:/analogo/rcv",
		FilePattern:     "*.log*",
		BackupDirectory: "C:/analogo/backup",
	}

	c, err := f.connect()
	if err != nil {
		log.Fatal(err)
	}
	dirs, err := f.dirList(c)
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range dirs {
		fmt.Println(dir.Name)
	}

}

func (f *FtpClient) connect() (*ftp.ServerConn, error) {
	c, err := ftp.Dial(f.Server, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = c.Login(f.UserName, f.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Do something with the FTP conn

	// if err := c.Quit(); err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	return c, nil
}

func (f *FtpClient) dirList(c *ftp.ServerConn) ([]*ftp.Entry, error) {

	entries, err := c.List(f.RemoteDir)
	if err != nil {
		log.Fatal(err)
	}

	dirs := []*ftp.Entry{}

	fmt.Println(dirs)
	for _, e := range entries {
		if e.Type == ftp.EntryTypeFolder {
			dirs = append(dirs, e)
		}
	}

	if err := c.Quit(); err != nil {
		log.Println(err)
		return nil, err
	}

	return dirs, nil
}

/*
func (f *FtpClient) connect1() (*goftp.Client, error) {

	client, err := goftp.DialConfig(f.Config(), f.Server)
	if err != nil {
		panic(err)
	}

	return client, nil
}

func (f *FtpClient) getDirList() ([]fs.FileInfo, error) {

	client, err := f.connect()
	if err != nil {
		panic(err)
	}

	files, err := client.ReadDir(f.RemoteDir)

	if err != nil {
		return nil, err
	}

	dirs := []fs.FileInfo{}

	for _, file := range files {
		if file.Name() == "SND" {
			continue
		}
		if !file.IsDir() {
			continue
		}

		dirs = append(dirs, file)

	}

	client.Close()
	return dirs, nil
}

func (f *FtpClient) getFileFromDir(dirName string) error {
	// for
	client, err := f.connect()
	if err != nil {
		// panic(err)
		return err
	}

	loop := true
	for loop {
		fileNames, err := f.getFileList(dirName, client)
		if err != nil {
			// panic(err)
			return err
		}

		for _, fileName := range fileNames {
			f.downloadFile(fileName, client)
		}
	}

	return nil

}

func (f *FtpClient) downloadFile(fileName string, client *goftp.Client) {
	fmt.Printf("downloaded fileName: %s\n", fileName)
	fp, err := os.Create(fileName)

	if err != nil {

	}

	err = client.Retrieve("aaa", fp)

}

func (f *FtpClient) getFileList(dirName string, client *goftp.Client) ([]string, error) {

	lastChar := f.RemoteDir[len(f.RemoteDir)-1:]
	dirDelim := ""

	if lastChar != "/" && lastChar != "\\" {
		dirDelim = "/"
	}

	// fmt.Printf("디렉토리 ReadDir %s\n", f.RemoteDir+dirDelim+dirName+dirDelim+f.FilePattern)
	files, err := client.ReadDir(f.RemoteDir + dirDelim + dirName + dirDelim + f.FilePattern)

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

	fmt.Println(fileNames)
	return fileNames, nil
}

func getFileExtension(fileName string) string {
	lastindex := strings.LastIndex(fileName, ".")

	if lastindex == -1 || len(fileName) <= 1 {
		return ""

	}
	return fileName[strings.LastIndex(fileName, ".")+1:]
}

func FtpRun() error {
	ftpClient := &FtpClient{
		Server:          "210.1.1.138",
		UserName:        "tstwas",
		Password:        "tstwas!!",
		RemoteDir:       "/APP/WAS/UBMWQ/LOG/log4j",
		LocalDir:        "C:/analogo/rcv",
		FilePattern:     "*",
		BackupDirectory: "C:/analogo/backup",
	}

	dirs, err := ftpClient.getDirList()

	if err != nil {
		return err
	}

	var wait sync.WaitGroup
	done := make(chan bool)
	wait.Add(len(dirs))
	for _, dir := range dirs {
		go func() {
			defer wait.Done()
			for {

				files, err := ftpClient.getFileList(dir.Name())

				if err != nil {
					fmt.Println(err)
					done <- true
				}

				for _, f := range files {
					fmt.Println(f)
				}
			}
		}()

	}

	wait.Wait()

	return nil
}
*/
