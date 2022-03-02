package pattern

import (
	"encoding/json"
	"fmt"
	"goproject/AnaloGo/lexer/logHeader"
	"goproject/AnaloGo/lexer/token"
	"os"

	"io/ioutil"
	"regexp"
	"strings"
)

type StringMatcher struct {
	Input       string   `json:"input"`
	ContainText []string `json:"containText"`
}
type RegexMatcher struct {
	Input        string `json:"input"`
	Test         string `json:"test"`
	Regexp       *regexp.Regexp
	MatchingKeys []string `json:"matchingKeys"`
}

type TokenMatcher struct {
	TokenType  string            `json:"tokenType"`
	Discard    bool              `json:"discard"`
	State      string            `json:"state"` // S, F, nil
	CopyObject map[string]string `json:"copyObject"`

	StringMatcher StringMatcher `json:"stringMatcher"`
	ExtMatcher1   StringMatcher `json:"extraMatcher1"`
	RegexMatcher  RegexMatcher  `json:"regexMatcher"`
	// PostProcessor struct {
	// 	ServiceName string `json:"ServiceName"`
	// } `json:"postProcessor"`

}

type TokenPatterns struct {
	LexPatterns []TokenMatcher `json:"tokens"`
}

var tokenPatterns TokenPatterns

func init() {
	path, _ := os.Getwd()
	println(path)
	fmt.Println("@@@@ pattern.init() @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	tokenPatterns = TokenPatterns{}
	tokenPatterns.Config("c:/project/goproject/AnaloGo/analogo/lex_pattern.json")
}

func (tp *TokenPatterns) Config(filePath string) {
	dat, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(dat, tp)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(tp.LexPatterns); i++ {
		if tp.LexPatterns[i].RegexMatcher.Test != "" {
			tp.LexPatterns[i].RegexMatcher.Regexp = regexp.MustCompile(tp.LexPatterns[i].RegexMatcher.Test)
		}
	}
}

func (s *StringMatcher) match(text string) bool {

	if text == "" {
		return true
	}
	if len(s.ContainText) == 0 {
		return true
	}
	for _, substr := range s.ContainText {
		if !strings.Contains(text, substr) {
			return false
		}
	}
	return true
}

func (r *RegexMatcher) match(text string) ([]string, bool) {

	if r.Regexp == nil {
		return nil, true
	}

	m := r.Regexp.FindStringSubmatch(text)

	if len(m) == 0 {
		return nil, false
	}

	if len(m) != len(r.MatchingKeys)+1 {
		// err := fmt.Errorf("[%s][%v] 의 매칭 건수[%d]와 매칭키의 갯수[%d]가 다릅니다",
		// 	r.Test, r.MatchingKeys, len(m)-1, len(r.MatchingKeys))
		// panic(err)
		fmt.Printf("    [%s][%v] 의 매칭 건수[%d]와 매칭키의 갯수[%d]가 다릅니다. %s %v\n", r.Test, r.MatchingKeys, len(m)-1, len(r.MatchingKeys), text, m)
		// return m[1:], false
		return nil, true
	}
	return m[1:], true
}

func (tm *TokenMatcher) match(header *logHeader.LogHeader) (map[string]string, bool) {

	text := header.Text

	if !tm.StringMatcher.match(text) {
		return nil, false
	}

	input := ""
	if tm.ExtMatcher1.Input != "" {
		switch tm.ExtMatcher1.Input {
		case "NDC":
			input = header.NDC
		case "Level":
			input = header.Level
		case "TotalString":
			input = header.TotalString
		}

		if !tm.ExtMatcher1.match(input) {
			return nil, false
		}
	}
	m, ok := tm.RegexMatcher.match(text)
	if ok {
		if len(m) == 0 {
			return nil, true
		}

		matchMap := make(map[string]string)

		for i, v := range m {

			matchMap[tm.RegexMatcher.MatchingKeys[i]] = v
		}
		return matchMap, true
	}
	return nil, false
}

func (tp *TokenPatterns) match(header *logHeader.LogHeader) (*token.Token, bool) {

	if header.Key == "" {
		return nil, false
	}

	// text := header.Text
	for _, v := range tp.LexPatterns {
		if m, ok := v.match(header); ok {
			token := &token.Token{
				TokenType:  v.TokenType,
				Header:     header,
				Properties: m,
				Status:     v.State,
				CopyObject: &v.CopyObject,
			}
			return token, ok
		}
	}
	// fmt.Println("패턴 검색에 실패 했습니다.", header.TotalString)
	token := &token.Token{
		TokenType:  "X",
		Header:     header,
		Properties: nil,
		Status:     "",
		CopyObject: nil,
	}
	return token, false
}

func Match(header *logHeader.LogHeader) (*token.Token, bool) {
	return tokenPatterns.match(header)
}
