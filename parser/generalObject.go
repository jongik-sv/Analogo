package parser

import (
	"bytes"
	"fmt"
	"goproject/AnaloGo/lexer/token"
	"strings"
	"time"
)

// ToDo : Merge 변경
// g.Properties.Merge(t) 구문을 parser.go의 Merge로 바꿀 것
// lex_pattern.ods 에서 Copy부분을 다시 한번 잘 확인할 것
type GeneralObject struct {
	ObjectType *string
	// ProcessType      string
	Key              string
	Properties       *Properties
	GlobalProperties *Properties
	Children         []*Node
	ChildrenMap      map[string]*Node
	// CurrentNode      *Node
}

// const (
// 	PARALELL   string = "Parallel"
// 	SEQUENTIAL string = "Sequential"
// )

/*
AST 구조 생성 컨셉
	1. g.Key로 자식이 존재 하나?
		Y : Parse
		N : 2번으로
	2. g.Key + 하위 프로세스 (입력키(hgch:XB0c:X6Px) --> g.Key(hgch; 현재위치) +  하위키(XB0c)
		Y : Parse
		N : NewObject
*/
func (g *GeneralObject) Parse(t *token.Token) int {
	var v *Node
	var ok bool = false
	var fn string = "Parse"
	var key string = t.Header.Key

	if v, ok = g.ChildrenMap[g.Key]; ok && v != nil {
		// g.Key의 가장 마지막 노드
		// 예) A(할) --> A(아) --> A(본)
		// g의 키가 A 인데 자식으로 A가 있을 경우 자식에게 토큰을 넘긴다.
		fn = "Child"

	} else {
		if g.Key == t.Header.Key {
			// 해당 키로 가장 마지막 후손 노드
			// A(할) --> A(아) --> A(본) 인데 A로 들어 왔을 경우 본인에서 파싱을 한다.
			// Sequential 일 경우 이렇게 됨
			fn = "Parse"
			// key = t.Header.Key
		} else {
			// 해당 키로 가장 마지막 후손 노드
			// A(할) --> A(아) --> A(본) 인데 A:B가 들어올 경우 A(본)에서 A:B의 자식이 있는지 확인한다.
			// A(할) --> A(아) --> A(본) 인데 A:B:C가 들어올 경우에도 C를 잘라내고 A(본)에서 A:B의 자식으로 파싱한다. 만약 없으면 에러!!!
			nextThreadKey := strings.Split(t.Header.Key[len(g.Key)+1:], ":")
			if len(nextThreadKey) == 0 {
				fmt.Println("잘못 들어왔어요")
				return ERR
			}

			key = g.Key + ":" + nextThreadKey[0]
			if v, ok = g.ChildrenMap[key]; ok && v != nil {
				fn = "Child"
			} else {
				fn = "Parse"
			}
		}
	}

	// v, ok := g.ChildrenMap[t.Header.Key]
	// if ok && v != nil {
	if fn == "Child" {
		// if g.CurrentNode != nil {
		// r := (*g.CurrentNode).Parse(t)
		r := (*v).Parse(t)
		if r == UGT {
			// SQL 노드는 끝나는 문장이 없으므로 SQL노드에서 SQL노드를 종료하고 UGT(UnGet ; 토큰을 소모 하지 않고 )을 리턴
			// 현재 노드에서 토큰을 다시 소모한다.
			delete(g.ChildrenMap, key)
			r = (*v).Parse(t)
		}
		switch r {
		case ACK:
			return ACK
		case RTN:
			delete(g.ChildrenMap, key)
			// delete(g.ChildrenMap, t.Header.Key)
			return ACK
		case ERR:
			return ERR
		}
	}

	objectType, object, se := objectDecode(t.TokenType)
	// SERVICE, TASK, PROCESS ... 등 오브젝트이면
	switch objectType {
	case OBJECT:
		if se == "End" {
			// g.Properties.Merge(t)
			Merge(t, g.Properties, g.GlobalProperties)
			if object != *g.ObjectType {
				println("XXXXXXXXXXXXXXX", t.TokenType, t.Header.Text)
				return ERR
			}
			return RTN
		}

		// 예: task start
		obj := NewGeneralObject(object, t.Header.Key, g.GlobalProperties)
		// obj.Properties.Merge(t)
		Merge(t, obj.Properties, obj.GlobalProperties)
		var node Node
		node = obj
		g.Children = append(g.Children, &node)
		g.ChildrenMap[t.Header.Key] = &node
	case QUERY:
		obj := NewQueryObject(object, t.Header.Key, g.GlobalProperties)
		// obj.Properties.Merge(t)
		Merge(t, obj.Properties, obj.GlobalProperties)
		var node Node
		node = obj
		g.Children = append(g.Children, &node)
		g.ChildrenMap[t.Header.Key] = &node
		g.Properties.Add("hasQuery", "Y", NORMAL)
	case PARAMETER: // 이건 오류다.
	default:
		switch t.TokenType {
		case "X":
			// 메시지인 경우 메시지를 등록한다.
			obj := NewMessageObject(t)
			var node Node
			node = obj
			g.Children = append(g.Children, &node)

		default:
			// 메시지가 아닌 경우 프로퍼티 내용을 넣는다.
			// g.Properties.Merge(t)
			Merge(t, g.Properties, g.GlobalProperties)
		}
	}
	return ACK
}

