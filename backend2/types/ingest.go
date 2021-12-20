package types

type Event struct {
	Timestamp string                 `json:"timestamp,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

//Ingest

type StructuredIngestRequest struct {
	Event Event `json:"event"`
}

//Search

type SearchIngestRequest struct {
	TimeFrom int64  `json:"time_from"`
	TimeTo   int64  `json:"time_to"`
	Query    string `json:"query"`
}

type SearchIngestResponse struct {
	SearchIngestRequest SearchIngestRequest `json:"search_ingest_request"`
	Events              []Event             `json:"events"`
}
