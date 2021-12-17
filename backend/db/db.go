package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Time series Creation parameters

const (
	Granularity = "seconds"
	MetaField   = "meta"
)

//GetClient Create a MongoDB Client
func GetClient() (client *mongo.Client, err error) {
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	//TODO: Cancel Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
}

type DB struct {
	Client *mongo.Client
	DbName string
}

//CreateCollection Create a Collection (Not a time series one)
func (db *DB) CreateCollection(collectionName string) (err error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// TODO: Parameterize database name
	err = db.Client.Database(db.DbName).CreateCollection(ctx, collectionName)
	return nil
}

//CreateTimeSeriesCollection Does what it says on the tin
func (db *DB) CreateTimeSeriesCollection(collectionName string) (err error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Not sure why I have to do this - this way
	timeSeriesCollectionOptions := options.TimeSeriesOptions{TimeField: "timestamp"}
	timeSeriesCollectionOptions.SetMetaField(MetaField).SetGranularity(Granularity)
	createOptions := options.CreateCollectionOptions{
		TimeSeriesOptions: &timeSeriesCollectionOptions,
	}

	err = db.Client.Database(db.DbName).CreateCollection(ctx, collectionName, &createOptions)
	return err
}

func (db *DB) InsertOne(collectionName string, data interface{}) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := db.Client.Database(db.DbName).Collection(collectionName)
	_, err = collection.InsertOne(ctx, data)
	return err
}

func (db *DB) FindOne(collectionName string, filter interface{}) *mongo.SingleResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := db.Client.Database(db.DbName).Collection(collectionName)
	data := collection.FindOne(ctx, filter)
	return data
}
