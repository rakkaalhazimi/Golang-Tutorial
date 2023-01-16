package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Waifu struct {
	Name string `bson:name`
}

// MongoDB
func getMongoClient(uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	return client, err
}

func listMongoDatabase(client *mongo.Client) {
	databases, _ := client.ListDatabaseNames(context.TODO(), bson.D{})

	for _, database := range databases {
		fmt.Println(database)
	}
}

func listMongoCollection(db *mongo.Database) {
	collections, _ := db.ListCollectionNames(context.TODO(), bson.D{})

	for _, collection := range collections {
		fmt.Println(collection)
	}
}

// DynamoDB
func getDynamoClient() *dynamodb.DynamoDB {
	sess, err := session.NewSessionWithOptions(session.Options{

		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String("ap-southeast-3"),
		},

		// Force enable Shared Config support
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		log.Fatal(err)
	}

	svc := dynamodb.New(sess)
	return svc
}

func listDynamoTable(svc *dynamodb.DynamoDB) {

	input := &dynamodb.ListTablesInput{}
	for {
		// Get the list of tables
		result, err := svc.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		for _, n := range result.TableNames {
			fmt.Println(*n)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}
	}
}

func getDynamoItem(svc *dynamodb.DynamoDB, tableName string, key map[string]*dynamodb.AttributeValue) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	var item map[string]interface{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
}

func main() {
	mongoClient, err := getMongoClient("mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	// listMongoDatabase(mongoClient)

	db := mongoClient.Database("rakka")
	// listMongoCollection(db)

	coll := db.Collection("rakka")

	opt := options.Update().SetUpsert(true)
	filter := bson.D{{"name", "Lappland"}}

	var updateField bson.D
	newWaifu := Waifu{Name: "Mostima"}

	updateBytes, err := bson.Marshal(newWaifu)
	if err != nil {
		log.Fatal(err)
	}

	err = bson.Unmarshal(updateBytes, &updateField)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updateField)
	fmt.Println(filter)

	update := bson.D{{"$set", updateField}}

	_, err = coll.UpdateOne(context.TODO(), filter, update, opt)
	if err != nil {
		log.Fatal(err)
	}

	result := coll.FindOne(context.TODO(), bson.D{})

	var resultMap map[string]interface{}
	result.Decode(&resultMap)
	fmt.Println(resultMap)

}
