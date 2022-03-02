package parser

import (
	"bytes"
	"goproject/AnaloGo/lexer/token"
)

type QueryObject struct {
	ObjectType       *string
	Key              string
	Properties       *Properties
	GlobalProperties *Properties
}

func (s *QueryObject) Parse(t *token.Token) int {

	if t.TokenType != "Parameter" {
		// 쿼리는 종료 부분이 없으므로 만들어 준다.
		str := s.Properties.GetFirst("StartTime")
		end := t.Header.Timestamp // 쿼리는 끝나는 시간을 안찍어주므로
		s.Properties.Add("EndTime", end, NORMAL)
		s.Properties.Add("RunTime", DiffMiliseconds(str, end), NORMAL)

		return UGT
	}
	Merge(t, s.Properties, s.GlobalProperties)
	return ACK
}

func (s *QueryObject) ToString() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	propertiesString := s.toStringProperties()
	out.WriteString(propertiesString)
	out.WriteString("\n}\n")
	return out.String()
}

func (s *QueryObject) toStringProperties() string {
	var out bytes.Buffer

	out.WriteString("\"property\":\n")
	propertiesString := s.Properties.ToString()

	out.WriteString(propertiesString)
	return out.String()
}

func (s *QueryObject) GetObjectType() string {
	return "Query"
}

func (s *QueryObject) GetChildren() []*Node {
	return nil
}
