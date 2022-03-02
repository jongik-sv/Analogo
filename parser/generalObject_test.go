package parser

import (
	"fmt"
	"goproject/AnaloGo/lexer/logHeader"
	"goproject/AnaloGo/lexer/token"
	"strings"
	"testing"

	_ "github.com/stretchr/testify/assert"
)

func TestLastIndex(t *testing.T) {
	tokenType := "PlSubProcess_Start"
	index := strings.LastIndex(tokenType, "_")
	println(index)
	println(tokenType[:index])
	println(tokenType[index+1:])
}

// func TestSendChild(str string, t *testing.T) {
// 	assert := assert.New(t)
// 	assert.Equal(1, 1)
// }

func TestBuildAST1(t *testing.T) {
	tokens := []token.Token{
		{TokenType: "Service_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Service [mo] start. Request Tag [ZQQ]"}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch", Text: "Process [Process_05xymvn](Process_05xymvn) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Activity_1obyfx7](parallelGreeting) start."}},
		{TokenType: "SubServ_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Sub-service [parallelGreeting] start."}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch", Text: "Process [Process_05xymvn](Process_05xymvn) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Gateway_0qw4j58](Gateway_0qw4j58) start."}},
		{TokenType: "PlSubProcess", Header: &logHeader.LogHeader{Key: "hgch", Text: "Parallel Sub-process run start. [Gateway_0qw4j58]"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Task [Activity_1msngyv](인사1 프로세스) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Activity_0nlxb65](인사2 메인 프로세스) start."}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Process [Activity_0nlxb65](인사2 메인 프로세스) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Gateway_111cvzu](Gateway_111cvzu) start."}},
		{TokenType: "PlSubProcess", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Parallel Sub-process run start. [Gateway_111cvzu]"}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Process [Activity_1msngyv](인사1 프로세스) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Task [Activity_0o9fhf8](인사1) start."}},
		{TokenType: "TaskInvokingClass", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Task [Activity_0o9fhf8](인사1) finish.(0ms)"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Task [Activity_1msngyv](인사1 프로세스) finish.(5ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Task [Activity_0mq11dx](인사2 프로세스) start."}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Process [Activity_0mq11dx](인사2 프로세스) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Task [Activity_1o33vaa](인사3 프로세스) start."}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Process [Activity_1o33vaa](인사3 프로세스) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Task [Activity_0wmahw6](인사2) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Task [Activity_0s0rbua](인사3) start."}},
		{TokenType: "TaskInvokingClass", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Task [Activity_0wmahw6](인사2) finish.(0ms)"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Task [Activity_0mq11dx](인사2 프로세스) finish.(1ms)"}},
		{TokenType: "TaskInvokingClass", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Task [Activity_0s0rbua](인사3) finish.(6ms)"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Task [Activity_1o33vaa](인사3 프로세스) finish.(7ms)"}},
		{TokenType: "PlSubProcess", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Parallel Sub-process run finish. [Gateway_111cvzu]"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Gateway_111cvzu](Gateway_111cvzu) finish.(26ms)"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Activity_0nlxb65](인사2 메인 프로세스) finish.(27ms)"}},
		{TokenType: "PlSubProcess", Header: &logHeader.LogHeader{Key: "hgch", Text: "Parallel Sub-process run finish. [Gateway_0qw4j58]"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Gateway_0qw4j58](Gateway_0qw4j58) finish.(41ms)"}},
		{TokenType: "SubServ_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Sub-service [parallelGreeting] finish."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Activity_1obyfx7](parallelGreeting) finish.(61ms)"}},
		{TokenType: "Service_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Service [mo] finish.(85ms)"}},
	}

	// fmt.Println(tokens)
	// sendChildChan := make(chan token.Token, 1000)
	// for _, v := range tokens {
	// 	sendChildChan <- v
	// }

	globalProperties := NewProperties()
	obj := NewGeneralObject("Request", "hgch", globalProperties)
	for i, v := range tokens {
		fmt.Println(i, v)
		obj.Parse(&v)

	}

	service := obj.Children[0]
	fmt.Println((*service).ToString())

}

