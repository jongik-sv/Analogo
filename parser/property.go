package parser

import (
	"bytes"
	"fmt"
	"strings"
)

// type Property interface {
// 	ToString() string
// 	add(value string)
// }

const (
	NORMAL = iota
	ARRAY
	SET
)

type Property struct {
	dataType int
	value    []string
}

func (p Property) ToString() string {
	var out bytes.Buffer

	if len(p.value) == 0 {
		return ""
	}

	if p.dataType == NORMAL {
		out.WriteString(fmt.Sprintf("\"%s\"", p.value[0]))
	} else {
		out.WriteString("[")

		var isFirst bool = true
		for _, v := range p.value {
			if isFirst {
				out.WriteString(fmt.Sprintf("\"%s\"", v))
				isFirst = false
			} else {
				out.WriteString(fmt.Sprintf(", \"%s\"", v))
			}
		}
		out.WriteString("]")
	}
	return out.String()
}

func (p *Property) append(value string) {
	p.value = append(p.value, value)
}

func (p *Property) update(value string) {
	p.value[0] = value
}

func (p *Property) set(value string) {
	for _, v := range p.value {
		if v == value {
			return
		}
	}
	p.value = append(p.value, value)
}

type Properties struct {
	propertyMap map[string]*Property
}

func (ps *Properties) Get(key string) []string {
	// pm := ps.mode(key)
	if v, ok := ps.propertyMap[key]; !ok {
		return nil
	} else {
		return v.value
	}
}

func (ps *Properties) GetJoinString(key string) string {
	v := ps.Get(key)
	if len(v) >= 1 {
		return strings.Join(v, ", ")
	}
	return ""
}

func (ps *Properties) GetFirst(key string) string {
	v := ps.Get(key)
	if len(v) >= 1 {
		return v[0]
	}
	return ""
}

func (ps *Properties) GetLast(key string) string {
	v := ps.Get(key)
	l := len(v)
	if l >= 1 {
		return v[l-1]
	}
	return ""
}

func (ps *Properties) GetSliceThroughSetMethod(key string) []string {
	sliceValue := ps.Get(key)
	set := make(map[string]int)
	for _, v := range sliceValue {
		set[v] = set[v] + 1
	}

	newKeySlice := []string{}
	for k := range set {
		newKeySlice = append(newKeySlice, k)
	}

	return newKeySlice

}

func (ps *Properties) Add(key, value string, dataType int) {
	if v, ok := ps.propertyMap[key]; ok {
		if v.dataType != dataType {
			fmt.Println(key, value, dataType, "데이터 타입이 맞지 않습니다.")
		}
		if dataType == ARRAY {
			(*v).append(value)
		} else if dataType == SET {
			(*v).set(value)
		} else if dataType == NORMAL {
			(*v).update(value)
		}

	} else {
		property := Property{value: []string{value}, dataType: dataType}
		ps.propertyMap[key] = &property
	}
}

func NewProperties() *Properties {
	return &Properties{propertyMap: make(map[string]*Property)}
}

func (ps *Properties) ToString_old() string {
	var out bytes.Buffer
	mapLen := len(ps.propertyMap)
	mapCnt := 0
	if mapLen == 0 {
		return "{}"
	}

	out.WriteString("{\n")
	for key, element := range ps.propertyMap {
		mapCnt++
		if element.ToString() == "" {
			continue
		}
		if mapCnt == mapLen {
			out.WriteString(fmt.Sprintf("\"%s\": %s\n", key, element.ToString()))
		} else {
			out.WriteString(fmt.Sprintf("\"%s\": %s,\n", key, element.ToString()))
		}
	}
	out.WriteString("}")
	return out.String()
}

func (ps *Properties) ToString() string {
	var out bytes.Buffer
	isFirst := true
	out.WriteString("{\n")

	for _, key := range propertyOrder[ps.GetFirst("ObjectType")] {

		if value, ok := ps.propertyMap[key]; ok {
			if isFirst {
				isFirst = false
				out.WriteString("\n")
			} else {
				out.WriteString(",\n")
			}
			out.WriteString(fmt.Sprintf("\"%s\": %s", key, value.ToString()))
		}
	}

	out.WriteString("\n}")
	return out.String()
}

/*
	for key, element := range ps.propertyMap {
		if element.ToString() == "" {
			continue
		}
		if mapCnt == mapLen {
			out.WriteString(fmt.Sprintf("\"%s\": %s\n", key, element.ToString()))
		} else {
			out.WriteString(fmt.Sprintf("\"%s\": %s,\n", key, element.ToString()))
		}
	}
*/
