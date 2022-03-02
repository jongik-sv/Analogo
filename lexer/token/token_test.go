package token

// func TestGetToken(t *testing.T) {

// 	header := logHeader.LogHeader{
// 		Server:      "운영계1",
// 		Chain:       "C10",
// 		Timestamp:   "2021-12-13 11:49:21,162",
// 		Level:       "DEBUG",
// 		Key:         "KFH7:dsdsd:dse34",
// 		Keys:        []string{"KFH7", "dsdsd", "dse34"},
// 		NDC:         "org.springframework.web.servlet.DispatcherServlet",
// 		Text:        "Failed to complete request: java.lang.RuntimeException: my exception",
// 		TotalString: "[2021-12-13 11:49:21,162][DEBUG][KFH7:dsdsd:dse34][org.springframework.web.servlet.DispatcherServlet]-Failed to complete request: java.lang.RuntimeException: my exception",
// 	}

// 	token := GetToken(header)

// 	assert.NotEqual(t, token.TokenType, NONE_KEY)

// 	header = logHeader.LogHeader{
// 		Server:      "운영계1",
// 		Chain:       "C10",
// 		Timestamp:   "2021-12-13 11:49:21,162",
// 		Level:       "DEBUG",
// 		Key:         "",
// 		Keys:        []string{},
// 		NDC:         "org.springframework.web.servlet.DispatcherServlet",
// 		Text:        "Failed to complete request: java.lang.RuntimeException: my exception",
// 		TotalString: "[2021-12-13 11:49:21,162][DEBUG][][org.springframework.web.servlet.DispatcherServlet]-Failed to complete request: java.lang.RuntimeException: my exception",
// 	}
// 	token = GetToken(header)

// 	assert.Equal(t, token.TokenType, NONE_KEY)

// }
