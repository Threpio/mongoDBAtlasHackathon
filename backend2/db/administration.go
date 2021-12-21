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

//ConfigureCollections to be run on startup to check that a number of different Collections exist
//Creates the collections if not already exist
func (db *DB) ConfigureCollections() error {

	collectionReferences := collections{
		Names: []string{"users", "repositories", "stats"},
		Users: nil,
		Repos: nil,
		Stats: nil,
	}
	for _, name := range collectionReferences.Names {
		exists, err := db.CheckIfCollectionExists(name)
		if err != nil {
			return err
		}
		if !exists {
			fmt.Println("Creating collection: " + name)
			err = db.CreateCollection(name, collectionReferences.GetCollectionOptions(name))
			if err != nil {
				return err
			}
		}
		if err := db.CreateCollection(name); err != nil {
			return err
		}
	}
	return nil
}

//CheckIfCollectionExists is used to check if a further DB initialisation is required.
//Returns true if a Collection with that name exists
func (db *DB) CheckIfCollectionExists(collectionName string) (bool, error) {
	//Bloody Pointers
	names, err := db.Client.Database(db.DatabaseName).ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		return false, err
	}
	fmt.Println(names)
	for _, name := range names {
		fmt.Println(name)
		if name == collectionName {
			fmt.Println("Collection exists")
			return true, nil
		}
	}
	return false, nil
}

//CreateCollection creates a collection with a specified name and with creationParameters
//It does NOT check if the collection already exists and will return an error for sure if it does.
func (db *DB) CreateCollection(collectionName string, collectionCreationParams ...*options.CreateCollectionOptions) error {
	if err := db.Client.Database(db.DatabaseName).CreateCollection(db.BasicContext, collectionName, collectionCreationParams...); err != nil {
		return err
	}
	return nil
}

type collectionsCursorResponse struct {
}
