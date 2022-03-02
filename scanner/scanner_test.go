package scanner

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListern(t *testing.T) {
	assert := assert.New(t)

	adapterChan := make(chan string, 100)
	brokerChan := make(chan string, 100)
	strSlice := []string{
		"[1] aaaaaa",
		"bbbbbb",
		"[1] cccccc",
		"ddddddd",
		"[1] efadfsa",
		"EOF",
	}

	go Listen(adapterChan, brokerChan)

	// scanner 역할
	go func() {

		for _, str := range strSlice {
			adapterChan <- str
			// sendLog(str)
		}

	}()

	i := 0
	go func() {
		defer close(brokerChan)
		for str := range brokerChan {
			fmt.Println(" str : ", str, "\n cnt : ", i, "\n time : ", time.Now())
			i++
		}

	}()

	time.Sleep(time.Second * 3)
	assert.Equal(3, i)
}

func TestSendLog(t *testing.T) {
	assert := assert.New(t)
	strSlice := []string{
		"[1] aaaaaa",
		"[1] ddddddd",
		"[1] efadfsa",
	}
	brokerChan := make(chan string, 100)
	go func() {
		for _, str := range strSlice {
			sendLog(str, brokerChan)
		}

	}()

	i := 0
	go func() {
		defer close(brokerChan)
		for str := range brokerChan {
			fmt.Println("str : ", str, " cnt : ", i)
			i++
		}

	}()

	time.Sleep(time.Second * 3)
	assert.Equal(3, i)

}
func TestIsStart(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, isStart("[****************]"))
	assert.Equal(false, isStart("****************]"))

}

func TestTrimReturnCharacter(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("aaaa", trimReturnCharacter("aaaa\n"))
	assert.Equal("aa\naa", trimReturnCharacter("aa\naa\n"))
	assert.Equal("\naaab", trimReturnCharacter("\naaab"))
}

