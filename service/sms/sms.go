package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SMSRequestBody struct {
	From      string `json:"from"`
	Text      string `json:"text"`
	To        string `json:"to"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

type sms struct {
	Key    string
	Secret string
}

func Send(phoneNumber, msg string) error {

	file, _ := ioutil.ReadFile("./conf/sms.json")

	sms := sms{}

	_ = json.Unmarshal(file, &sms)

	body := SMSRequestBody{
		APIKey:    sms.Key,
		APISecret: sms.Secret,
		To:        phoneNumber + "123",
		From:      "Test app",
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

	defer resp.Body.Close()

	return nil
}
