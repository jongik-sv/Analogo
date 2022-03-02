package mongoRepository

import (
	"fmt"
	"testing"
	"time"
	// "github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	mongo := &Repository{
		URI:            "mongodb://210.1.1.40:27017",
		DataBase:       "test",
		CollectionName: "oasis4",
	}

	doc := `
	{
		"_id":"2021-12-23 11:36:42,001$운영계$12mF$parallelGreeting",
		"property":
		{
		
		"ObjectType": "Request",
		"Server": "운영계",
		"Chain": "AAA",
		"LogKey": "12mF",
		"Services": ["parallelGreeting"],
		"RequestTag": "OhE",
		"Index": ["payload:[null]", "2"],
		"StartTime": "2021-12-23 11:36:42,001",
		"EndTime": "2021-12-23 11:36:42,082",
		"RunTime": "81",
		"_id": "2021-12-23 11:36:42,001$운영계$12mF$parallelGreeting",
		"finish": "finish"
		},
		"child":[
		"Mapped to com.dongkuk.dmes.cr.frm.inbound.ServiceController#service(String, Map, String, String)",
		"Read 'application/octet-stream' to []",
		{
		"property":
		{
		
		"ObjectType": "Query",
		"SQL": "\n    select\n        next_val as id_val \n    from\n        hibernate_sequence for update\n            ",
		"StartTime": "2021-12-23 11:36:42,007",
		"EndTime": "2021-12-23 11:36:42,010",
		"RunTime": "3"
		}
		}
		,
		{
		"property":
		{
		
		"ObjectType": "Query",
		"SQL": "\n    update\n        hibernate_sequence \n    set\n        next_val= ? \n    where\n        next_val=?",
		"StartTime": "2021-12-23 11:36:42,010",
		"EndTime": "2021-12-23 11:36:42,016",
		"RunTime": "6"
		}
		}
		,
		{
		"property":
		{
		
		"ObjectType": "Query",
		"SQL": "\n    insert \n    into\n        request_log\n        (content, id) \n    values\n        (?, ?)",
		"Param": ["payload:[null]", "2"],
		"StartTime": "2021-12-23 11:36:42,019",
		"EndTime": "2021-12-23 11:36:42,021",
		"RunTime": "2"
		}
		}
		,
		{
		"property":
		{
		
		"ObjectType": "Service",
		"Name": "parallelGreeting",
		"StartTime": "2021-12-23 11:36:42,025",
		"EndTime": "2021-12-23 11:36:42,080",
		"RunTime": "54"
		},
		"child":[
		"Transaction [txBiz] started.",
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "StartEvent_1",
		"TaskID": "StartEvent_1",
		"LogKey": "12mF",
		"StartTime": "2021-12-23 11:36:42,055",
		"EndTime": "2021-12-23 11:36:42,055",
		"RunTime": "0"
		}}
		,
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "Gateway_1udd4th",
		"TaskID": "Gateway_1udd4th",
		"LogKey": "12mF",
		"StartTime": "2021-12-23 11:36:42,056",
		"EndTime": "2021-12-23 11:36:42,058",
		"RunTime": "2"
		}}
		,
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "Gateway_0qw4j58",
		"TaskID": "Gateway_0qw4j58",
		"ProcessType": "Parallel Sub-process",
		"LogKey": "12mF",
		"StartTime": "2021-12-23 11:36:42,058",
		"EndTime": "2021-12-23 11:36:42,073",
		"RunTime": "14"
		},
		"child":[
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "인사2 메인 프로세스",
		"TaskID": "Activity_0nlxb65",
		"LogKey": "12mF:fUsu",
		"StartTime": "2021-12-23 11:36:42,060",
		"EndTime": "2021-12-23 11:36:42,072",
		"RunTime": "11"
		},
		"child":[
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "Event_1smi0ye",
		"TaskID": "Event_1smi0ye",
		"LogKey": "12mF:fUsu",
		"StartTime": "2021-12-23 11:36:42,067",
		"EndTime": "2021-12-23 11:36:42,067",
		"RunTime": "0"
		}}
		,
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "Gateway_0xdapvd",
		"TaskID": "Gateway_0xdapvd",
		"LogKey": "12mF:fUsu",
		"StartTime": "2021-12-23 11:36:42,067",
		"EndTime": "2021-12-23 11:36:42,067",
		"RunTime": "0"
		}}
		,
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "Gateway_111cvzu",
		"TaskID": "Gateway_111cvzu",
		"ProcessType": "Parallel Sub-process",
		"LogKey": "12mF:fUsu",
		"StartTime": "2021-12-23 11:36:42,067",
		"EndTime": "2021-12-23 11:36:42,072",
		"RunTime": "4"
		},
		"child":[
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "인사2 프로세스",
		"TaskID": "Activity_0mq11dx",
		"LogKey": "12mF:fUsu:G0PH",
		"StartTime": "2021-12-23 11:36:42,070",
		"EndTime": "2021-12-23 11:36:42,071",
		"RunTime": "1"
		},
		"child":[
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "인사2",
		"TaskID": "Activity_0wmahw6",
		"LogKey": "12mF:fUsu:G0PH",
		"Class": "com.dongkuk.dmes.cr.biz.plan.Hello",
		"Method": "greeting",
		"StartTime": "2021-12-23 11:36:42,070",
		"EndTime": "2021-12-23 11:36:42,071",
		"RunTime": "0"
		},
		"child":[
		"Parameter name binding of name",
		"kjj, hello."]
		}
		]
		}
		,
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "인사3 프로세스",
		"TaskID": "Activity_1o33vaa",
		"LogKey": "12mF:fUsu:S65g",
		"StartTime": "2021-12-23 11:36:42,070",
		"EndTime": "2021-12-23 11:36:42,071",
		"RunTime": "1"
		},
		"child":[
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "인사3",
		"TaskID": "Activity_0s0rbua",
		"LogKey": "12mF:fUsu:S65g",
		"Class": "com.dongkuk.dmes.cr.biz.plan.Hello",
		"Method": "greeting",
		"StartTime": "2021-12-23 11:36:42,070",
		"EndTime": "2021-12-23 11:36:42,071",
		"RunTime": "0"
		},
		"child":[
		"Parameter name binding of name",
		"kjj, hello."]
		}
		]
		}
		]
		}
		,
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "Event_1035s1u",
		"TaskID": "Event_1035s1u",
		"LogKey": "12mF:fUsu",
		"StartTime": "2021-12-23 11:36:42,072",
		"EndTime": "2021-12-23 11:36:42,072",
		"RunTime": "0"
		}}
		]
		}
		,
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "인사1 프로세스",
		"TaskID": "Activity_1msngyv",
		"LogKey": "12mF:7ES6",
		"StartTime": "2021-12-23 11:36:42,060",
		"EndTime": "2021-12-23 11:36:42,066",
		"RunTime": "5"
		},
		"child":[
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "인사1",
		"TaskID": "Activity_0o9fhf8",
		"LogKey": "12mF:7ES6",
		"Class": "com.dongkuk.dmes.cr.biz.plan.Hello",
		"Method": "greeting",
		"StartTime": "2021-12-23 11:36:42,064",
		"EndTime": "2021-12-23 11:36:42,065",
		"RunTime": "0"
		},
		"child":[
		"Parameter name binding of name",
		"kjj, hello."]
		}
		]
		}
		]
		}
		,
		{
		"property":
		{
		
		"ObjectType": "Task",
		"Name": "Event_0t44im8",
		"TaskID": "Event_0t44im8",
		"LogKey": "12mF",
		"StartTime": "2021-12-23 11:36:42,073",
		"EndTime": "2021-12-23 11:36:42,073",
		"RunTime": "0"
		}}
		,
		"Transaction [txBiz] has been committed."]
		}
		,
		"Using 'application/json', given [*/*] and supported [application/json, application/*+json, application/json, application/*+json]",
		"Writing [{}]",
		"Completed 200 OK"]
		}
		
	`
	mongo.Connect()
	mongo.Insert(doc)
}