func TestChannelStringSize(t *testing.T) {

	log := []string{
		"[2021-12-23 11:35:50,019][DEBUG][][org.hibernate.internal.SessionFactoryImpl(268)]-Session factory constructed with filter configurations : {}",
		"[2021-12-23 11:35:50,019][DEBUG][][org.hibernate.internal.SessionFactoryImpl(269)]-Instantiating session factory with properties: {gopherProxySet=false, awt.toolkit=sun.lwawt.macosx.LWCToolkit, hibernate.format_sql=true, java.specification.version=11, sun.cpu.isalist=, hibernate.connection.handling_mode=DELAYED_ACQUISITION_AND_HOLD, sun.jnu.encoding=UTF-8, hibernate.dialect=org.hibernate.dialect.MySQL5Dialect, java.class.path=/Users/cothe/IdeaProjects/dmes-cr-pmf/dmes-cr-pmf-web-jar/build/classes/java/main:/Users/cothe/IdeaProjects/dmes-cr-pmf/dmes-cr-pmf-web-jar/build/resources/main:/Users/cothe/IdeaProjects/dmes-cr-pmf/dmes-cr-pmf-core/build/classes/java/main:/Users/cothe/IdeaProjects/dmes-cr-pmf/dmes-cr-pmf-core/build/resources/main:/Users/cothe/.m2/repository/com/dongkuk/dmes/cr/dmes-cr-frm-config/0.0.1/dmes-cr-frm-config-0.0.1.jar:/Users/cothe/.m2/repository/org/apache/logging/log4j/log4j-slf4j-impl/2.17.0/log4j-slf4j-impl-2.17.0.jar:/Users/cothe/.m2/repository/org/apache/logging/log4j/log4j-core/2.17.0/log4j-core-2.17.0.jar:/Users/cothe/.m2/repository/org/apache/logging/log4j/log4j-api/2.17.0/log4j-api-2.17.0.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework.boot/spring-boot-starter-web/2.6.1/145ac0cfb81982608ef0d19e32699c0eeeb3c2ab/spring-boot-starter-web-2.6.1.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework.boot/spring-boot-starter-thymeleaf/2.6.1/7eb008866400aa6eeafd6ef73c92e975ba35811b/spring-boot-starter-thymeleaf-2.6.1.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.slf4j/slf4j-api/1.7.32/cdcff33940d9f2de763bc41ea05a0be5941176c3/slf4j-api-1.7.32.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework.boot/spring-boot-starter-json/2.6.1/b76ad5d869508cc78aaf00bb7ab92e8a472b26ed/spring-boot-starter-json-2.6.1.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework.boot/spring-boot-starter/2.6.1/39d77929829c82b102248630047abf0f69b96a7b/spring-boot-starter-2.6.1.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework.boot/spring-boot-starter-tomcat/2.6.1/6a0744acdd17b860e2ba3edb95c3c7cec36e6b58/spring-boot-starter-tomcat-2.6.1.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework/spring-webmvc/5.3.13/cea31c85fa84dbd9f8df14a3ca62ab57c25cabe4/spring-webmvc-5.3.13.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework/spring-web/5.3.13/66d95a5d2d436961b4cae036723f4c7a764fc14c/spring-web-5.3.13.jar:/Users/cothe/.m2/repository/org/thymeleaf/thymeleaf-spring5/3.0.12.RELEASE/thymeleaf-spring5-3.0.12.RELEASE.jar:/Users/cothe/.m2/repository/org/thymeleaf/extras/thymeleaf-extras-java8time/3.0.4.RELEASE/thymeleaf-extras-java8time-3.0.4.RELEASE.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/com.fasterxml.jackson.datatype/jackson-datatype-jsr310/2.13.0/4c143877fc733befe6189151c8b95d84acd06941/jackson-datatype-jsr310-2.13.0.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/com.fasterxml.jackson.module/jackson-module-parameter-names/2.13.0/62a5e3b6cc5aacc0ff2354053f3cc2a7821dfa9/jackson-module-parameter-names-2.13.0.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/com.fasterxml.jackson.datatype/jackson-datatype-jdk8/2.13.0/6d82a502e03f61d6dc6e47e5f7d6168141419c92/jackson-datatype-jdk8-2.13.0.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/com.fasterxml.jackson.core/jackson-databind/2.13.0/889672a1721d6d85b2834fcd29d3fda92c8c8891/jackson-databind-2.13.0.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework.boot/spring-boot-autoconfigure/2.6.1/9ac07afd64da0cce435792ba1328c93edcfbb2fb/spring-boot-autoconfigure-2.6.1.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework.boot/spring-boot/2.6.1/f670cee55752c1f1b304508e18bafd000e543174/spring-boot-2.6.1.jar:/Users/cothe/.m2/repository/jakarta/annotation/jakarta.annotation-api/1.3.5/jakarta.annotation-api-1.3.5.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework/spring-core/5.3.13/d2a6c3372dd337e08144f9f49f386b8ec7a8080d/spring-core-5.3.13.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.yaml/snakeyaml/1.29/6d0cdafb2010f1297e574656551d7145240f6e25/snakeyaml-1.29.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.apache.tomcat.embed/tomcat-embed-websocket/9.0.55/4e6dc3646d00887497bd465bc4a63bfb0a7b10ab/tomcat-embed-websocket-9.0.55.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.apache.tomcat.embed/tomcat-embed-core/9.0.55/6ab68425d34f35e93cf97e1950c2c710161d8ce1/tomcat-embed-core-9.0.55.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.apache.tomcat.embed/tomcat-embed-el/9.0.55/d8b69643d1566712cf849a7e8e95c917f8aed1d3/tomcat-embed-el-9.0.55.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework/spring-context/5.3.13/e328db1c30ffe1c58328e4ab42cd3855a5307469/spring-context-5.3.13.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework/spring-aop/5.3.13/e0fddf47af3fbbec69a403c058c23505612ca329/spring-aop-5.3.13.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework/spring-beans/5.3.13/1d90c96b287253ec371260c35fbbea719c24bad6/spring-beans-5.3.13.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework/spring-expression/5.3.13/8f7448f4fb296a92855fd0afea3375ce41061e84/spring-expression-5.3.13.jar:/Users/cothe/.m2/repository/org/thymeleaf/thymeleaf/3.0.12.RELEASE/thymeleaf-3.0.12.RELEASE.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/com.fasterxml.jackson.core/jackson-annotations/2.13.0/15be89db6609bd6fda3dc309bacf0318a312c03f/jackson-annotations-2.13.0.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/com.fasterxml.jackson.core/jackson-core/2.13.0/e957ec5442966e69cef543927bdc80e5426968bb/jackson-core-2.13.0.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework/spring-jcl/5.3.13/3aa15be194887fbda3912ecbb4ab6ec8dfdebdb0/spring-jcl-5.3.13.jar:/Users/cothe/.m2/repository/org/attoparser/attoparser/2.0.5.RELEASE/attoparser-2.0.5.RELEASE.jar:/Users/cothe/.m2/repository/org/unbescape/unbescape/1.1.6.RELEASE/unbescape-1.1.6.RELEASE.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework.data/spring-data-jpa/2.6.0/bd08ea8db76c7c82397307dda2e253180c31b7ec/spring-data-jpa-2.6.0.jar:/Users/cothe/.m2/repository/javax/persistence/javax.persistence-api/2.2/javax.persistence-api-2.2.jar:/Users/cothe/.m2/repository/org/mybatis/mybatis/3.5.6/mybatis-3.5.6.jar:/Users/cothe/.m2/repository/com/dongkuk/oasis-core/4.2.0/oasis-core-4.2.0.jar:/Users/cothe/.m2/repository/com/zaxxer/HikariCP/4.0.3/HikariCP-4.0.3.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.hibernate/hibernate-entitymanager/5.6.1.Final/f08814f6b4a5dcaafcc63dee1ab73fe5708ebf5c/hibernate-entitymanager-5.6.1.Final.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework/spring-orm/5.3.13/cfcd1ea05a881200ddf7c67a1127a0f7c2efcf06/spring-orm-5.3.13.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/javax.servlet/javax.servlet-api/4.0.1/a27082684a2ff0bf397666c3943496c44541d1ca/javax.servlet-api-4.0.1.jar:/Users/cothe/.m2/repository/org/mybatis/mybatis-spring/2.0.6/mybatis-spring-2.0.6.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/mysql/mysql-connector-java/8.0.27/f1da9f10a3de6348725a413304aab6d0aa04f923/mysql-connector-java-8.0.27.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework.data/spring-data-commons/2.6.0/5a9afaa6e0a4cd74183a794f467c9b4a546b4cbe/spring-data-commons-2.6.0.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework/spring-tx/5.3.13/4b7346f190d6a6f4b983bc23d1f0145c2ff2dbb7/spring-tx-5.3.13.jar:/Users/cothe/.m2/repository/io/github/thecodinglog/method-invoker/0.2.0/method-invoker-0.2.0.jar:/Users/cothe/.m2/repository/org/jdom/jdom2/2.0.6/jdom2-2.0.6.jar:/Users/cothe/.m2/repository/org/jboss/jboss-vfs/3.1.0.Final/jboss-vfs-3.1.0.Final.jar:/Users/cothe/.m2/repository/com/google/code/findbugs/annotations/3.0.1/annotations-3.0.1.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.hibernate/hibernate-core/5.6.1.Final/e587b28eef7a5dd7c742d5dd20e16c9c38b2e34d/hibernate-core-5.6.1.Final.jar:/Users/cothe/.m2/repository/org/hibernate/common/hibernate-commons-annotations/5.1.2.Final/hibernate-commons-annotations-5.1.2.Final.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.jboss.logging/jboss-logging/3.4.2.Final/e517b8a93dd9962ed5481345e4d262fdd47c4217/jboss-logging-3.4.2.Final.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/net.bytebuddy/byte-buddy/1.11.22/8b4c7fa5562a09da1c2a9ab0873cb51f5034d83f/byte-buddy-1.11.22.jar:/Users/cothe/.m2/repository/org/jboss/spec/javax/transaction/jboss-transaction-api_1.2_spec/1.1.1.Final/jboss-transaction-api_1.2_spec-1.1.1.Final.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.springframework/spring-jdbc/5.3.13/70d775617131bfd87370f590c94c55319f7964ff/spring-jdbc-5.3.13.jar:/Users/cothe/.m2/repository/net/jcip/jcip-annotations/1.0/jcip-annotations-1.0.jar:/Users/cothe/.m2/repository/com/google/code/findbugs/jsr305/3.0.1/jsr305-3.0.1.jar:/Users/cothe/.m2/repository/antlr/antlr/2.7.7/antlr-2.7.7.jar:/Users/cothe/.m2/repository/org/jboss/jandex/2.2.3.Final/jandex-2.2.3.Final.jar:/Users/cothe/.m2/repository/com/fasterxml/classmate/1.5.1/classmate-1.5.1.jar:/Users/cothe/.m2/repository/javax/xml/bind/jaxb-api/2.3.1/jaxb-api-2.3.1.jar:/Users/cothe/.m2/repository/javax/activation/javax.activation-api/1.2.0/javax.activation-api-1.2.0.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.glassfish.jaxb/jaxb-runtime/2.3.5/a169a961a2bb9ac69517ec1005e451becf5cdfab/jaxb-runtime-2.3.5.jar:/Users/cothe/.m2/repository/jakarta/xml/bind/jakarta.xml.bind-api/2.3.3/jakarta.xml.bind-api-2.3.3.jar:/Users/cothe/.gradle/caches/modules-2/files-2.1/org.glassfish.jaxb/txw2/2.3.5/ec8930fa62e7b1758b1664d135f50c7abe86a4a3/txw2-2.3.5.jar:/Users/cothe/.m2/repository/com/sun/istack/istack-commons-runtime/3.0.12/istack-commons-runtime-3.0.12.jar:/Users/cothe/.m2/repository/com/sun/activation/jakarta.activation/1.2.2/jakarta.activation-1.2.2.jar, java.vm.vendor=AdoptOpenJDK, sun.arch.data.model=64, java.vendor.url=https://adoptopenjdk.net/, catalina.useNaming=false, user.timezone=Asia/Seoul, jakarta.persistence.sharedCache.mode=UNSPECIFIED, os.name=Mac OS X, java.vm.specification.version=11, javax.persistence.validation.mode=AUTO, jakarta.persistence.nonJtaDataSource=HikariDataSource (HikariPool-1), sun.java.launcher=SUN_STANDARD, user.country=KR, sun.boot.library.path=/Library/Java/JavaVirtualMachines/adoptopenjdk-11.jdk/Contents/Home/lib, sun.java.command=com.dongkuk.dmes.cr.pmf.DmesCrPmfWebJarApplication, spring.application.admin.enabled=true, javax.persistence.nonJtaDataSource=HikariDataSource (HikariPool-1), http.nonProxyHosts=local|*.local|169.254/16|*.169.254/16, com.sun.management.jmxremote=, javax.persistence.sharedCache.mode=UNSPECIFIED, jdk.debug=release, spring.liveBeansView.mbeanDomain=, sun.cpu.endian=little, user.home=/Users/cothe, user.language=en, java.specification.vendor=Oracle Corporation, java.version.date=2021-04-20, java.home=/Library/Java/JavaVirtualMachines/adoptopenjdk-11.jdk/Contents/Home, spring.profiles.active=dev, file.separator=/, spring.output.ansi.enabled=always, java.vm.compressedOopsMode=Zero based, line.separator=",
		", hibernate.persistenceUnitName=biz, java.specification.name=Java Platform API Specification, java.vm.specification.vendor=Oracle Corporation, FILE_LOG_CHARSET=UTF-8, hibernate.transaction.coordinator_class=class org.hibernate.resource.transaction.backend.jdbc.internal.JdbcResourceLocalTransactionCoordinatorBuilderImpl, java.awt.graphicsenv=sun.awt.CGraphicsEnvironment, java.awt.headless=true, jakarta.persistence.validation.mode=AUTO, hibernate.hbm2ddl.auto=create, sun.management.compiler=HotSpot 64-Bit Tiered Compilers, ftp.nonProxyHosts=local|*.local|169.254/16|*.169.254/16, dmes.cr.logfile.path=/Users/cothe/logs, java.runtime.version=11.0.11+9, user.name=cothe, spring.jmx.enabled=true, path.separator=:, os.version=10.16, java.runtime.name=OpenJDK Runtime Environment, file.encoding=UTF-8, hibernate.ejb.persistenceUnitName=biz, hibernate.user_sql_comments=true, spring.beaninfo.ignore=true, java.vm.name=OpenJDK 64-Bit Server VM, hibernate.show_sql=false, java.vendor.version=AdoptOpenJDK-11.0.11+9, java.vendor.url.bug=https://github.com/AdoptOpenJDK/openjdk-support/issues, java.io.tmpdir=/var/folders/sy/7273b__s20z63__cxv7cxtcr0000gn/T/, catalina.home=/private/var/folders/sy/7273b__s20z63__cxv7cxtcr0000gn/T/tomcat.8080.12899743477260524331, dmes.cr.module.name=pmf, java.version=11.0.11, hibernate.physical_naming_strategy=com.dongkuk.dmes.cr.frm.SnakePhysicalNamingStrategy, user.dir=/Users/cothe/IdeaProjects/dmes-cr-pmf, os.arch=x86_64, java.vm.specification.name=Java Virtual Machine Specification, PID=49311, java.awt.printerjob=sun.lwawt.macosx.CPrinterJob, sun.os.patch.level=unknown, CONSOLE_LOG_CHARSET=UTF-8, catalina.base=/private/var/folders/sy/7273b__s20z63__cxv7cxtcr0000gn/T/tomcat.8080.12899743477260524331, hibernate.boot.CfgXmlAccessService.key=org.hibernate.boot.registry.StandardServiceRegistryBuilder$1@50672905, java.library.path=/Users/cothe/Library/Java/Extensions:/Library/Java/Extensions:/Network/Library/Java/Extensions:/System/Library/Java/Extensions:/usr/lib/java:., java.vendor=AdoptOpenJDK, java.vm.info=mixed mode, java.vm.version=11.0.11+9, hibernate.bytecode.use_reflection_optimizer=false, java.rmi.server.randomIDs=true, sun.io.unicode.encoding=UnicodeBig, hibernate.connection.datasource=HikariDataSource (HikariPool-1), socksNonProxyHosts=local|*.local|169.254/16|*.169.254/16, java.class.version=55.0}",
		"[2021-12-23 11:35:50,037][DEBUG][][org.hibernate.secure.spi.JaccIntegrator(84)]-Skipping JACC integration as it was not enabled",
	}

	adaterChan := make(chan string, 1000)
	brokerChan := make(chan string, 1000)

	go Listen(adaterChan, brokerChan)

	go func() {
		for i := 0; i < 100; i++ {
			for _, line := range log {
				adaterChan <- line
				// fmt.Println(line)
			}
		}
		adaterChan <- "<<<<LOG EOF>>>>"
		close(adaterChan)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		i := 0

		for line := range brokerChan {
			if line == "<<<<LOG EOF>>>>" {
				wg.Done()
				return
			}
			fmt.Println(i, line)
			fmt.Printf("\n\n\n")
			i++
		}

	}()
	wg.Wait()
}
