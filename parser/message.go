package parser

import (
	"goproject/AnaloGo/lexer/token"
)

type MessageObject struct {
	message string
}

func (m MessageObject) Parse(t *token.Token) int {
	return ACK
}

func (m MessageObject) ToString() string {
	return "\"" + m.message + "\""
}

func (m MessageObject) GetObjectType() string {
	return "Message"
}

func (m MessageObject) GetChildren() []*Node {
	return nil
}
