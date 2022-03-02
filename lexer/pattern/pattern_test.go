package pattern

import (
	"fmt"
	"goproject/AnaloGo/lexer/logHeader"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegMatch(t *testing.T) {
	r := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)
	fmt.Printf("%#v\n", r.FindStringSubmatch(`2015-05-27`))
	fmt.Printf("%#v\n", r.SubexpNames())

}

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	tt := &TokenPatterns{}
	tt.Config("./lex_pattern_sample.json")
	assert.Equal(len(tt.LexPatterns), 3)

	// for i, v := range tt.LexPatterns {
	// 	fmt.Println(i, "TokenType : ", v.TokenType)
	// 	fmt.Println(i, "Discard : ", v.Discard)
	// 	fmt.Println(i, "State : ", v.State)
	// 	fmt.Println(i, "StringMatcher.Input : ", v.StringMatcher.Input)
	// 	fmt.Println(i, "StringMatcher.ContainText : ", v.StringMatcher.ContainText)
	// 	fmt.Println(i, "RegexMatcher.Input : ", v.RegexMatcher.Input)
	// 	fmt.Println(i, "RegexMatcher.Test : ", v.RegexMatcher.Test)
	// 	fmt.Println(i, "RegexMatcher.Regex : ", v.RegexMatcher.Regexp)
	// 	fmt.Println(i, "RegexMatcher.MatchingKeys : ", v.RegexMatcher.MatchingKeys)
	// }
}

func TestStringMatch(t *testing.T) {
	assert := assert.New(t)

	s := &StringMatcher{ContainText: []string{"Service ", "start."}}

	assert.Equal(s.match("Service [mo] start."), true)
	assert.Equal(s.match("Service [mo]"), false)
	assert.Equal(s.match("[mo] start."), false)
	assert.Equal(s.match("Servicestart."), false)
	assert.Equal(s.match("Service [mo] start.."), true)

	s = &StringMatcher{ContainText: []string{}}
	assert.Equal(s.match("Service [mo] start."), true)
	assert.Equal(s.match("Service [mo]"), true)
	assert.Equal(s.match("[mo] start."), true)
	assert.Equal(s.match("Servicestart."), true)
	assert.Equal(s.match("Service [mo] start.."), true)

}

func TestRegexMatch(t *testing.T) {
	assert := assert.New(t)

	r := &RegexMatcher{
		Test:         "Service \\[(.*)\\] finish.*\\(([0-9]+)ms\\)$",
		MatchingKeys: []string{"Service.ServiceName", "Service.RunTime"},
	}

	r.Regexp = regexp.MustCompile(r.Test)

	m, ok := r.match("Service [mo] finish with exceptions.(569ms)")

	assert.Equal(ok, true)
	assert.Equal(len(m), 2)
	// assert.Equal(m[0], "Service [mo] finish with exceptions.(569ms)")
	assert.Equal(m[0], "mo")
	assert.Equal(m[1], "569")
}

func TestTokenMatch(t *testing.T) {
	assert := assert.New(t)

	tm := &TokenMatcher{
		TokenType:     "Service_End",
		Discard:       false,
		State:         "F",
		StringMatcher: StringMatcher{ContainText: []string{"Service ", "finish"}},
		RegexMatcher: RegexMatcher{
			Test:         "Service \\[(.*)\\] finish.*\\(([0-9]+)ms\\)$",
			MatchingKeys: []string{"Service.ServiceName", "Service.RunTime"},
		},
	}

	tm.RegexMatcher.Regexp = regexp.MustCompile(tm.RegexMatcher.Test)
	m, ok := tm.match(&logHeader.LogHeader{Text: "Service [mo] finish with exceptions.(569ms)"})

	fmt.Println(m)
	assert.Equal(ok, true)
}

func TestTokenMatchPost(t *testing.T) {
	assert := assert.New(t)

	tm := &TokenMatcher{
		TokenType:     "Service_End",
		Discard:       false,
		State:         "F",
		StringMatcher: StringMatcher{ContainText: []string{"POST "}},
		RegexMatcher: RegexMatcher{
			Test:         "^POST `([^\"]*)`, parameters=\\{(.*)\\}",
			MatchingKeys: []string{"Address", "Parameters"},
		},
	}

	tm.RegexMatcher.Regexp = regexp.MustCompile(tm.RegexMatcher.Test)
	m, ok := tm.match(&logHeader.LogHeader{Text: "POST `/pmf/service/greeting?name=Kim`, parameters={masked}"})

	fmt.Println(m)
	assert.Equal(ok, true)
}

func TestTokenPattern(t *testing.T) {
	assert := assert.New(t)

	header := &logHeader.LogHeader{
		Key:  "AAAA",
		Text: "Service [mo] finish with exceptions.(569ms)",
		OK:   true,
	}

	tt := &TokenPatterns{}
	tt.Config("./lex_pattern_sample.json")
	// fmt.Println(tt)
	assert.Equal(len(tt.LexPatterns), 3)
	p, ok := tt.match(header)
	// fmt.Println(p)
	assert.Equal(ok, true)
	assert.Equal(len(p.Properties), 2)
	assert.Equal(p.Properties["Service.ServiceName"], "mo")
	assert.Equal(p.Properties["Service.RunTime"], "569")
	assert.Equal(p.Header.OK, true)
	assert.Equal(p.Status, "F")

	header = &logHeader.LogHeader{
		Key:   "AAAA",
		Level: "DEBUG",
		NDC:   "org.springframework.web.servlet.DispatcherServlet(108)",
		Text:  "POST \"/pmf/service/mo?moNo=M000164&action=subservice&name=JeongjinKim\", parameters={masked}",
		OK:    true,
	}
	p, _ = tt.match(header)
	assert.Equal(p.TokenType, "X")
	// fmt.Println(p)
}

func TestConfig2(t *testing.T) {
	// assert := assert.New(t)

	tt := &TokenPatterns{}
	tt.Config("./lex_pattern.json")
	// assert.Equal(len(tt.LexPatterns), 2)

	// for i, v := range tt.LexPatterns {
	// 	fmt.Println(i, "TokenType : ", v.TokenType)
	// 	fmt.Println(i, "Discard : ", v.Discard)
	// 	fmt.Println(i, "State : ", v.State)
	// 	fmt.Println(i, "StringMatcher.Input : ", v.StringMatcher.Input)
	// 	fmt.Println(i, "StringMatcher.ContainText : ", v.StringMatcher.ContainText)
	// 	fmt.Println(i, "RegexMatcher.Input : ", v.RegexMatcher.Input)
	// 	fmt.Println(i, "RegexMatcher.Test : ", v.RegexMatcher.Test)
	// 	fmt.Println(i, "RegexMatcher.Regex : ", v.RegexMatcher.Regexp)
	// 	fmt.Println(i, "RegexMatcher.MatchingKeys : ", v.RegexMatcher.MatchingKeys)
	// }
}
