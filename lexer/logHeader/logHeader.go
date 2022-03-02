package logHeader

import (
	"fmt"
	"strings"
)

type LogHeader struct {
	Server      string
	Chain       string
	Timestamp   string
	Level       string
	Key         string
	Keys        []string
	NDC         string // Nested Diagnostic Contexts
	Text        string
	TotalString string
	OK          bool
}

// type Property struct {
// 	Key   string
// 	Value string
// }

const (
	TIMESTAMP = iota
	LEVEL
	KEY
	NDC
)

func GetLogHeader(log string) *LogHeader {
	position := 0

	if log[0:1] != "[" {
		// fmt.Println("@@@ error : ", log)
		return makeErrorLogHeader(log)
	}

	for i, v := range log {
		if i > 20 && v == '-' {
			position = i
			break
		}
	}

	if position == 0 {
		return makeErrorLogHeader(log)
	}

	// if strings.Contains(log, "\"") {
	// }
	log = strings.ReplaceAll(strings.ReplaceAll(log, "\"", "'"), "\n", "\\n")

	headerTexts := strings.Split(log[1:position-1], "][")

	if len(headerTexts) != 4 {
		return makeErrorLogHeader(log)
	}
	logHeader := &LogHeader{Key: ""}
	logHeader.Text = log[position+1:]
	logHeader.Timestamp = headerTexts[TIMESTAMP]
	logHeader.Level = headerTexts[LEVEL]
	logHeader.Key = headerTexts[KEY]
	logHeader.Keys = strings.Split(headerTexts[KEY], ":") // 자식 프로세스 처리
	logHeader.NDC = headerTexts[NDC]
	logHeader.TotalString = log
	logHeader.OK = true

	return logHeader
}

func makeErrorLogHeader(log string) *LogHeader {

	logHeader := &LogHeader{Key: ""}
	logHeader.Text = ""
	logHeader.Timestamp = ""
	logHeader.Level = ""
	logHeader.Key = ""
	logHeader.NDC = ""
	logHeader.TotalString = log
	logHeader.OK = false
	return logHeader
}

func (l *LogHeader) Print() {
	fmt.Println("{")
	fmt.Printf("\t\"Server\":\"%s\",\n", l.Server)
	fmt.Printf("\t\"Chain\":\"%s\",\n", l.Chain)
	fmt.Printf("\t\"Timestamp\":\"%s\",\n", l.Timestamp)
	fmt.Printf("\t\"Level\":\"%s\",\n", l.Level)
	fmt.Printf("\t\"Key\":\"%s\",\n", l.Key)
	fmt.Print("\t\"Keys\":[")
	for i, v := range l.Keys {
		if i == 0 {
			fmt.Printf("\"%s\"", v)
		} else {
			fmt.Printf(", \"%s\"", v)
		}

	}
	fmt.Printf("],\n")
	fmt.Printf("\t\"NDC\":\"%s\",\n", l.NDC)
	fmt.Printf("\t\"Text\":\"%s\",\n", l.Text)
	fmt.Printf("\t\"TotalString\":\"%s\",\n", l.TotalString)
	fmt.Printf("\t\"OK\":%v\n", l.OK)
	fmt.Println("}")
}

func (l *LogHeader) GetData(k string) string {

	switch k {
	case "Server":
		return l.Server
	case "Chain":
		return l.Chain
	case "Timestamp":
		return l.Timestamp
	case "Level":
		return l.Level
	case "NDC":
		return l.NDC
	case "Key":
		return l.Key
	case "Text":
		return l.Text
	case "TotalString":
		return l.TotalString

	}

	return ""
}