func TestBuildAST(t *testing.T) {
	tokens := []token.Token{
		{TokenType: "POST", Header: &logHeader.LogHeader{Key: "hgch", Text: "POST \"/pmf/service/mo?moNo=M000164&action=subservice&name=JeongjinKim\", parameters={masked}"}},
		{TokenType: "Service_Start", Properties: map[string]string{"Name": "mo", "Tag": "ZQQ"}, Header: &logHeader.LogHeader{Key: "hgch", Text: "Service [mo] start. Request Tag [ZQQ]"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch", Text: "Transaction [txBiz] started."}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch", Text: "Process [Process_05xymvn](Process_05xymvn) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [StartEvent_1](StartEvent_1) start."}},
		{TokenType: "Task_End", Properties: map[string]string{"Name": "StartEvent_1", "runtime": "1"}, Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [StartEvent_1](StartEvent_1) finish.(0ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Gateway_1xtcgie](Gateway_1xtcgie) start."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Gateway_1xtcgie](Gateway_1xtcgie) finish.(0ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Activity_1obyfx7](parallelGreeting) start."}},
		{TokenType: "SubServ_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Sub-service [parallelGreeting] start."}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch", Text: "Process [Process_05xymvn](Process_05xymvn) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [StartEvent_1](StartEvent_1) start."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [StartEvent_1](StartEvent_1) finish.(0ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Gateway_1udd4th](Gateway_1udd4th) start."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Gateway_1udd4th](Gateway_1udd4th) finish.(0ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Gateway_0qw4j58](Gateway_0qw4j58) start."}},
		{TokenType: "PlSubProcess", Header: &logHeader.LogHeader{Key: "hgch", Text: "Parallel Sub-process run start. [Gateway_0qw4j58]"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Task [Activity_1msngyv](인사1 프로세스) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Activity_0nlxb65](인사2 메인 프로세스) start."}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Process [Activity_0nlxb65](인사2 메인 프로세스) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Event_1smi0ye](Event_1smi0ye) start."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Event_1smi0ye](Event_1smi0ye) finish.(0ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Gateway_0xdapvd](Gateway_0xdapvd) start."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Gateway_0xdapvd](Gateway_0xdapvd) finish.(0ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Gateway_111cvzu](Gateway_111cvzu) start."}},
		{TokenType: "PlSubProcess", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Parallel Sub-process run start. [Gateway_111cvzu]"}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Process [Activity_1msngyv](인사1 프로세스) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Task [Activity_0o9fhf8](인사1) start."}},
		{TokenType: "TaskInvokingClass", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Parameter name binding of name"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "JeongjinKim, hello."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Task [Activity_0o9fhf8](인사1) finish.(0ms)"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:rN5f", Text: "Task [Activity_1msngyv](인사1 프로세스) finish.(5ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Task [Activity_0mq11dx](인사2 프로세스) start."}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Process [Activity_0mq11dx](인사2 프로세스) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Task [Activity_1o33vaa](인사3 프로세스) start."}},
		{TokenType: "Process", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Process [Activity_1o33vaa](인사3 프로세스) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Task [Activity_0wmahw6](인사2) start."}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Task [Activity_0s0rbua](인사3) start."}},
		{TokenType: "TaskInvokingClass", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Parameter name binding of name"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "JeongjinKim, hello."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Task [Activity_0wmahw6](인사2) finish.(0ms)"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c:X6Px", Text: "Task [Activity_0mq11dx](인사2 프로세스) finish.(1ms)"}},
		{TokenType: "TaskInvokingClass", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Parameter name binding of name"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "JeongjinKim, hello."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Task [Activity_0s0rbua](인사3) finish.(6ms)"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c:Ba1e", Text: "Task [Activity_1o33vaa](인사3 프로세스) finish.(7ms)"}},
		{TokenType: "PlSubProcess", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Parallel Sub-process run finish. [Gateway_111cvzu]"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Gateway_111cvzu](Gateway_111cvzu) finish.(26ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Event_1035s1u](Event_1035s1u) start."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Event_1035s1u](Event_1035s1u) finish.(0ms)"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch:XB0c", Text: "Task [Activity_0nlxb65](인사2 메인 프로세스) finish.(27ms)"}},
		{TokenType: "PlSubProcess", Header: &logHeader.LogHeader{Key: "hgch", Text: "Parallel Sub-process run finish. [Gateway_0qw4j58]"}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Gateway_0qw4j58](Gateway_0qw4j58) finish.(41ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Event_0t44im8](Event_0t44im8) start."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Event_0t44im8](Event_0t44im8) finish.(0ms)"}},
		{TokenType: "SubServ_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Sub-service [parallelGreeting] finish."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Activity_1obyfx7](parallelGreeting) finish.(61ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Activity_0591vg1](MO 조회) start."}},
		{TokenType: "Query", Properties: map[string]string{"SQL": "select 1 from dual where a = :a"}, Header: &logHeader.LogHeader{Key: "hgch", Text: "Executing prepared SQL statement [select /* sample */mo_no, weight from mo_hea where mo_no = ?]"}},
		{TokenType: "Parameter", Properties: map[string]string{"Value": "value"}, Header: &logHeader.LogHeader{Key: "hgch", Text: "Setting SQL statement parameter value: column index 1, parameter value [M000164], value class [java.lang.String], SQL type unknown"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch", Text: "1 row(s) selected."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Activity_0591vg1](MO 조회) finish.(5ms)"}},
		{TokenType: "Task_Start", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Event_19hiwi1](Event_19hiwi1) start."}},
		{TokenType: "Task_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Task [Event_19hiwi1](Event_19hiwi1) finish.(0ms)"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch", Text: "Transaction [txBiz] has been committed."}},
		{TokenType: "Service_End", Header: &logHeader.LogHeader{Key: "hgch", Text: "Service [mo] finish.(85ms)"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch", Text: "Using 'application/json', given [*/*] and supported [application/json, application/*+json, application/json, application/*+json]"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch", Text: "Writing [{mos=TypedObject{object=[{mo_no=M000164, weight=12300}], type=java.util.List<java.util.Map<java.lang (truncated)...]"}},
		{TokenType: "X", Header: &logHeader.LogHeader{Key: "hgch", Text: "Completed 200 OK"}},
	}

	globalProperties := NewProperties()
	obj := NewGeneralObject("Zero", "hgch", globalProperties)
	for _, v := range tokens {
		obj.Parse(&v)
	}
	// service := obj.Children[0]
	fmt.Println((*obj).ToString())

}
