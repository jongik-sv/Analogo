package parser

import (
	"fmt"
	"goproject/AnaloGo/lexer/token"
	"strconv"
	"strings"
	"time"
)

const (
	OBJECT int = iota
	MESSAGE
	QUERY
	PARAMETER
)

const (
	ACK int = iota
	RTN
	UGT // unget
	ERR
)

type Node interface {
	Parse(t *token.Token) int
	// Run() bool
	ToString() string
	GetObjectType() string
	GetChildren() []*Node
}

// func Merge() *Properties {
// 	for k, v := range b {
// 		a.Add(k, v)
// 	}
// 	return a
// }

func NewGeneralObject(object string, key string, globalProperties *Properties) *GeneralObject {
	obj := &GeneralObject{
		ObjectType: &object,
		// ProcessType:      "Normal",
		Key:              key,
		Properties:       NewProperties(),
		GlobalProperties: globalProperties,
		Children:         []*Node{},
		ChildrenMap:      make(map[string]*Node),
	}
	if object == "Request" {
		obj.GlobalProperties.Add("ObjectType", object, NORMAL)
	}
	obj.Properties.Add("ObjectType", object, NORMAL)
	return obj
}

func NewQueryObject(object, key string, globalProperties *Properties) *QueryObject {
	obj := &QueryObject{
		ObjectType:       &object,
		Key:              key,
		Properties:       NewProperties(),
		GlobalProperties: globalProperties,
		// Parameters:       []string{},
	}
	obj.Properties.Add("ObjectType", object, NORMAL)
	return obj
}

func NewMessageObject(token *token.Token) *MessageObject {
	var message MessageObject = MessageObject{message: token.Header.Text}
	return &message
}

func objectDecode(tokenType string) (int, string, string) {
	switch tokenType {
	case "Query":
		return QUERY, tokenType, tokenType
	case "Parameter":
		return PARAMETER, tokenType, tokenType
	default:
		if strings.Contains(tokenType, "_Start") || strings.Contains(tokenType, "_End") {
			index := strings.LastIndex(tokenType, "_")
			return OBJECT, tokenType[:index], tokenType[index+1:]
		}
		return MESSAGE, tokenType, ""
	}
}

func Merge(t *token.Token, p, g *Properties) {

	// tokenProperty := &t.Properties

	// 예시 : "ServiceName":"Name,Global.Services.Set"
	if t.CopyObject == nil {
		return
	}
	for source, targets := range *t.CopyObject {

		value := ""
		sourceSelector, sourceMapKey := getSourceKey(source)
		switch sourceSelector {
		case "Header":
			value = t.Header.GetData(sourceMapKey)
		case "default":
			value = t.Properties[sourceMapKey]
		default:
			value = "error[unknwon selector] 선택자를 잘못 지정했습니다."
		}

		for _, target := range strings.Split(targets, ",") {

			targetSelector, targetMapKey, targetDataType := getTargetKey(strings.TrimSpace(target))
			switch targetSelector {
			case "Global":
				g.Add(targetMapKey, value, targetDataType)
			case "Local":
				p.Add(targetMapKey, value, targetDataType)
			}

			// fmt.Println(sourceSelector, sourceMapKey, target)
		}
	}
}

func getSourceKey(source string) (string, string) {
	sourceSlice := strings.Split(source, ".")

	sourceSelector := ""
	sourceMapKey := ""
	if len(sourceSlice) > 1 {
		sourceSelector = sourceSlice[0]
		sourceMapKey = sourceSlice[1]
	} else {
		sourceSelector = "default"
		sourceMapKey = sourceSlice[0]
	}

	return sourceSelector, sourceMapKey

}

func getTargetKey(target string) (string, string, int) {
	targetSlice := strings.Split(target, ".")
	targetSelector := ""
	targetMapKey := ""
	targetDataType := NORMAL
	if len(targetSlice) >= 2 && (targetSlice[0] == "Global" || targetSlice[0] == "Local") {
		targetSelector = targetSlice[0]
		targetSlice = targetSlice[1:]
	} else {
		targetSelector = "Local"
	}

	targetMapKey = targetSlice[0]

	if len(targetSlice) >= 2 {
		switch strings.ToUpper(targetSlice[1]) {
		case "SLICE":
			fallthrough
		case "ARR":
			fallthrough
		case "ARRAY":
			targetDataType = ARRAY
		case "SET":
			fallthrough
		case "MAP":
			targetDataType = SET
		default:
			targetDataType = NORMAL
		}

	}

	return targetSelector, targetMapKey, targetDataType
}

func convertTime(s string) (time.Time, error) {
	strSlice := strings.Split(s, " ")
	strDay := strSlice[0]
	timeSlice := strings.Split(strSlice[1], ",")
	strTime := timeSlice[0]
	strMTime := timeSlice[1]
	timeString := fmt.Sprintf("%sT%s.%s+09:00", strDay, strTime, strMTime)
	return time.Parse(time.RFC3339, timeString)
}

func DiffMiliseconds(str, end string) string {
	st, err := convertTime(str)
	if err != nil {
		return ""
	}
	et, err := convertTime(end)
	if err != nil {
		return ""
	}
	return strconv.Itoa(int(et.Sub(st) / time.Millisecond))
}
