package models

type Status float64

const (
	Processed             Status = 100
	Sent                  Status = 101
	Queued                Status = 102
	RiskHold              Status = 401
	InvalidSenderId       Status = 402
	InvalidPhoneNumber    Status = 403
	UnsupportedNumberType Status = 404
	InsufficientBalance   Status = 405
	UserInBlacklist       Status = 406
	CouldNotRoute         Status = 407
	InternalServerError   Status = 500
	GatewayError          Status = 501
	RejectedByGateway     Status = 502
)

type SendErrorCode int
type SendError struct {
	SendErrorCode
	Text string
}
type Text struct {
	Phone string `json:"phone" binding:"required"`
	//From is the value of the senderID in case the company has bought one
	From    string `json:"from" binding:"required"`
	Message string `json:"message" binding:"required"`
	Status  `json:"-"`
	Cost    float64 `json:"-"`
	//MessageID is a value assigned to the message by the 3rd party provider
	MessageID string `json:"-"`
	//BatchID represents the ID of every batch of SMS
	//Every request with an arrays of SMS's is treated as a batch
}
