package models

type Meter struct {
	ID          string  `json:"ID"  bson:"_id,omitempty"`
	MeterNumber string  `json:"meterNumber"`
	Location    *string `json:"location"`
}
