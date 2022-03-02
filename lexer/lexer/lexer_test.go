package lexer

import (
	"fmt"
	"goproject/AnaloGo/lexer/logHeader"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// "goproject/AnaloGo/lexer/token"

func TestMatch(t *testing.T) {
	header := logHeader.LogHeader{
		Server:      "운영계1",
		Chain:       "",
		Timestamp:   "2021-12-10 18:16:44,718",
		Level:       "INFO",
		Key:         "L454",
		Keys:        []string{"L454"},
		NDC:         "com.dongkuk.oasis.service.SpringServiceStarter(???)",
		Text:        "Service [mo] start.",
		TotalString: "[2021-12-10 18:16:44,718][INFO ][L454][com.dongkuk.oasis.service.SpringServiceStarter]-Service [mo] start.",
	}

	lexer := &Lexer{}
	token := lexer.match(&header)
	fmt.Println(*token)
	assert.Equal(t, token.TokenType, "Service_Start")

}

func TestMatch2(t *testing.T) {
	header := logHeader.LogHeader{
		Server:      "운영계1",
		Chain:       "",
		Timestamp:   "2021-12-23 11:36:41,906",
		Level:       "DEBUG",
		Key:         "E75f",
		Keys:        []string{"E75f"},
		NDC:         "org.springframework.web.servlet.mvc.method.annotation.RequestResponseBodyMethodProcessor(108)",
		Text:        "Writing [{hello=TypedObject{object=Kim, hello., type=class java.lang.String}}]",
		TotalString: "[2021-12-23 11:36:41,906][DEBUG][E75f][org.springframework.web.servlet.mvc.method.annotation.RequestResponseBodyMethodProcessor(108)]-Writing [{hello=TypedObject{object=Kim, hello., type=class java.lang.String}}]",
	}

	lexer := &Lexer{}
	token := lexer.match(&header)
	fmt.Println(*token)
	fmt.Println(token.Header.Text)
	assert.Equal(t, token.TokenType, "X")

}

func TestGetToken(t *testing.T) {

	header1 := logHeader.LogHeader{
		Server:      "운영계1",
		Chain:       "C10",
		Timestamp:   "2021-12-13 11:49:21,162",
		Level:       "DEBUG",
		Key:         "KFH7:dsdsd:dse34",
		Keys:        []string{"KFH7", "dsdsd", "dse34"},
		NDC:         "org.springframework.web.servlet.DispatcherServlet",
		Text:        "Failed to complete request: java.lang.RuntimeException: my exception",
		TotalString: "[2021-12-13 11:49:21,162][DEBUG][KFH7:dsdsd:dse34][org.springframework.web.servlet.DispatcherServlet]-Failed to complete request: java.lang.RuntimeException: my exception",
	}
	header2 := logHeader.LogHeader{
		Server:      "운영계1",
		Chain:       "C10",
		Timestamp:   "2021-12-13 11:49:21,162",
		Level:       "DEBUG",
		Key:         "aaa",
		Keys:        []string{"aaa"},
		NDC:         "org.springframework.web.servlet.DispatcherServlet",
		Text:        "Failed to complete request: java.lang.RuntimeException: my exception",
		TotalString: "[2021-12-13 11:49:21,162][DEBUG][aaa][org.springframework.web.servlet.DispatcherServlet]-Failed to complete request: java.lang.RuntimeException: my exception",
	}

	lexer := Lexer{ReciveChan: make(chan logHeader.LogHeader, 10), ExpireTime: 2}
	lexer.ReciveChan <- header1
	lexer.ReciveChan <- header2
	token := lexer.GetToken()
	fmt.Println(token.TokenType)
	token = lexer.GetToken()
	fmt.Println(token.TokenType)
	// assert.NotEqual(t, token.TokenType, token.NONE_KEY)
	token = lexer.GetToken()
	assert.Equal(t, token.TokenType, "TIMEOUT")
}

func TestTokenPrint(t *testing.T) {
	logs := []string{
		"[2021-12-10 18:16:44,718][INFO ][L454][com.dongkuk.oasis.service.SpringServiceStarter]-Service [mo] start.",
		"[2021-12-23 11:36:42,273][INFO ][kNgo][com.dongkuk.oasis.service.SpringServiceStarter(37)]-Service [mo] finish.(156ms)",
	}
	for _, log := range logs {
		header := logHeader.GetLogHeader(log)
		ch := make(chan logHeader.LogHeader, 10)
		ch <- *header
		l := &Lexer{ReciveChan: ch}
		tk := l.GetToken()
		tk.Print()
	}

}

func TestBrokerEmul(t *testing.T) {
	logs := []string{
		`[2021-12-23 11:39:39,181][DEBUG][pU66][org.hibernate.SQL(144)]-
			select
				next_val as id_val 
			from
				hibernate_sequence for update
					`,
		`[2021-12-23 11:39:39,184][DEBUG][pU66][org.hibernate.SQL(144)]-
			update
				hibernate_sequence 
			set
				next_val= ? 
			where
				next_val=?`,
		`[2021-12-23 11:39:39,188][DEBUG][pU66][org.hibernate.SQL(144)]-
			insert 
			into
				request_log
				(content, id) 
			values
				(?, ?)`,
		`[2021-12-23 11:39:39,189][TRACE][pU66][org.hibernate.type.descriptor.sql.BasicBinder(64)]-binding parameter [1] as [VARCHAR] - [payload:[null]]`,
		`[2021-12-23 11:39:39,189][TRACE][pU66][org.hibernate.type.descriptor.sql.BasicBinder(64)]-binding parameter [2] as [BIGINT] - [3]`,
		`[2021-12-23 11:39:39,193][INFO ][pU66][com.dongkuk.oasis.service.SpringServiceStarter(23)]-Service [mo] start. Request Tag [WAh]`,
		`[2021-12-23 11:39:39,214][INFO ][pU66][com.dongkuk.oasis.transaction.SpringTransactionHandler(82)]-Transaction [txBiz] started.`,
		`[2021-12-23 11:39:39,215][INFO ][pU66][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Process_05xymvn](Process_05xymvn) start.`,
		`[2021-12-23 11:39:39,215][INFO ][pU66][com.dongkuk.oasis.model.event.DefaultStartEvent(24)]-Task [StartEvent_1](StartEvent_1) start.`,
		`[2021-12-23 11:39:39,215][INFO ][pU66][com.dongkuk.oasis.model.event.DefaultStartEvent(31)]-Task [StartEvent_1](StartEvent_1) finish.(0ms)`,
		`[2021-12-23 11:39:39,215][INFO ][pU66][com.dongkuk.oasis.model.gateway.DefaultExclusiveGateway(24)]-Task [Gateway_1xtcgie](Gateway_1xtcgie) start.`,
		`[2021-12-23 11:39:39,216][INFO ][pU66][com.dongkuk.oasis.model.gateway.DefaultExclusiveGateway(31)]-Task [Gateway_1xtcgie](Gateway_1xtcgie) finish.(0ms)`,
		`[2021-12-23 11:39:39,216][INFO ][pU66][com.dongkuk.oasis.model.activity.JavaServiceTask(24)]-Task [Activity_0qep0bz](MO등록) start.`,
		`[2021-12-23 11:39:39,217][INFO ][pU66][com.dongkuk.oasis.executors.PlainJavaServiceTaskExecutable(72)]-Invoking class : [com.dongkuk.dmes.cr.biz.plan.mo.Mo], method : [newMo]`,
		`[2021-12-23 11:39:39,253][INFO ][pU66][com.dongkuk.dmes.cr.frm.access.log.MybatisSqlLogger(37)]-Mybatis SQL : 
		/* com.dongkuk.dmes.cr.biz.plan.mo.MoRepository.addMo */
		insert into biz.mo_hea(mo_no, weight)
				values (?, ?)`,
		`[2021-12-23 11:39:39,253][INFO ][pU66][com.dongkuk.dmes.cr.frm.access.log.MybatisSqlLogger(43)]-binding parameter [0] as [java.lang.String] - [M000164]`,
		`[2021-12-23 11:39:39,253][INFO ][pU66][com.dongkuk.dmes.cr.frm.access.log.MybatisSqlLogger(43)]-binding parameter [1] as [java.lang.Integer] - [12300]`,
		`[2021-12-23 11:39:39,267][INFO ][pU66][com.dongkuk.oasis.model.activity.JavaServiceTask(31)]-Task [Activity_0qep0bz](MO등록) finish.(50ms)`,
		`[2021-12-23 11:39:39,268][INFO ][pU66][com.dongkuk.oasis.model.activity.SqlScriptTask(24)]-Task [Activity_1k6inig](MO 조회) start.`,
		`[2021-12-23 11:39:39,281][DEBUG][pU66][org.springframework.jdbc.core.JdbcTemplate(711)]-Executing prepared SQL query`,
		`[2021-12-23 11:39:39,282][DEBUG][pU66][org.springframework.jdbc.core.JdbcTemplate(643)]-Executing prepared SQL statement [select mo_no, weight from mo_hea where mo_no = ?]`,
		`[2021-12-23 11:39:39,284][TRACE][pU66][org.springframework.jdbc.core.StatementCreatorUtils(221)]-Setting SQL statement parameter value: column index 1, parameter value [M000164], value class [java.lang.String], SQL type unknown`,
		`[2021-12-23 11:39:39,295][INFO ][pU66][com.dongkuk.oasis.model.activity.SqlScriptTask(31)]-Task [Activity_1k6inig](MO 조회) finish.(27ms)`,
		`[2021-12-23 11:39:39,296][INFO ][pU66][com.dongkuk.oasis.model.event.DefaultEndEvent(24)]-Task [Event_0t44im8](Event_0t44im8) start.`,
		`[2021-12-23 11:39:39,296][INFO ][pU66][com.dongkuk.oasis.model.event.DefaultEndEvent(31)]-Task [Event_0t44im8](Event_0t44im8) finish.(0ms)`,
		`[2021-12-23 11:39:39,296][INFO ][pU66][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Process_05xymvn](Process_05xymvn) finish.(81ms)`,
		`[2021-12-23 11:39:39,301][INFO ][pU66][com.dongkuk.oasis.transaction.SpringTransactionHandler(117)]-Transaction [txBiz] has been committed.`,
		`[2021-12-23 11:39:39,301][INFO ][pU66][com.dongkuk.oasis.service.SpringServiceStarter(37)]-Service [mo] finish.(108ms)`,
	}

	toLexCh := make(chan logHeader.LogHeader, 100)

	lex := Lexer{
		Server:     "운영계1",
		Chain:      "AAA",
		Key:        "pU66",
		ReciveChan: toLexCh,
		ExpireTime: 3,
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		lex.Run()
		wg.Done()
	}()

	for _, logString := range logs {
		header := logHeader.GetLogHeader(logString)
		toLexCh <- *header
	}

	wg.Wait()
}

func TestBrokerEmulWithErrorException(t *testing.T) {
	logs := []string{
		`[2021-12-23 11:39:39,449][DEBUG][Nkut][org.hibernate.SQL(144)]-
			select
				next_val as id_val 
			from
				hibernate_sequence for update
					`,
		`[2021-12-23 11:39:39,451][DEBUG][Nkut][org.hibernate.SQL(144)]-
			update
				hibernate_sequence 
			set
				next_val= ? 
			where
				next_val=?`,
		`[2021-12-23 11:39:39,455][DEBUG][Nkut][org.hibernate.SQL(144)]-
			insert 
			into
				request_log
				(content, id) 
			values
				(?, ?)`,
		`[2021-12-23 11:39:39,455][TRACE][Nkut][org.hibernate.type.descriptor.sql.BasicBinder(64)]-binding parameter [1] as [VARCHAR] - [payload:[null]]`,
		`[2021-12-23 11:39:39,456][TRACE][Nkut][org.hibernate.type.descriptor.sql.BasicBinder(64)]-binding parameter [2] as [BIGINT] - [5]`,
		`[2021-12-23 11:39:39,460][INFO ][Nkut][com.dongkuk.oasis.service.SpringServiceStarter(23)]-Service [mo] start. Request Tag [3ge]`,
		`[2021-12-23 11:39:39,471][INFO ][Nkut][com.dongkuk.oasis.transaction.SpringTransactionHandler(82)]-Transaction [txBiz] started.`,
		`[2021-12-23 11:39:39,471][INFO ][Nkut][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Process_05xymvn](Process_05xymvn) start.`,
		`[2021-12-23 11:39:39,471][INFO ][Nkut][com.dongkuk.oasis.model.event.DefaultStartEvent(24)]-Task [StartEvent_1](StartEvent_1) start.`,
		`[2021-12-23 11:39:39,471][INFO ][Nkut][com.dongkuk.oasis.model.event.DefaultStartEvent(31)]-Task [StartEvent_1](StartEvent_1) finish.(0ms)`,
		`[2021-12-23 11:39:39,472][INFO ][Nkut][com.dongkuk.oasis.model.gateway.DefaultExclusiveGateway(24)]-Task [Gateway_1xtcgie](Gateway_1xtcgie) start.`,
		`[2021-12-23 11:39:39,472][INFO ][Nkut][com.dongkuk.oasis.model.gateway.DefaultExclusiveGateway(31)]-Task [Gateway_1xtcgie](Gateway_1xtcgie) finish.(0ms)`,
		`[2021-12-23 11:39:39,472][INFO ][Nkut][com.dongkuk.oasis.model.activity.JavaServiceTask(24)]-Task [Activity_0esc4cw](MO등록_exception) start.`,
		`[2021-12-23 11:39:39,472][INFO ][Nkut][com.dongkuk.oasis.executors.PlainJavaServiceTaskExecutable(72)]-Invoking class : [com.dongkuk.dmes.cr.biz.plan.mo.Mo], method : [throwException]`,
		`[2021-12-23 11:39:39,473][INFO ][Nkut][com.dongkuk.oasis.model.activity.JavaServiceTask(41)]-Task [Activity_0esc4cw](MO등록_exception) finish with exceptions.(1ms)`,
		`[2021-12-23 11:39:39,474][INFO ][Nkut][com.dongkuk.oasis.process.CoreProcessStarter(39)]-Process [Process_05xymvn](Process_05xymvn) finish with exceptions.(2ms)`,
		`[2021-12-23 11:39:39,481][INFO ][Nkut][com.dongkuk.oasis.transaction.SpringTransactionHandler(157)]-Transaction [txBiz] has been rolled back.`,
		`[2021-12-23 11:39:39,482][ERROR][Nkut][com.dongkuk.oasis.service.StopWatchServiceStarter(39)]-my exception
		java.lang.RuntimeException: my exception
			at com.dongkuk.dmes.cr.biz.plan.mo.Mo.throwException(Mo.java:30) ~[main/:?]
			at jdk.internal.reflect.NativeMethodAccessorImpl.invoke0(Native Method) ~[?:?]
			at jdk.internal.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62) ~[?:?]
			at jdk.internal.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43) ~[?:?]
			at java.lang.reflect.Method.invoke(Method.java:566) ~[?:?]
			at io.github.thecodinglog.methodinvoker.StrictMethodInvoker.invoke(StrictMethodInvoker.java:60) ~[method-invoker-0.2.0.jar:?]
			at com.dongkuk.oasis.executors.PlainJavaServiceTaskExecutable.execute(PlainJavaServiceTaskExecutable.java:94) ~[oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.executors.JavaServiceTaskExecutable.execute(JavaServiceTaskExecutable.java:46) ~[oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.executors.CoreElementExecutor.execute(CoreElementExecutor.java:78) ~[oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.executors.StopWatchElementExecutor.execute(StopWatchElementExecutor.java:36) ~[oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.process.CoreProcessStarter.start(CoreProcessStarter.java:99) ~[oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.process.AbstractStopWatchProcessStarter.start(AbstractStopWatchProcessStarter.java:32) ~[oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.process.StopWatchProcessStarter.start(StopWatchProcessStarter.java:26) ~[oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.service.CoreServiceStarter.lambda$start$3(CoreServiceStarter.java:94) ~[oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.transaction.SpringTransactionHandler.execute(SpringTransactionHandler.java:215) ~[oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.service.CoreServiceStarter.start(CoreServiceStarter.java:93) ~[oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.service.SpringServiceStarter.start(SpringServiceStarter.java:33) ~[oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.service.WatchServiceStarter.logAndServiceStart(WatchServiceStarter.java:36) [oasis-core-4.2.0.jar:?]
			at com.dongkuk.oasis.service.StopWatchServiceStarter.start(StopWatchServiceStarter.java:29) [oasis-core-4.2.0.jar:?]
			at com.dongkuk.dmes.cr.frm.inbound.ServiceController.service(ServiceController.java:56) [dmes-cr-frm-config-0.0.1.jar:?]
			at jdk.internal.reflect.NativeMethodAccessorImpl.invoke0(Native Method) ~[?:?]
			at jdk.internal.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62) ~[?:?]
			at jdk.internal.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43) ~[?:?]
			at java.lang.reflect.Method.invoke(Method.java:566) ~[?:?]
			at org.springframework.web.method.support.InvocableHandlerMethod.doInvoke(InvocableHandlerMethod.java:205) [spring-web-5.3.13.jar:5.3.13]
			at org.springframework.web.method.support.InvocableHandlerMethod.invokeForRequest(InvocableHandlerMethod.java:150) [spring-web-5.3.13.jar:5.3.13]
			at org.springframework.web.servlet.mvc.method.annotation.ServletInvocableHandlerMethod.invokeAndHandle(ServletInvocableHandlerMethod.java:117) [spring-webmvc-5.3.13.jar:5.3.13]
			at org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerAdapter.invokeHandlerMethod(RequestMappingHandlerAdapter.java:895) [spring-webmvc-5.3.13.jar:5.3.13]
			at org.springframework.web.servlet.mvc.method.annotation.RequestMappingHandlerAdapter.handleInternal(RequestMappingHandlerAdapter.java:808) [spring-webmvc-5.3.13.jar:5.3.13]
			at org.springframework.web.servlet.mvc.method.AbstractHandlerMethodAdapter.handle(AbstractHandlerMethodAdapter.java:87) [spring-webmvc-5.3.13.jar:5.3.13]
			at org.springframework.web.servlet.DispatcherServlet.doDispatch(DispatcherServlet.java:1067) [spring-webmvc-5.3.13.jar:5.3.13]
			at org.springframework.web.servlet.DispatcherServlet.doService(DispatcherServlet.java:963) [spring-webmvc-5.3.13.jar:5.3.13]
			at org.springframework.web.servlet.FrameworkServlet.processRequest(FrameworkServlet.java:1006) [spring-webmvc-5.3.13.jar:5.3.13]
			at org.springframework.web.servlet.FrameworkServlet.doPost(FrameworkServlet.java:909) [spring-webmvc-5.3.13.jar:5.3.13]
			at javax.servlet.http.HttpServlet.service(HttpServlet.java:681) [tomcat-embed-core-9.0.55.jar:4.0.FR]
			at org.springframework.web.servlet.FrameworkServlet.service(FrameworkServlet.java:883) [spring-webmvc-5.3.13.jar:5.3.13]
			at javax.servlet.http.HttpServlet.service(HttpServlet.java:764) [tomcat-embed-core-9.0.55.jar:4.0.FR]
			at org.apache.catalina.core.ApplicationFilterChain.internalDoFilter(ApplicationFilterChain.java:227) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.core.ApplicationFilterChain.doFilter(ApplicationFilterChain.java:162) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.tomcat.websocket.server.WsFilter.doFilter(WsFilter.java:53) [tomcat-embed-websocket-9.0.55.jar:9.0.55]
			at org.apache.catalina.core.ApplicationFilterChain.internalDoFilter(ApplicationFilterChain.java:189) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.core.ApplicationFilterChain.doFilter(ApplicationFilterChain.java:162) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at com.dongkuk.dmes.cr.frm.filters.TraceIdLoggingFilter$1.process(TraceIdLoggingFilter.java:23) [dmes-cr-frm-config-0.0.1.jar:?]
			at com.dongkuk.oasis.logger.MDCTemplate.mdc(MDCTemplate.java:34) [oasis-core-4.2.0.jar:?]
			at com.dongkuk.dmes.cr.frm.filters.TraceIdLoggingFilter.doFilter(TraceIdLoggingFilter.java:29) [dmes-cr-frm-config-0.0.1.jar:?]
			at org.apache.catalina.core.ApplicationFilterChain.internalDoFilter(ApplicationFilterChain.java:189) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.core.ApplicationFilterChain.doFilter(ApplicationFilterChain.java:162) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.springframework.web.filter.RequestContextFilter.doFilterInternal(RequestContextFilter.java:100) [spring-web-5.3.13.jar:5.3.13]
			at org.springframework.web.filter.OncePerRequestFilter.doFilter(OncePerRequestFilter.java:119) [spring-web-5.3.13.jar:5.3.13]
			at org.apache.catalina.core.ApplicationFilterChain.internalDoFilter(ApplicationFilterChain.java:189) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.core.ApplicationFilterChain.doFilter(ApplicationFilterChain.java:162) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.springframework.web.filter.FormContentFilter.doFilterInternal(FormContentFilter.java:93) [spring-web-5.3.13.jar:5.3.13]
			at org.springframework.web.filter.OncePerRequestFilter.doFilter(OncePerRequestFilter.java:119) [spring-web-5.3.13.jar:5.3.13]
			at org.apache.catalina.core.ApplicationFilterChain.internalDoFilter(ApplicationFilterChain.java:189) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.core.ApplicationFilterChain.doFilter(ApplicationFilterChain.java:162) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.springframework.web.filter.CharacterEncodingFilter.doFilterInternal(CharacterEncodingFilter.java:201) [spring-web-5.3.13.jar:5.3.13]
			at org.springframework.web.filter.OncePerRequestFilter.doFilter(OncePerRequestFilter.java:119) [spring-web-5.3.13.jar:5.3.13]
			at org.apache.catalina.core.ApplicationFilterChain.internalDoFilter(ApplicationFilterChain.java:189) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.core.ApplicationFilterChain.doFilter(ApplicationFilterChain.java:162) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.core.StandardWrapperValve.invoke(StandardWrapperValve.java:197) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.core.StandardContextValve.invoke(StandardContextValve.java:97) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.authenticator.AuthenticatorBase.invoke(AuthenticatorBase.java:540) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.core.StandardHostValve.invoke(StandardHostValve.java:135) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.valves.ErrorReportValve.invoke(ErrorReportValve.java:92) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.core.StandardEngineValve.invoke(StandardEngineValve.java:78) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.catalina.connector.CoyoteAdapter.service(CoyoteAdapter.java:357) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.coyote.http11.Http11Processor.service(Http11Processor.java:382) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.coyote.AbstractProcessorLight.process(AbstractProcessorLight.java:65) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.coyote.AbstractProtocol$ConnectionHandler.process(AbstractProtocol.java:895) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.tomcat.util.net.NioEndpoint$SocketProcessor.doRun(NioEndpoint.java:1722) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.tomcat.util.net.SocketProcessorBase.run(SocketProcessorBase.java:49) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.tomcat.util.threads.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1191) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.tomcat.util.threads.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:659) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at org.apache.tomcat.util.threads.TaskThread$WrappingRunnable.run(TaskThread.java:61) [tomcat-embed-core-9.0.55.jar:9.0.55]
			at java.lang.Thread.run(Thread.java:829) [?:?]`,
		`[2021-12-23 11:39:39,496][INFO ][Nkut][com.dongkuk.oasis.service.SpringServiceStarter(46)]-Service [mo] finish with exceptions.(35ms)`,
	}

	toLexCh := make(chan logHeader.LogHeader, 100)

	lex := Lexer{
		Server:     "운영계1",
		Chain:      "AAA",
		Key:        "Nkut",
		ReciveChan: toLexCh,
		ExpireTime: 3,
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		lex.Run()
		wg.Done()
	}()

	for _, logString := range logs {
		header := logHeader.GetLogHeader(logString)
		toLexCh <- *header
	}

	wg.Wait()
}

func TestBrokerEmulExceptProcess(t *testing.T) {
	logs := []string{
		`[2021-12-23 11:39:40,122][DEBUG][wwAP][org.hibernate.SQL(144)]-
		select
			next_val as id_val 
		from
			hibernate_sequence for update
				`,
		`[2021-12-23 11:39:40,125][DEBUG][wwAP][org.hibernate.SQL(144)]-
		update
			hibernate_sequence 
		set
			next_val= ? 
		where
			next_val=?`,
		`[2021-12-23 11:39:40,128][DEBUG][wwAP][org.hibernate.SQL(144)]-
		insert 
		into
			request_log
			(content, id) 
		values
			(?, ?)`,
		`[2021-12-23 11:39:40,129][TRACE][wwAP][org.hibernate.type.descriptor.sql.BasicBinder(64)]-binding parameter [1] as [VARCHAR] - [payload:[null]]`,
		`[2021-12-23 11:39:40,129][TRACE][wwAP][org.hibernate.type.descriptor.sql.BasicBinder(64)]-binding parameter [2] as [BIGINT] - [9]`,
		`[2021-12-23 11:39:40,132][INFO ][wwAP][com.dongkuk.oasis.service.SpringServiceStarter(23)]-Service [mo] start. Request Tag [xXI]`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.transaction.SpringTransactionHandler(82)]-Transaction [txBiz] started.`,
		// `[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Process_05xymvn](Process_05xymvn) start.`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultStartEvent(24)]-Task [StartEvent_1](StartEvent_1) start.`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultStartEvent(31)]-Task [StartEvent_1](StartEvent_1) finish.(0ms)`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultExclusiveGateway(24)]-Task [Gateway_1xtcgie](Gateway_1xtcgie) start.`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultExclusiveGateway(31)]-Task [Gateway_1xtcgie](Gateway_1xtcgie) finish.(0ms)`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.model.activity.SubServiceCallTask(24)]-Task [Activity_1obyfx7](parallelGreeting) start.`,
		`[2021-12-23 11:39:40,148][INFO ][wwAP][com.dongkuk.oasis.executors.SubServiceCallTaskExecutable(68)]-Sub-service [parallelGreeting] start.`,
		// `[2021-12-23 11:39:40,148][INFO ][wwAP][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Process_05xymvn](Process_05xymvn) start.`,
		`[2021-12-23 11:39:40,148][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultStartEvent(24)]-Task [StartEvent_1](StartEvent_1) start.`,
		`[2021-12-23 11:39:40,148][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultStartEvent(31)]-Task [StartEvent_1](StartEvent_1) finish.(0ms)`,
		`[2021-12-23 11:39:40,149][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(24)]-Task [Gateway_1udd4th](Gateway_1udd4th) start.`,
		`[2021-12-23 11:39:40,149][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(31)]-Task [Gateway_1udd4th](Gateway_1udd4th) finish.(0ms)`,
		`[2021-12-23 11:39:40,149][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(24)]-Task [Gateway_0qw4j58](Gateway_0qw4j58) start.`,
		`[2021-12-23 11:39:40,149][INFO ][wwAP][com.dongkuk.oasis.executors.ParallelGatewayExecutable(71)]-Parallel Sub-process run start. [Gateway_0qw4j58]`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:NPrG][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(24)]-Task [Activity_1msngyv](인사1 프로세스) start.`,
		// `[2021-12-23 11:39:40,150][INFO ][wwAP:NPrG][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Activity_1msngyv](인사1 프로세스) start.`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(24)]-Task [Activity_0nlxb65](인사2 메인 프로세스) start.`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:NPrG][com.dongkuk.oasis.model.activity.JavaServiceTask(24)]-Task [Activity_0o9fhf8](인사1) start.`,
		// `[2021-12-23 11:39:40,150][INFO ][wwAP:Jeme][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Activity_0nlxb65](인사2 메인 프로세스) start.`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:NPrG][com.dongkuk.oasis.executors.PlainJavaServiceTaskExecutable(72)]-Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.event.DefaultStartEvent(24)]-Task [Event_1smi0ye](Event_1smi0ye) start.`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.event.DefaultStartEvent(31)]-Task [Event_1smi0ye](Event_1smi0ye) finish.(0ms)`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(24)]-Task [Gateway_0xdapvd](Gateway_0xdapvd) start.`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(31)]-Task [Gateway_0xdapvd](Gateway_0xdapvd) finish.(0ms)`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(24)]-Task [Gateway_111cvzu](Gateway_111cvzu) start.`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:Jeme][com.dongkuk.oasis.executors.ParallelGatewayExecutable(71)]-Parallel Sub-process run start. [Gateway_111cvzu]`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:NPrG][com.dongkuk.dmes.cr.biz.plan.Hello(10)]-JeongjinKim, hello.`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:NPrG][com.dongkuk.oasis.model.activity.JavaServiceTask(31)]-Task [Activity_0o9fhf8](인사1) finish.(1ms)`,
		`[2021-12-23 11:39:40,152][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(24)]-Task [Activity_0mq11dx](인사2 프로세스) start.`,
		// `[2021-12-23 11:39:40,152][INFO ][wwAP:NPrG][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Activity_1msngyv](인사1 프로세스) finish.(1ms)`,
		// `[2021-12-23 11:39:40,152][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Activity_0mq11dx](인사2 프로세스) start.`,
		`[2021-12-23 11:39:40,152][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.model.activity.JavaServiceTask(24)]-Task [Activity_0wmahw6](인사2) start.`,
		`[2021-12-23 11:39:40,152][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.executors.PlainJavaServiceTaskExecutable(72)]-Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]`,
		`[2021-12-23 11:39:40,152][INFO ][wwAP:NPrG][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(31)]-Task [Activity_1msngyv](인사1 프로세스) finish.(2ms)`,
		`[2021-12-23 11:39:40,153][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(24)]-Task [Activity_1o33vaa](인사3 프로세스) start.`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:QmYu][com.dongkuk.dmes.cr.biz.plan.Hello(10)]-JeongjinKim, hello.`,
		// `[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Activity_1o33vaa](인사3 프로세스) start.`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.model.activity.JavaServiceTask(31)]-Task [Activity_0wmahw6](인사2) finish.(1ms)`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.model.activity.JavaServiceTask(24)]-Task [Activity_0s0rbua](인사3) start.`,
		// `[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Activity_0mq11dx](인사2 프로세스) finish.(2ms)`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.executors.PlainJavaServiceTaskExecutable(72)]-Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(31)]-Task [Activity_0mq11dx](인사2 프로세스) finish.(2ms)`,
		`[2021-12-23 11:39:40,155][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.dmes.cr.biz.plan.Hello(10)]-JeongjinKim, hello.`,
		`[2021-12-23 11:39:40,155][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.model.activity.JavaServiceTask(31)]-Task [Activity_0s0rbua](인사3) finish.(0ms)`,
		// `[2021-12-23 11:39:40,156][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Activity_1o33vaa](인사3 프로세스) finish.(1ms)`,
		// `[2021-12-23 11:39:40,156][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(31)]-Task [Activity_1o33vaa](인사3 프로세스) finish.(1ms)`,
		// `[2021-12-23 11:39:40,156][INFO ][wwAP:Jeme][com.dongkuk.oasis.executors.ParallelGatewayExecutable(91)]-Parallel Sub-process run finish. [Gateway_111cvzu]`,
		`[2021-12-23 11:39:40,156][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(31)]-Task [Gateway_111cvzu](Gateway_111cvzu) finish.(5ms)`,
		`[2021-12-23 11:39:40,157][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.event.DefaultEndEvent(24)]-Task [Event_1035s1u](Event_1035s1u) start.`,
		`[2021-12-23 11:39:40,157][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.event.DefaultEndEvent(31)]-Task [Event_1035s1u](Event_1035s1u) finish.(0ms)`,
		// `[2021-12-23 11:39:40,157][INFO ][wwAP:Jeme][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Activity_0nlxb65](인사2 메인 프로세스) finish.(6ms)`,
		`[2021-12-23 11:39:40,157][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(31)]-Task [Activity_0nlxb65](인사2 메인 프로세스) finish.(6ms)`,
		// `[2021-12-23 11:39:40,157][INFO ][wwAP][com.dongkuk.oasis.executors.ParallelGatewayExecutable(91)]-Parallel Sub-process run finish. [Gateway_0qw4j58]`,
		`[2021-12-23 11:39:40,157][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(31)]-Task [Gateway_0qw4j58](Gateway_0qw4j58) finish.(8ms)`,
		`[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultEndEvent(24)]-Task [Event_0t44im8](Event_0t44im8) start.`,
		`[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultEndEvent(31)]-Task [Event_0t44im8](Event_0t44im8) finish.(0ms)`,
		// `[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Process_05xymvn](Process_05xymvn) finish.(9ms)`,
		`[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.executors.SubServiceCallTaskExecutable(72)]-Sub-service [parallelGreeting] finish.`,
		`[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.model.activity.SubServiceCallTask(31)]-Task [Activity_1obyfx7](parallelGreeting) finish.(14ms)`,
		`[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.model.activity.SqlScriptTask(24)]-Task [Activity_0591vg1](MO 조회) start.`,
		`[2021-12-23 11:39:40,159][DEBUG][wwAP][org.springframework.jdbc.core.JdbcTemplate(711)]-Executing prepared SQL query`,
		`[2021-12-23 11:39:40,159][DEBUG][wwAP][org.springframework.jdbc.core.JdbcTemplate(643)]-Executing prepared SQL statement [select /* sample */mo_no, weight from mo_hea where mo_no = ?]`,
		`[2021-12-23 11:39:40,159][TRACE][wwAP][org.springframework.jdbc.core.StatementCreatorUtils(221)]-Setting SQL statement parameter value: column index 1, parameter value [M000164], value class [java.lang.String], SQL type unknown`,
		`[2021-12-23 11:39:40,164][INFO ][wwAP][com.dongkuk.oasis.model.activity.SqlScriptTask(31)]-Task [Activity_0591vg1](MO 조회) finish.(5ms)`,
		`[2021-12-23 11:39:40,164][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultEndEvent(24)]-Task [Event_19hiwi1](Event_19hiwi1) start.`,
		`[2021-12-23 11:39:40,164][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultEndEvent(31)]-Task [Event_19hiwi1](Event_19hiwi1) finish.(0ms)`,
		// `[2021-12-23 11:39:40,164][INFO ][wwAP][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Process_05xymvn](Process_05xymvn) finish.(21ms)`,
		`[2021-12-23 11:39:40,167][INFO ][wwAP][com.dongkuk.oasis.transaction.SpringTransactionHandler(117)]-Transaction [txBiz] has been committed.`,
		`[2021-12-23 11:39:40,167][INFO ][wwAP][com.dongkuk.oasis.service.SpringServiceStarter(37)]-Service [mo] finish.(34ms)`,
	}

	toLexCh := make(chan logHeader.LogHeader, 100)

	lex := Lexer{
		Server:     "운영계1",
		Chain:      "AAA",
		Key:        "wwAP",
		ReciveChan: toLexCh,
		ExpireTime: 3,
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		lex.Run()
		wg.Done()
	}()

	for _, logString := range logs {
		header := logHeader.GetLogHeader(logString)
		toLexCh <- *header
	}

	wg.Wait()
}

