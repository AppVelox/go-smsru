package sms

import (
	"time"
	"strconv"
	"fmt"
	"errors"
	"net/url"
	"io/ioutil"
	"log"
	"strings"
	"net/http"
)

const IQSMSRU_API_URL = "http://gate.iqsms.ru"


func NewIQSmsRuClient(login string, password string,sender string) *IQSMSRuClient {

	c := &IQSMSRuClient{
		ApiLogin:   login,
		ApiPassword: password,
		Http:        &http.Client{},
		Sender:      sender,
	}

	return c
}


func(c *IQSMSRuClient) NewSms(phone string, text string) *IQSms{
	return &IQSms{
		Phone:   phone,
		Text: text,
		Sender: c.Sender,
	}
}

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

func (c *IQSMSRuClient) SmsSend(p *IQSms) (Response, error) {
	var params = url.Values{}

	params.Set("phone", p.Phone)
	params.Set("text", p.Text)

	if len(p.WapUrl) > 0 {
		params.Set("wapurl", p.WapUrl)
	}

	if len(p.Sender) > 0 {
		params.Set("sender", p.Sender)
	}

	if len(p.Flash) > 0 {
		params.Set("flash", p.Flash)
	}

	if p.ScheduleTime.After(time.Now()) {
		val := strconv.FormatInt(p.ScheduleTime.Unix(), 10)
		params.Set("scheduleTime", val)
	}
	if len(p.StatusQueueName) > 0 {
		params.Set("statusQueueName", p.StatusQueueName)
	}

	res, err := c.makeRequest("/send", params)
	if err != nil {
		return Response{}, err
	}

	res.Phone = p.Phone

	return res, nil
}

// SmsStatus will get a status of message
func (c *IQSMSRuClient) SmsStatus(id string) (Response, error) {
	params := url.Values{}
	params.Set("id", id)

	res, err := c.makeRequest("/status", params)
	if err != nil {
		return Response{}, err
	}

	return res, nil
}