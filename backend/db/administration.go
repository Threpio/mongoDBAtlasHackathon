package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var collections []string{"users", "repositories", "stats"}

type collections struct {
	Names []string
	Users *options.CreateCollectionOptions
	Repos *options.CreateCollectionOptions
	Stats *options.CreateCollectionOptions
}

//GetCollectionOptions is the only way I have figured for there to be a struct that contains the options and be used.
// This cannot be the best way to do this.
func (c collections) GetCollectionOptions(name string) *options.CreateCollectionOptions {
	switch name {
	case "users":
		return c.Users
	case "repositories":
		return c.Repos
	case "stats":
		return c.Stats
	default:
		return nil
	}
}

//TODO: Clean these functions up
//TODO: Make it so that the TimeSeries collections have different parameters

//ConfigureCollections to be run on startup to check that a number of different Collections exist
//Creates the collections if not already exist
func (db *DB) ConfigureCollections() error {

	collectionReferences := collections{
		Names: []string{"users", "repositories", "stats"},
		Users: &options.CreateCollectionOptions{},
		Repos: &options.CreateCollectionOptions{},
		Stats: &options.CreateCollectionOptions{},
	}
	for _, name := range collectionReferences.Names {
		exists, err := db.CheckIfCollectionExists(name)
		if err != nil {
			return err
		}
		if !exists {
			fmt.Println("Creating collection: " + name)
			options := options.CreateCollectionOptions{}
			options.SetMaxDocuments(1000)
			options.SetCapped(true)
			options.SetSizeInBytes(100000000)
			if err = db.CreateCollection(name, &options); err != nil {
				return err
			}
		}
	}
	return db.ConfigureTimeSeriesCollections()
}

//CheckIfCollectionExists is used to check if a further DB initialisation is required.
//Returns true if a Collection with that name exists
func (db *DB) CheckIfCollectionExists(collectionName string) (bool, error) {
	//Bloody Pointers
	names, err := db.Client.Database(db.DatabaseName).ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		return false, err
	}
	for _, name := range names {
		if name == collectionName {
			return true, nil
		}
	}
	return false, nil
}

//CreateCollection creates a collection with a specified name and with creationParameters
//It does NOT check if the collection already exists and will return an error for sure if it does.
func (db *DB) CreateCollection(collectionName string, collectionCreationParams ...*options.CreateCollectionOptions) error {
	return db.Client.Database(db.DatabaseName).CreateCollection(db.BasicContext, collectionName, collectionCreationParams...)
}

func (db *DB) ConfigureTimeSeriesCollections() error {
	exists, err := db.CheckIfCollectionExists("timeSeries")
	if err != nil {
		return err
	}
	if !exists {
		fmt.Println("Creating collection: timeSeries")

		timeField := "timestamp"
		metaField := "meta"
		granularity := "seconds"

		options := options.CreateCollectionOptions{
			Capped: nil,
			TimeSeriesOptions: &options.TimeSeriesOptions{
				TimeField:   timeField,
				MetaField:   &metaField,
				Granularity: &granularity,
			},
		}

		if err = db.CreateCollection("timeSeries", &options); err != nil {
			return err
		}
	}
	return nil
}