func TestBrokerEmulWithMuiltiInstance(t *testing.T) {
	logs := []string{
		`[2021-12-23 11:39:40,122][DEBUG][wwAP][org.hibernate.SQL(144)]-
		select
			next_val as id_val 
		from
			hibernate_sequence for update
				`,
		`[2021-12-23 11:39:40,125][DEBUG][wwAP][org.hibernate.SQL(144)]-
		update
			hibernate_sequence 
		set
			next_val= ? 
		where
			next_val=?`,
		`[2021-12-23 11:39:40,128][DEBUG][wwAP][org.hibernate.SQL(144)]-
		insert 
		into
			request_log
			(content, id) 
		values
			(?, ?)`,
		`[2021-12-23 11:39:40,129][TRACE][wwAP][org.hibernate.type.descriptor.sql.BasicBinder(64)]-binding parameter [1] as [VARCHAR] - [payload:[null]]`,
		`[2021-12-23 11:39:40,129][TRACE][wwAP][org.hibernate.type.descriptor.sql.BasicBinder(64)]-binding parameter [2] as [BIGINT] - [9]`,
		`[2021-12-23 11:39:40,132][INFO ][wwAP][com.dongkuk.oasis.service.SpringServiceStarter(23)]-Service [mo] start. Request Tag [xXI]`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.transaction.SpringTransactionHandler(82)]-Transaction [txBiz] started.`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Process_05xymvn](Process_05xymvn) start.`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultStartEvent(24)]-Task [StartEvent_1](StartEvent_1) start.`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultStartEvent(31)]-Task [StartEvent_1](StartEvent_1) finish.(0ms)`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultExclusiveGateway(24)]-Task [Gateway_1xtcgie](Gateway_1xtcgie) start.`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultExclusiveGateway(31)]-Task [Gateway_1xtcgie](Gateway_1xtcgie) finish.(0ms)`,
		`[2021-12-23 11:39:40,143][INFO ][wwAP][com.dongkuk.oasis.model.activity.SubServiceCallTask(24)]-Task [Activity_1obyfx7](parallelGreeting) start.`,
		`[2021-12-23 11:39:40,148][INFO ][wwAP][com.dongkuk.oasis.executors.SubServiceCallTaskExecutable(68)]-Sub-service [parallelGreeting] start.`,
		`[2021-12-23 11:39:40,148][INFO ][wwAP][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Process_05xymvn](Process_05xymvn) start.`,
		`[2021-12-23 11:39:40,148][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultStartEvent(24)]-Task [StartEvent_1](StartEvent_1) start.`,
		`[2021-12-23 11:39:40,148][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultStartEvent(31)]-Task [StartEvent_1](StartEvent_1) finish.(0ms)`,
		`[2021-12-23 11:39:40,149][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(24)]-Task [Gateway_1udd4th](Gateway_1udd4th) start.`,
		`[2021-12-23 11:39:40,149][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(31)]-Task [Gateway_1udd4th](Gateway_1udd4th) finish.(0ms)`,
		`[2021-12-23 11:39:40,149][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(24)]-Task [Gateway_0qw4j58](Gateway_0qw4j58) start.`,
		`[2021-12-23 11:39:40,149][INFO ][wwAP][com.dongkuk.oasis.executors.ParallelGatewayExecutable(71)]-Parallel Sub-process run start. [Gateway_0qw4j58]`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:NPrG][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(24)]-Task [Activity_1msngyv](인사1 프로세스) start.`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:NPrG][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Activity_1msngyv](인사1 프로세스) start.`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(24)]-Task [Activity_0nlxb65](인사2 메인 프로세스) start.`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:NPrG][com.dongkuk.oasis.model.activity.JavaServiceTask(24)]-Task [Activity_0o9fhf8](인사1) start.`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:Jeme][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Activity_0nlxb65](인사2 메인 프로세스) start.`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:NPrG][com.dongkuk.oasis.executors.PlainJavaServiceTaskExecutable(72)]-Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.event.DefaultStartEvent(24)]-Task [Event_1smi0ye](Event_1smi0ye) start.`,
		`[2021-12-23 11:39:40,150][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.event.DefaultStartEvent(31)]-Task [Event_1smi0ye](Event_1smi0ye) finish.(0ms)`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(24)]-Task [Gateway_0xdapvd](Gateway_0xdapvd) start.`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(31)]-Task [Gateway_0xdapvd](Gateway_0xdapvd) finish.(0ms)`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(24)]-Task [Gateway_111cvzu](Gateway_111cvzu) start.`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:Jeme][com.dongkuk.oasis.executors.ParallelGatewayExecutable(71)]-Parallel Sub-process run start. [Gateway_111cvzu]`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:NPrG][com.dongkuk.dmes.cr.biz.plan.Hello(10)]-JeongjinKim, hello.`,
		`[2021-12-23 11:39:40,151][INFO ][wwAP:NPrG][com.dongkuk.oasis.model.activity.JavaServiceTask(31)]-Task [Activity_0o9fhf8](인사1) finish.(1ms)`,
		`[2021-12-23 11:39:40,152][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(24)]-Task [Activity_0mq11dx](인사2 프로세스) start.`,
		`[2021-12-23 11:39:40,152][INFO ][wwAP:NPrG][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Activity_1msngyv](인사1 프로세스) finish.(1ms)`,
		`[2021-12-23 11:39:40,152][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Activity_0mq11dx](인사2 프로세스) start.`,
		`[2021-12-23 11:39:40,152][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.model.activity.JavaServiceTask(24)]-Task [Activity_0wmahw6](인사2) start.`,
		`[2021-12-23 11:39:40,152][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.executors.PlainJavaServiceTaskExecutable(72)]-Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]`,
		`[2021-12-23 11:39:40,152][INFO ][wwAP:NPrG][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(31)]-Task [Activity_1msngyv](인사1 프로세스) finish.(2ms)`,
		`[2021-12-23 11:39:40,153][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(24)]-Task [Activity_1o33vaa](인사3 프로세스) start.`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:QmYu][com.dongkuk.dmes.cr.biz.plan.Hello(10)]-JeongjinKim, hello.`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.process.CoreProcessStarter(22)]-Process [Activity_1o33vaa](인사3 프로세스) start.`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.model.activity.JavaServiceTask(31)]-Task [Activity_0wmahw6](인사2) finish.(1ms)`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.model.activity.JavaServiceTask(24)]-Task [Activity_0s0rbua](인사3) start.`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Activity_0mq11dx](인사2 프로세스) finish.(2ms)`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.executors.PlainJavaServiceTaskExecutable(72)]-Invoking class : [com.dongkuk.dmes.cr.biz.plan.Hello], method : [greeting]`,
		`[2021-12-23 11:39:40,154][INFO ][wwAP:Jeme:QmYu][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(31)]-Task [Activity_0mq11dx](인사2 프로세스) finish.(2ms)`,
		`[2021-12-23 11:39:40,155][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.dmes.cr.biz.plan.Hello(10)]-JeongjinKim, hello.`,
		`[2021-12-23 11:39:40,155][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.model.activity.JavaServiceTask(31)]-Task [Activity_0s0rbua](인사3) finish.(0ms)`,
		`[2021-12-23 11:39:40,156][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Activity_1o33vaa](인사3 프로세스) finish.(1ms)`,
		`[2021-12-23 11:39:40,156][INFO ][wwAP:Jeme:D7Ib][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(31)]-Task [Activity_1o33vaa](인사3 프로세스) finish.(1ms)`,
		`[2021-12-23 11:39:40,156][INFO ][wwAP:Jeme][com.dongkuk.oasis.executors.ParallelGatewayExecutable(91)]-Parallel Sub-process run finish. [Gateway_111cvzu]`,
		`[2021-12-23 11:39:40,156][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(31)]-Task [Gateway_111cvzu](Gateway_111cvzu) finish.(5ms)`,
		`[2021-12-23 11:39:40,157][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.event.DefaultEndEvent(24)]-Task [Event_1035s1u](Event_1035s1u) start.`,
		`[2021-12-23 11:39:40,157][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.event.DefaultEndEvent(31)]-Task [Event_1035s1u](Event_1035s1u) finish.(0ms)`,
		`[2021-12-23 11:39:40,157][INFO ][wwAP:Jeme][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Activity_0nlxb65](인사2 메인 프로세스) finish.(6ms)`,
		`[2021-12-23 11:39:40,157][INFO ][wwAP:Jeme][com.dongkuk.oasis.model.activity.DefaultInlineSubProcess(31)]-Task [Activity_0nlxb65](인사2 메인 프로세스) finish.(6ms)`,
		`[2021-12-23 11:39:40,157][INFO ][wwAP][com.dongkuk.oasis.executors.ParallelGatewayExecutable(91)]-Parallel Sub-process run finish. [Gateway_0qw4j58]`,
		`[2021-12-23 11:39:40,157][INFO ][wwAP][com.dongkuk.oasis.model.gateway.DefaultParallelGateway(31)]-Task [Gateway_0qw4j58](Gateway_0qw4j58) finish.(8ms)`,
		`[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultEndEvent(24)]-Task [Event_0t44im8](Event_0t44im8) start.`,
		`[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultEndEvent(31)]-Task [Event_0t44im8](Event_0t44im8) finish.(0ms)`,
		`[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Process_05xymvn](Process_05xymvn) finish.(9ms)`,
		`[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.executors.SubServiceCallTaskExecutable(72)]-Sub-service [parallelGreeting] finish.`,
		`[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.model.activity.SubServiceCallTask(31)]-Task [Activity_1obyfx7](parallelGreeting) finish.(14ms)`,
		`[2021-12-23 11:39:40,158][INFO ][wwAP][com.dongkuk.oasis.model.activity.SqlScriptTask(24)]-Task [Activity_0591vg1](MO 조회) start.`,
		`[2021-12-23 11:39:40,159][DEBUG][wwAP][org.springframework.jdbc.core.JdbcTemplate(711)]-Executing prepared SQL query`,
		`[2021-12-23 11:39:40,159][DEBUG][wwAP][org.springframework.jdbc.core.JdbcTemplate(643)]-Executing prepared SQL statement [select /* sample */mo_no, weight from mo_hea where mo_no = ?]`,
		`[2021-12-23 11:39:40,159][TRACE][wwAP][org.springframework.jdbc.core.StatementCreatorUtils(221)]-Setting SQL statement parameter value: column index 1, parameter value [M000164], value class [java.lang.String], SQL type unknown`,
		`[2021-12-23 11:39:40,164][INFO ][wwAP][com.dongkuk.oasis.model.activity.SqlScriptTask(31)]-Task [Activity_0591vg1](MO 조회) finish.(5ms)`,
		`[2021-12-23 11:39:40,164][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultEndEvent(24)]-Task [Event_19hiwi1](Event_19hiwi1) start.`,
		`[2021-12-23 11:39:40,164][INFO ][wwAP][com.dongkuk.oasis.model.event.DefaultEndEvent(31)]-Task [Event_19hiwi1](Event_19hiwi1) finish.(0ms)`,
		`[2021-12-23 11:39:40,164][INFO ][wwAP][com.dongkuk.oasis.process.CoreProcessStarter(29)]-Process [Process_05xymvn](Process_05xymvn) finish.(21ms)`,
		`[2021-12-23 11:39:40,167][INFO ][wwAP][com.dongkuk.oasis.transaction.SpringTransactionHandler(117)]-Transaction [txBiz] has been committed.`,
		`[2021-12-23 11:39:40,167][INFO ][wwAP][com.dongkuk.oasis.service.SpringServiceStarter(37)]-Service [mo] finish.(34ms)`,
	}

	toLexCh := make(chan logHeader.LogHeader, 100)

	lex := Lexer{
		Server:     "운영계1",
		Chain:      "AAA",
		Key:        "wwAP",
		ReciveChan: toLexCh,
		ExpireTime: 3,
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		lex.Run()
		wg.Done()
	}()

	for _, logString := range logs {
		header := logHeader.GetLogHeader(logString)
		toLexCh <- *header
	}

	wg.Wait()
}
