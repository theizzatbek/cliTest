package cmd

import (
	"bytes"
	"cliTest/config"
	"encoding/json"
	"net/http"
)

type SMSRequestBody struct {
	From      string `json:"from"`
	Text      string `json:"text"`
	To        string `json:"to"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

func Send(phoneNumber, msg string) error {

	sms := config.GetInstance().Nexmo

	body := SMSRequestBody{
		APIKey:    sms.Key,
		APISecret: sms.Secret,
		To:        phoneNumber,
		From:      config.GetInstance().Application,
		Text:      msg,
	}

	smsBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(smsBody))
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	return nil
}
