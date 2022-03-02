package mongoRepository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	URI            string
	DataBase       string
	CollectionName string
	client         *mongo.Client
	collection     *mongo.Collection
}

func (r *Repository) Connect() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(r.URI))
	if err != nil {
		panic(err)
	}
	r.client = client
	/*
		defer func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()
	*/
	collection := client.Database(r.DataBase).Collection(r.CollectionName)
	r.collection = collection
}

func (r *Repository) Insert(docString string) {
	var doc interface{}
	err := bson.UnmarshalExtJSON([]byte(docString), true, &doc)
	if err != nil {
		fmt.Println(docString)
		panic(err)
	}

	result, err := r.collection.InsertOne(context.TODO(), doc)
	fmt.Println("I : ", result.InsertedID)
	if err != nil {
		fmt.Println(docString)
		fmt.Printf("error: %v\n", err)
		// fmt.Println("error: %v", err)
		// panic(err)
	}

	// id := doc["_id"]
	// filter := bson.D{{"_id", ""}}
}

func (r *Repository) Delete(id string) {

	projection := bson.D{
		{"_id", 1},
		{"insertTime", 1},
	}

	opts := options.FindOneAndDelete().SetProjection(projection)

	var deletedDocument bson.M
	err := r.collection.FindOneAndDelete(context.TODO(), bson.D{{"_id", id}}, opts).Decode(&deletedDocument)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}
	fmt.Printf("D :  %v,  %v\n", deletedDocument, err)
}

/*
	var result bson.M
	r.collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&result)
	fmt.Printf("%v\n", result)
	res, err := r.collection.DeleteMany(context.TODO(), bson.D{{"_id", id}})
	fmt.Println("D : ", id, ", ", res.DeletedCount)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		log.Fatal(err)
	}
*/
