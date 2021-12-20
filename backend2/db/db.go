package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	dbName     string = ""
	dbUser     string = ""
	dbPassword string = ""
	dbCluster  string = ""
)

type DB struct {
	Client       *mongo.Client
	DatabaseName string
	BasicContext context.Context
}

// NewDB returns a DB object that is not connected but is optioned correctly
func NewDB() (*DB, error) {

	//uri := fmt.Sprintf("mongodb+srv://%s:%s@%s.h9r61.mongodb.net/%s?retryWrites=true&w=majority", dbUser, dbPassword, dbCluster, dbName)
	uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000"
	fmt.Println(uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	DB := &DB{
		Client:       client,
		DatabaseName: dbName,
		BasicContext: context.Background(),
	}
	return DB, nil
}

//InsertOne inserts one document into the specified collection
//Uses bson.M which does NOT care about the order of the json fields and data
func (db *DB) InsertOne(collectionName string, data bson.M ) error {
	collection := db.Client.Database(db.DatabaseName).Collection(collectionName)
	_, err := collection.InsertOne(db.BasicContext, data)
	return err
}

//FindOne returns one document from the specified collection
//Returns a bson.M that needs to be decoded ( bson.M.decode() )
//You need to check if the result.Err() is mongo.ErrNoDocuments if it returns Err
func (db *DB) FindOne(collectionName string, filter bson.M) (*mongo.SingleResult, error) {
	collection := db.Client.Database(db.DatabaseName).Collection(collectionName)
	result := collection.FindOne(db.BasicContext, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result, nil
}

//FindAllFromFilter returns a mongo.Cursor that can be iterated over
//In order to return all the documents you need to have an empty filter
func (db *DB) FindAllFromFilter(collectionName string, filter bson.M) (*mongo.Cursor, error) {
	collection := db.Client.Database(db.DatabaseName).Collection(collectionName)
	cursor, err := collection.Find(db.BasicContext, filter)
	if err != nil {
		return nil, err
	}
	// Get a list of all returned documents and print them out.
	//See the mongo.Cursor documentation for more examples of using cursors.
	// var results []bson.M if err = cursor.All(context.TODO(), &results);
	// err != nil { log.Fatal(err) }
	//for _, result := range results { fmt.Println(result) }
	return cursor, nil
}

