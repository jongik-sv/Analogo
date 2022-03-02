package broker

import (
	"fmt"
	"goproject/AnaloGo/lexer/lexer"
	"goproject/AnaloGo/lexer/logHeader"
	"goproject/AnaloGo/mongoRepository"
)

type Broker struct {
	Server        string
	Chain         string
	ChildChannels map[string]chan logHeader.LogHeader
	Repository    *mongoRepository.Repository
}

func (b *Broker) Listen(scannerChan chan string) error {
	b.ChildChannels = make(map[string]chan logHeader.LogHeader)

	b.Repository = &mongoRepository.Repository{
		URI:            "mongodb://210.1.1.40:27017",
		DataBase:       "test",
		CollectionName: b.Chain,
	}
	b.Repository.Connect()
	//b.ChildChannels = childChannels

	// 어따가 쓸거냐면 lexer가 다 끝나면 이쪽으로 자기 번호를 보내주면
	// child 정리 할거다.
	reciveChanFromChild := make(chan string)
	for {
		if scannerChan == nil {
			return fmt.Errorf("subscription channels cannot be nil %d", 1000)
		}

		select {
		case logString := <-scannerChan:
			b.sendChild(logString, reciveChanFromChild)
		case removeLogKey := <-reciveChanFromChild:
			b.removeChild(removeLogKey)
		}
	}
}

func (b *Broker) sendChild(logString string, reciveChanFromChild chan string) {
	header := logHeader.GetLogHeader(logString)
	if header != nil && header.Key != "" {
		logKey := header.Keys[0]
		if _, ok := b.ChildChannels[logKey]; !ok {
			b.createChild(logKey)

		}
		b.ChildChannels[logKey] <- *header
	}
}

func (b *Broker) createChild(logKey string) {
	toLexCh := make(chan logHeader.LogHeader, 100)
	b.ChildChannels[logKey] = toLexCh
	lex := lexer.Lexer{
		Server:     b.Server,
		Chain:      b.Chain,
		Key:        logKey,
		ReciveChan: toLexCh,
		ExpireTime: 1000,
		Repository: b.Repository,
	}

	// var wg sync.WaitGroup
	// wg.Add(1)
	go func() {
		lex.Run()
		delete(b.ChildChannels, logKey)
		// wg.Done()
	}()
	// wg.Wait()
}

func (b *Broker) removeChild(removeLogKey string) {
	// fmt.Printf("%s 이거 지워주세요", removeLogKey)

	close(b.ChildChannels[removeLogKey])
	delete(b.ChildChannels, removeLogKey)
}
