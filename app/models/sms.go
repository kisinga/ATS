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
	Phone   string `json:"phone" binding:"required"`
	Message string `json:"message" binding:"required"`
}
