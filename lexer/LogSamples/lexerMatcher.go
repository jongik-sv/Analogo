package main

import (
	"bufio"
	"fmt"
	"goproject/AnaloGo/lexer/logHeader"
	"goproject/AnaloGo/lexer/pattern"
	"goproject/AnaloGo/scanner"
	"os"
	"time"
)

func main() {

	adapterChan := make(chan string, 100)
	brokerChan := make(chan string, 100)

	fileName := "./pmf-debug.log"

	go func(fileName string) {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("error : ", "파일을 찾을 수 없습니다.", fileName)
			return
		}
		defer file.Close()
		fileScanner := bufio.NewScanner(file)

		for fileScanner.Scan() {
			line := fileScanner.Text()
			adapterChan <- line
		}
		close(adapterChan)
	}(fileName)

	go func(adapterChan, brokerChan chan string) {
		scanner.Listen(adapterChan, brokerChan)
	}(adapterChan, brokerChan)

	go func(brokerChan chan string) {
		for str := range brokerChan {
			h := logHeader.GetLogHeader(str)
			if !h.OK {
				continue
			}
			if h.Key != "" {
				t, _ := pattern.Match(h)
				// if t.TokenType == "X" {
				if t.TokenType[0] != '@' {
					fmt.Printf("%-23s  %s\n", t.TokenType, t.Header.TotalString)
				}

			}

		}
	}(brokerChan)
	time.Sleep(time.Second * 5)
}
