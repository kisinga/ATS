package sms

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/kisinga/ATS/app/models"
)

type SMS interface {
	Send(text models.Text) (models.Text, error)
}

type AfricasTalking struct {
	uri      string
	Username string
	Key      string
	SenderID string
}

func NewSMS(config AfricasTalking) SMS {
	return &config
}

func (a *AfricasTalking) Send(text models.Text) (models.Text, error) {
	values := url.Values{
		"username": []string{a.Username},
		"to":       []string{text.Phone},
		"message":  []string{text.Message},
		"from":     []string{},
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", a.uri, strings.NewReader(values.Encode()))
	if err != nil {
		return text, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("apikey", a.Key)
	resp, err := client.Do(req)
	if err != nil {
		return text, err
	}

	type RecipientData struct {
		StatusCode models.Status `json:"StatusCode"`
		Number     string        `json:"number"`
		Cost       string        `json:"cost"`
		Status     string        `json:"status"`
		MessageID  string        `json:"messageId"`
		Parts      string        `json:"key"`
	}
	type Data struct {
		Message    string          `json:"Message"`
		Recipients []RecipientData `json:"Recipients"`
	}
	response := &struct {
		SMSMessageData Data `json:"SMSMessageData"`
	}{}

	if resp.StatusCode >= 300 {
		return text, errors.New(resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return text, err
	}
	if len(response.SMSMessageData.Recipients) == 0 {
		return text, errors.New("nil recipient response")
	}

	// text.Cost, err = strconv.ParseFloat(strings.Split(response.SMSMessageData.Recipients[0].Cost, " ")[1], 64)
	// if err != nil {
	// 	return text, err
	// }
	// text.MessageID = response.SMSMessageData.Recipients[0].MessageID
	// text.Status = response.SMSMessageData.Recipients[0].StatusCode

	return text, nil
}
