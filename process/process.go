package process

import (
	"goproject/AnaloGo/adapter"
	"goproject/AnaloGo/broker"
	"goproject/AnaloGo/scanner"
	"strings"
	"sync"
)

type AdapterInterface interface {
	Start()
}

func Run(server, chain, mediaType string, config map[string]string) {

	var wait sync.WaitGroup
	wait.Add(1) // 하나라도 끝나면 종료 할거야

	chanFromAdapterToScanner := make(chan string, 1000)
	go func() {
		var adaterInterface AdapterInterface
		if mediaType == "file" {
			adaterInterface = runFileAdater(server, chain, chanFromAdapterToScanner)
		} else if mediaType == "kafka" {
			// TODO: kafka 인 경우 runKafakaAdapter 로 바꾸어야 됨
			adaterInterface = runFileAdater(server, chain, chanFromAdapterToScanner)
		}
		adaterInterface.Start()
		wait.Done()
	}()

	chanFromScannerToBroker := make(chan string, 1000)
	go func() {
		scanner.Listen(chanFromAdapterToScanner, chanFromScannerToBroker)
		wait.Done()
	}()

	go func() {
		broker := &broker.Broker{Server: server, Chain: chain}
		broker.Listen(chanFromScannerToBroker)
		wait.Done()
	}()

	wait.Wait()

}

func runFileAdater(server, chain string, chanFromAdapterToScanner chan string) AdapterInterface {

	fileAdapter := &adapter.FileAdapter{
		Server:          server,
		Chain:           chain,
		Directory:       "C:/temp",
		FilePattern:     strings.ToLower(chain) + ".log*",
		BackupDirectory: "C:/temp/backup/",
		Sender:          chanFromAdapterToScanner,
	}
	return fileAdapter

}
