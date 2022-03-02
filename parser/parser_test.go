package parser

import (
	"fmt"
	"goproject/AnaloGo/lexer/logHeader"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenFromLexer(t *testing.T) {
	log := "[2021-12-10 18:16:44,718][INFO ][L454][com.dongkuk.oasis.service.SpringServiceStarter]-Service [mo] start."
	header := logHeader.GetLogHeader(log)
	ch := make(chan logHeader.LogHeader, 10)
	ch <- *header
	// p := Parser{ReciveChan: ch}
	// p.Lexer = &lexer.Lexer{
	// 	Server:     p.Server,
	// 	Chain:      p.Chain,
	// 	ReciveChan: p.ReciveChan,
	// }
	// tk := p.Lexer.GetToken()
	// tk.Print()

}

// const (
// 		NORMAL = iota
// 		ARRAY
// 		SET
// 	)

func TestGetTargetKey(t *testing.T) {
	targetSelector, targetMapKey, targetDataType := getTargetKey("Parameter.ARRAY")
	assert.Equal(t, "Local", targetSelector)
	assert.Equal(t, "Parameter", targetMapKey)
	assert.Equal(t, ARRAY, targetDataType)

}
func TestGetSourceKey(t *testing.T) {

}

func TestTimeCalc(t *testing.T) {
	str := "2021-12-23 11:36:41,871"
	end := "2021-12-23 11:36:45,342"

	st, err := convertTime(str)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	et, err := convertTime(end)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	assert.Equal(t, "3471", strconv.Itoa(int(et.Sub(st)/time.Millisecond)))
	// tm, err := strconv.ParseFloat(et.Sub(st), 64)
	// fmt.Println(st, et, int64(et.Sub(st)/time.Millisecond))

}
