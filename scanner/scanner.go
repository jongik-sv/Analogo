package scanner

import (
	"fmt"
	"strings"
	"time"
)

// var logKeyRG *regexp.Regexp

const START_CHAR = "["

func Listen(adapterChan, brokerChan chan string) error {
	buff := make([]string, 0, 1000)
	if adapterChan == nil {
		return fmt.Errorf("subscription channels cannot be nil %d", 1000)
	}
	// t := time.Now()
	// tick := time.Tick(time.Second * 2)
	t := time.NewTimer(time.Second * 1)
	tickCounter := 0
	for {
		select {
		case logString := <-adapterChan:
			tickCounter = 0
			// 디버그 기능 , 종료 기능
			if logString == "<<<<LOG EOF>>>>" {
				sendLog(strings.Join(buff, "\n"), brokerChan)
				sendLog("<<<<LOG EOF>>>>", brokerChan)
			}
			logString = trimReturnCharacter(logString)
			st := isStart(logString)

			if st {
				// 시작하는 로그가 올 경우 기존에 버퍼에 로그가 남아 있다면 보낸다.
				// 로그가 시작과 끝이 잘 되어 있으면 이 로직은 필요가 없다.
				if len(buff) > 0 {
					sendLog(strings.Join(buff, "\n"), brokerChan)
					buff = buff[:0]
				}
			}

			buff = append(buff, logString)

		case <-t.C:
			tickCounter++
			if tickCounter > 3 {
				if len(buff) > 0 {
					fmt.Println("3초 틱 발생해서 전송")
					sendLog(strings.Join(buff, "\n"), brokerChan)
					buff = buff[:0]
					tickCounter = 0
				}
			}
		}

		if adapterChan == nil {
			return nil
		}
	}

}

// todo : 여기서 만든 lex가 한 프로세스 분석을 다 끝내면 채널에 자기자신의 로그키를 보내서
//        채널 삭제를 하는 로직을 개발해야 한다.

func sendLog(log string, brokerChan chan string) error {
	if brokerChan == nil {
		return fmt.Errorf("subscription channels cannot be nil %d", 1000)
	}

	brokerChan <- log
	return nil
}

func isStart(str string) bool {
	return strings.HasPrefix(str, START_CHAR)
}

func trimReturnCharacter(str string) string {
	return strings.TrimRight(str, "\n")
}
