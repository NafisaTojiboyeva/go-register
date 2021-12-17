package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NafisaTojiboyeva/go-register/models"
)

func SendConfirmCode(phone, code string) {

	body := models.SMSRequestBody{
		APIKey:    "f0fb1f37",
		APISecret: "mKYcnfgDzxy8twqd",
		From:      "Hacker",
		To:        phone,
		Text:      code,
	}

	smsBody, err := json.Marshal(body)

	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post(
		"https://rest.nexmo.com/sms/json",
		"application/json",
		bytes.NewBuffer(smsBody),
	)

	defer response.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(responseBody))
}
