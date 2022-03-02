package lexer

import (
	"goproject/AnaloGo/lexer/logHeader"
	"goproject/AnaloGo/lexer/pattern"
	"goproject/AnaloGo/lexer/token"
	"goproject/AnaloGo/mongoRepository"
	"goproject/AnaloGo/parser"

	"time"
)

type Lexer struct {
	Server           string
	Chain            string
	Key              string
	ReciveChan       chan logHeader.LogHeader
	SendToParentChan chan string
	prevToken        *token.Token
	curToken         *token.Token
	ExpireTime       time.Duration
	expireTimer      *time.Timer
	autoSaveTimer    *time.Ticker
	autoSaveCnt      int
	isAutoSaved      bool
	Repository       *mongoRepository.Repository
}

func (l *Lexer) GetToken() *token.Token {

	for {
		select {
		case header := <-l.ReciveChan:
			// 로그를 수신했을 경우 자동저장 대기 시간을 0으로 변경, 자동저장 아닌 상태로 변경 --> 로그로 인해 AST의 수정이 있으므로
			l.autoSaveCnt = 0
			l.isAutoSaved = false
			return l.match(&header)
		case <-l.expireTimer.C:
			return l.getCustomeToken("TIMEOUT")
		case <-l.autoSaveTimer.C:
			l.autoSaveCnt++
			return l.getCustomeToken("AUTOSAVE")
		}
	}

}

func (l *Lexer) GetCurToken() *token.Token {
	return l.curToken
}

func (l *Lexer) match(header *logHeader.LogHeader) *token.Token {
	t, _ := pattern.Match(header)
	l.prevToken = l.curToken
	l.curToken = t
	return l.curToken
}

func (l *Lexer) getCustomeToken(tokenType string) *token.Token {
	// Server      string
	// Chain       string
	// Timestamp   string
	// Level       string
	// Key         string
	// Keys        []string
	// NDC         string // Nested Diagnostic Contexts
	// Text        string
	// TotalString string
	// OK          bool

	header := logHeader.LogHeader{
		Timestamp: l.curToken.Header.Timestamp,
		Key:       l.curToken.Header.Key,
	}
	t := &token.Token{TokenType: tokenType, Header: &header}
	l.prevToken = l.curToken
	l.curToken = t

	return l.curToken
}

func (l *Lexer) Run() {
	globalProperties := parser.NewProperties()

	obj := parser.NewGeneralObject("Request", l.Key, globalProperties)

	l.autoSaveCnt = 0
	l.isAutoSaved = false
	startTime := ""
	isFirst := true
	// 0이면 1000초까지 기다린다.
	if l.ExpireTime != 0 {
		l.expireTimer = time.NewTimer(time.Second * l.ExpireTime)
	} else {
		l.expireTimer = time.NewTimer(time.Second * 1000)
	}
	l.autoSaveTimer = time.NewTicker(time.Second * 2)

	defer func() {
		obj.MakeId()
		doc := (*obj).ToString()

		// 저장 코드 필요
		// fmt.Println("--------------------------------------------------------")
		// fmt.Println((*obj).ToString())
		// 채널이 닫힌 후 채널의 전체 내용을 제거한다.
		close(l.ReciveChan)
		l.autoSaveTimer.Stop()
		for range l.ReciveChan {
		}
		_id := (*obj).GlobalProperties.GetFirst("_id")
		l.Repository.Delete(_id)
		l.Repository.Insert(doc)

	}()
FOR:
	for {
		curToken := l.GetToken()
		if isFirst {
			isFirst = false
			startTime = curToken.Header.Timestamp
			obj.GlobalProperties.Add("StartTime", startTime, parser.NORMAL)
			obj.GlobalProperties.Add("Server", l.Server, parser.NORMAL)
			obj.GlobalProperties.Add("Chain", l.Chain, parser.NORMAL)
		}
		switch curToken.TokenType {
		case "EOF":
			obj.GlobalProperties.Add("finish", "NORMAL", parser.NORMAL)

			return
		case "TIMEOUT":
			obj.GlobalProperties.Add("finish", "TIMEOUT", parser.NORMAL)

			break FOR
		case "AUTOSAVE":
			if obj.GlobalProperties.GetFirst("finish") == "finish" {
				// lex의 Service End 에서 넣어준다.
				// Service가 여러개 있을 경우
				// start --> finish --> ... -> start --> finish
				// Service End가 되었으면 finish에 finish가 들아간다.

				// 여기에 EndTime 이 들어오도록 해야 한다.
				obj.GlobalProperties.Add("EndTime", l.curToken.Header.Timestamp, parser.NORMAL)
				// fmt.Println("startTime", startTime)
				// fmt.Println("endTime", l.curToken.Header.Timestamp)

				t := parser.DiffMiliseconds(startTime, l.curToken.Header.Timestamp)

				obj.GlobalProperties.Add("RunTime", t, parser.NORMAL)
				return
			}

			l.autoSaveCnt++
			if l.autoSaveCnt > 5 && !l.isAutoSaved {
				// 5초 이상 로그 수신이 없고 자동저장을 안했을 경우
				obj.GlobalProperties.Add("finish", "AUTOSAVE", parser.NORMAL)

				// 자동 저장 로직 구현
				l.isAutoSaved = true

			}

			// 저장
		default:
			obj.Parse(curToken)
		}

	}
}
