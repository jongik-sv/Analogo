package broker

import (
	"bufio"
	"fmt"
	"goproject/AnaloGo/mongoRepository"
	"goproject/AnaloGo/scanner"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAdapterEmul(t *testing.T) {

	adapterChan := make(chan string, 100)
	brokerChan := make(chan string, 100)

	fileName := "../lexer/LogSamples/pmf.log"
	// fileName := "../lexer/LogSamples/kNgo.log"

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

	go func() {
		scanner.Listen(adapterChan, brokerChan)
	}()

	go func() {
		broker := Broker{Server: "운영계", Chain: "oasis4"}
		broker.Repository = &mongoRepository.Repository{
			URI:            "mongodb://210.1.1.40:27017",
			DataBase:       "test",
			CollectionName: broker.Chain,
		}
		broker.Repository.Connect()
		broker.Listen(brokerChan)
	}()

	time.Sleep(time.Second * 10)

}

func TestListern(t *testing.T) {
	assert := assert.New(t)
	ch := make(chan string, 100)
	broker := &Broker{Server: "운영계1", Chain: "C10"}
	go broker.Listen(ch)
	ch <- "[2021-12-13 11:49:21,162][DEBUG][AAAA:BBBB:CCCC][org.springframework.web.servlet.DispatcherServlet]-Failed to complete request: java.lang.RuntimeException: my exception"

	// fmt.Println(len(broker.ChildChannels))

	_, ok := broker.ChildChannels["AAAA"]
	assert.Equal(ok, true)
	// assert.NotEqual(broker.ChildChannels, nil)

	broker.removeChild("AAAA")
	_, ok2 := broker.ChildChannels["AAAA"]
	assert.Equal(ok2, false)
	// assert.Equal(broker.ChildChannels, nil)

}