func TestDelete(t *testing.T) {
	mongo := &Repository{
		URI:            "mongodb://210.1.1.40:27017",
		DataBase:       "test",
		CollectionName: "oasis4",
	}

	now := time.Now().Format("2006-01-02 15:04:05.000000")
	doc := fmt.Sprintf("{\"_id\":\"%s\", \"insertTime\":\"%s\"}", "aaaa", now)

	mongo.Connect()
	mongo.Delete("aaaa")
	// time.Sleep(time.Second * 10)
	mongo.Insert(doc)
	// mongo.Delete("aaaa")
	// mongo.Insert(doc)
	// mongo.Delete("aaaa")
	// mongo.Insert(doc)
	// mongo.Delete("aaaa")
	// mongo.Insert(doc)
	// mongo.Delete("aaaa")
}

/*
func TestUpsert(t *testing.T) {
	// 칼럼을 선택해서 upsert 한다. --> 로그에선 사용 안함
	mongoRepository := &Repository{
		URI:            "mongodb://210.1.1.40:27017",
		DataBase:       "test",
		CollectionName: "oasis4",
	}

	mongoRepository.Connect()

	coll := *mongoRepository.collection
	id := "aaaa"

	// Find the document for which the _id field matches id and set the email to
	// "newemail@example.com".
	// Specify the Upsert option to insert a new document if a document matching
	// the filter isn't found.
	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"email1", "newemail@example.com1111"}}}}
	var updatedDocument bson.M
	err := coll.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
		opts,
	).Decode(&updatedDocument)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}
	fmt.Printf("updated document %v", updatedDocument)

}
*/
