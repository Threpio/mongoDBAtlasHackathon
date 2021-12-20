package ingest

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/db"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/logger"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/types"
	"go.mongodb.org/mongo-driver/bson"
)

type Controller struct {
	DB     db.DB
	Router func(r chi.Router)
	logger logger.Logger
}

func NewController(db db.DB, logger logger.Logger) (*Controller, error) {
	controller := &Controller{
		DB:     db,
		logger: logger,
	}
	controller.Router = controller.ingestRouter()
	return controller, nil
}

func (c *Controller) structuredIngest(in *types.StructuredIngestRequest) error {

	return nil
}

func (c *Controller) searchIngest(in *types.SearchIngestRequest) (*types.SearchIngestResponse, error) {

	//TODO: Validate input request within params
	var filter bson.M

	if err := bson.UnmarshalExtJSON([]byte(in.Query), true, &filter); err != nil {
		return nil, err
	}

	cursor, err := c.DB.FindAllFromFilter("test", filter)
	if err != nil {
		return nil, err
	}

	var results []bson.M
	var singleResult types.Event
	if err := cursor.All(nil, &results); err != nil {
		return nil, err
	}
	//TODO: Something with this result.
	for _, result := range results {
		if err := bson.UnmarshalExtJSON([]byte(result["event"].(string)), true, &singleResult); err != nil {
			c.logger.Debug("failed to unmarshall an Ext JSON string in the searchIngest function")
			return nil, err
		}
		fmt.Println(result)
	}

	//TODO: Iterate over results
	//TODO: Check errors

	//Convert a string to a bson.M
	//var filter bson.M
	//if err := bson.UnmarshalExtJSON([]byte(in.Query), true, &filter); err != nil {
	//	return nil, err
	//}

	return nil, nil
}
