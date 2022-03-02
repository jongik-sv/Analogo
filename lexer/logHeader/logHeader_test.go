package logHeader

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPretoken(t *testing.T) {
	log := "[2021-12-13 11:49:21,162][DEBUG][KFH7:dsdsd:dse34][org.springframework.web.servlet.DispatcherServlet]-Failed to complete request: java.lang.RuntimeException: my exception"

	// logHeader := &LogHeader{TotalString: str}

	preToken := GetLogHeader(log)

	fmt.Println(TIMESTAMP)
	fmt.Println(preToken.Timestamp)
	fmt.Println(preToken.Level)
	fmt.Println(preToken.Key)
	fmt.Println(preToken.Keys)
	fmt.Println(preToken.NDC)
	fmt.Println(preToken.Text)
	fmt.Println(preToken.TotalString)
	assert.Equal(t, TIMESTAMP, 0)
	assert.Equal(t, preToken.Timestamp, "2021-12-13 11:49:21,162")
	assert.Equal(t, preToken.Level, "DEBUG")
	assert.Equal(t, preToken.Key, "KFH7:dsdsd:dse34")
	assert.Equal(t, preToken.Keys, []string{"KFH7", "dsdsd", "dse34"})
	assert.Equal(t, preToken.NDC, "org.springframework.web.servlet.DispatcherServlet")
	assert.Equal(t, preToken.Text, "Failed to complete request: java.lang.RuntimeException: my exception")

}

func TestPrint(t *testing.T) {
	log := "[2021-12-13 11:49:21,162][DEBUG][KFH7:dsdsd:dse34][org.springframework.web.servlet.DispatcherServlet]-Failed to complete request: java.lang.RuntimeException: my exception"
	header := GetLogHeader(log)
	header.Print()
}

func TestGetToken(t *testing.T) {
	log := `[2021-12-23 11:35:49,044][DEBUG][][org.hibernate.jpa.internal.util.LogHelper(102)]-PersistenceUnitInfo [
		name: biz
		persistence provider classname: null
		classloader: jdk.internal.loader.ClassLoaders$AppClassLoader@2c13da15
		excludeUnlistedClasses: true
		JTA datasource: null
		Non JTA datasource: HikariDataSource (null)
		Transaction type: RESOURCE_LOCAL
		PU root URL: file:/Users/cothe/IdeaProjects/dmes-cr-pmf/dmes-cr-pmf-web-jar/build/classes/java/main/
		Shared Cache Mode: UNSPECIFIED
		Validation Mode: AUTO
		Jar files URLs []
		Managed classes names [
			com.dongkuk.dmes.cr.biz.plan.mo.MoHeaEntity]
		Mapping files names []
		Properties []`

	header := GetLogHeader(log)
	header.Print()

}

func TestGetData(t *testing.T) {
	log := "[2021-12-13 11:49:21,162][DEBUG][KFH7:dsdsd:dse34][org.springframework.web.servlet.DispatcherServlet]-Failed to complete request: java.lang.RuntimeException: my exception"
	h := GetLogHeader(log)
	// Server      string
	// Chain       string
	// Timestamp   string
	// Level       string
	// Key         string
	// Keys        []string
	// NDC         string // Nested Diagnostic Contexts
	// Text        string
	// TotalString string
	// OK          bool

	assert.Equal(t, "2021-12-13 11:49:21,162", h.GetData("Timestamp"))
	assert.Equal(t, "DEBUG", h.GetData("Level"))
	assert.Equal(t, "KFH7:dsdsd:dse34", h.GetData("Key"))
	assert.Equal(t, "org.springframework.web.servlet.DispatcherServlet", h.GetData("NDC"))
	assert.Equal(t, "Failed to complete request: java.lang.RuntimeException: my exception", h.GetData("Text"))
}
