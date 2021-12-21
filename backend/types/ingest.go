package types

type Event struct {
	Timestamp int64                  `json:"timestamp,omitempty" bson:"timestamp"`
	Data      map[string]interface{} `json:"data,omitempty" bson:"data"`
}

//Ingest

type StructuredIngestRequest struct {
	Events []Event `json:"event"`
}

//Search

type SearchIngestRequest struct {
	TimeFrom int64             `json:"time_from"`
	TimeTo   int64             `json:"time_to"`
}

