package parser

// type PropertyOrder struct {
// 	Order map[string][]string
// }

// func (o *PropertyOrder) InitOrder() {

var propertyOrder = map[string][]string{
	"Request": []string{
		"ObjectType",
		"Server",
		"Chain",
		"LogKey",
		"Services",
		"RequestTag",
		"Address ",
		"Parameters ",
		"Index",
		"Exception",
		"StartTime",
		"EndTime",
		"RunTime",
		"_id",
		"finish",
	},
	"Service": []string{
		"ObjectType",
		"Name",
		"StartTime",
		"EndTime",
		"RunTime"},
	"Task": []string{
		"ObjectType",
		"Name",
		"TaskID",
		"ProcessType",
		"LogKey",
		"Class",
		"Method",
		"hasQuery",
		"StartTime",
		"EndTime",
		"RunTime"},
	"Query": []string{
		"ObjectType",
		"SQL",
		"Param",
		"StartTime",
		"EndTime",
		"RunTime"},
}

// }

// func (o *PropertyOrder) GetObjectOrder(objectType string) []string {
// 	return o.Order[objectType]
// }
