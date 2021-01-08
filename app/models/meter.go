package models

type Meter struct {
	MeterNumber string  `json:"meterNumber"`
	Location    *string `json:"location"`
	*BaseModel
}

func (Meter) IsBaseObject() {}
