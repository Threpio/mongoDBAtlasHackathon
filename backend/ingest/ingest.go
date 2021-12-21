package ingest

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/db"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/logger"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/types"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
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

func (c *Controller) structuredIngest(in *types.StructuredIngestRequest) (count int, err error) {
	count = 0
	var toInsert []interface{}
	for _, e := range in.Events {
		data := bson.M{
			"timestamp": time.Unix(e.Timestamp, 0),
			"meta":      e.Data,
		}
		toInsert = append(toInsert, data)
		count++
	}

	return count, c.DB.InsertMany("timeSeries", toInsert)
}

func (c *Controller) searchIngest(in *types.SearchIngestRequest) ([]byte, error) {

	//TODO: Validate input request within params
	// If in.Query == ""/nil then make it {}
	//  ...{
	//	... "timestamp": {
	//	..... $gte: ISODate("2021-12-21T10:58:52.000Z"),
	//	..... $lt: ISODate("2021-12-21T10:58:59.000Z")
	//	..... }
	//	... }

	query := bson.M{
		"timestamp": bson.M{
			"$gte": time.Unix(in.TimeFrom, 0),
			"$lt":  time.Unix(in.TimeTo, 0),
		},
	}

	cursor, err := c.DB.FindAllFromFilter("timeSeries", query)
	if err != nil {
		return nil, err
	}

	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}
	body, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	// TODO: Handle null

	return body, nil
}