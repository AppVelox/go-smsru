package sms

import (
	"fmt"
	"errors"
	"net/url"
	"io/ioutil"
	"log"
	"strings"
)

const IQSMSRU_API_URL = "http://gate.iqsms.ru"


func (c *IQSMSRuClient) makeRequest(endpoint string, params url.Values) (Response, error) {
	params.Set("login", c.ApiLogin)
	params.Set("password", c.ApiPassword)

	url := IQSMSRU_API_URL + endpoint + "?" + params.Encode()

	resp, err := c.Http.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	var messageAndStatus = strings.Split(string(body), "=")

	fmt.Println(messageAndStatus)

	if len(messageAndStatus) != 2 {
		msg := fmt.Sprintf(string(body))
		return Response{}, errors.New(msg)
	}

	res := Response{Status: messageAndStatus[1]}
	res.Id = messageAndStatus[0]
	return res, nil

}






func (c *IQSMSRuClient) SmsSend(p *CommonSms) (Response, error) {
	var params = url.Values{}

	params.Set("phone", p.Phone)
	params.Set("text", p.Message)

	if len(p.Sender) > 0 {
		params.Set("sender", p.Sender)
	}

	res, err := c.makeRequest("/send", params)
	if err != nil {
		return Response{}, err
	}

	res.Phone = p.Phone

	return res, nil
}

// //SmsStatus will get a status of message
func (c *IQSMSRuClient) SmsStatus(id string, phone string) (Response, error) {
	params := url.Values{}
	params.Set("id", id)

	res, err := c.makeRequest("/status", params)
	if err != nil {
		return Response{}, err
	}
	res.Phone = phone

	return res, nil
}