func (g *GeneralObject) ToString() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	if g.Properties.GetFirst("ObjectType") == "Request" {
		_id := g.GlobalProperties.GetFirst("_id")
		out.WriteString(fmt.Sprintf("\"_id\":\"%s\", \n", _id))
		now := time.Now().Format("2006-01-02 15:04:05.000")
		out.WriteString(fmt.Sprintf("\"insertTime\":\"%s\", \n", now))
		out.WriteString(fmt.Sprintf("\"expireAfterSeconds\":%d, \n", 10))
		// out.WriteString(fmt.Sprintf("\"ExpireAt\":ISODate(\"%s\"), \n", strings.Replace(now, " ", "T", 1)+"Z"))
	}
	propertiesString := g.toStringProperties()
	out.WriteString(propertiesString)
	childrenString := g.toStringChildren()
	if childrenString != "" {
		out.WriteString(",\n")
		out.WriteString(childrenString)
	}
	out.WriteString("}\n")
	return out.String()
}

func (g *GeneralObject) MakeId() {
	if g.Properties.GetFirst("ObjectType") == "Request" {
		// "_id" : "2022-01-07 15:18:06, 151$UBIUPMA2$0$M472020010",
		startTime := g.GlobalProperties.GetFirst("StartTime")
		server := g.GlobalProperties.GetFirst("Server")
		logKey := g.GlobalProperties.GetFirst("LogKey")
		service := g.GlobalProperties.GetJoinString("Services")
		_id := startTime + "$" + server + "$" + logKey + "$" + service
		g.GlobalProperties.Add("_id", _id, NORMAL)
	}
}

// func (g *GeneralObject) ExpireAt() {
// 	if g.Properties.GetFirst("ObjectType") == "Request" {
// 		startTime := g.GlobalProperties.GetFirst("StartTime")
// 		server := g.GlobalProperties.GetFirst("Server")
// 		logKey := g.GlobalProperties.GetFirst("LogKey")
// 		service := g.GlobalProperties.GetJoinString("Services")
// 		_id := startTime + "$" + server + "$" + logKey + "$" + service
// 		g.GlobalProperties.Add("_id", _id, NORMAL)
// 	}
// }

func (g *GeneralObject) toStringProperties() string {
	var out bytes.Buffer

	out.WriteString("\"property\":\n")
	propertiesString := ""
	if *g.ObjectType == "Request" {
		propertiesString = g.GlobalProperties.ToString()
	} else {
		propertiesString = g.Properties.ToString()
	}
	out.WriteString(propertiesString)
	return out.String()
}

func (g *GeneralObject) toStringChildren() string {
	if g.Children == nil || len(g.Children) == 0 {
		return ""
	}
	var out bytes.Buffer
	var children string
	var isFirst bool = true
	out.WriteString("\"child\":[")
	for _, v := range g.Children {
		if isFirst {
			children = fmt.Sprintf("\n%s", (*v).ToString())
			isFirst = false
		} else {
			children = fmt.Sprintf(",\n%s", (*v).ToString())
		}
		// fmt.Println(i, children)
		out.WriteString(children)
	}
	out.WriteString("]\n")
	return out.String()
}

func (g *GeneralObject) GetObjectType() string {
	return *g.ObjectType
}

func (g *GeneralObject) GetChildren() []*Node {
	return g.Children
}
