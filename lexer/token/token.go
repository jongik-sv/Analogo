package token

import (
	"fmt"
	"goproject/AnaloGo/lexer/logHeader"
)

type Token struct {
	TokenType  string
	Status     string
	Header     *logHeader.LogHeader
	Properties map[string]string
	CopyObject *map[string]string
}

const (
	NONE_KEY  = "NONE_KEY"
	OTHER_KEY = "OTHER_KEY"
)

func (t *Token) Print() {
	fmt.Println("{")
	fmt.Printf("\t\"TokenType\":\"%s\",\n", t.TokenType)
	fmt.Printf("\t\"Status\":\"%s\",\n", t.Status)
	t.Header.Print()

	if t.Properties != nil {
		fmt.Println("{")
		i := 0
		for k, v := range t.Properties {
			if i == 0 {
				fmt.Printf("\"%s\":\"%s\"\n", k, v)
			} else {
				fmt.Printf(",\"%s\":\"%s\"\n", k, v)
			}
		}
		fmt.Println("}")
	}
	fmt.Println("}")
}
