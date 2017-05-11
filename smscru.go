package sms

import (
	"net/url"
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
	"errors"
)

const SMSCRU_API_URL = "https://smsc.ru"

var SMSCcodeStatus map[int]string = map[int]string{
	-3: "Message not found",
	-1: "Wait for sending",
	0:  "Transmitted to operator",
	1:  "Delivered",
	2:  "Message was read",
	3:  "Overdue",
	20: "Undeliverable",
	22: "Incorrect number",
	23: "Forbidden",
	24: "Not enough money",
	25: "Inaccessible number",
}

func (c *SmsCRuClient) NewSms(to string, text string) *CommonSms {
	return &CommonSms{
		Phone:   to,
		Message: text,
		Sender:  c.Sender,
	}
}

func (c *SmsCRuClient) NewTestSms(to string, text string) *CommonSms {
	return &CommonSms{
		Phone:   to,
		Message: text,
		Sender:  c.Sender,
		Test:    true,
	}
}

func (c *SmsCRuClient) makeRequest(endpoint string, params url.Values) (Response, []byte, error) {
	params.Set("login", c.ApiLogin)
	params.Set("psw", c.ApiPassword)

	url := SMSCRU_API_URL + endpoint + "?" + params.Encode()

	resp, err := c.Http.Get(url)
	if err != nil {
		return Response{}, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return Response{}, body, nil

}

func (c *SmsCRuClient) SmsSend(p *CommonSms) (Response, error) {
	var params = url.Values{}

	params.Set("phones", p.Phone)
	params.Set("fmt", "3")
	params.Set("mes", p.Message)

	if p.Test {
		params.Set("cost", "1")
	}

	if len(p.Sender) > 0 {
		params.Set("sender", p.Sender)
	}

	log.Printf("Trying to send message: '%s' to %s", p.Message, p.Phone)

	res, body, err := c.makeRequest("/sys/send.php", params)
	if err != nil {
		return Response{}, err
	}

	var f struct {
		Id        int          `json:"id"`
		Cnt       int          `json:"cnt"`
		Error     string       `json:"error"`
		ErrorCode int          `json:"error_code"`
		Cost      string       `json:"cost"`
	}

	err = json.Unmarshal(body, &f)

	if err != nil {
		return Response{}, err
	}

	res.Phone = p.Phone



	if f.Id > 0 {
		res.Id = strconv.Itoa(f.Id)
	}

	if len(f.Error) > 0 {
		//res.Status = f.Error
		return Response{}, errors.New(f.Error)
	} else {
		res.Status = "OK"
	}
	if len(f.Cost) > 0 {
		res.Id = "Mock_Id"
	}

	return res, nil
}

// SmsStatus will get a status of message
func (c *SmsCRuClient) SmsStatus(id string, phone string) (Response, error) {
	params := url.Values{}
	params.Set("id", id)
	params.Set("phone", phone)
	params.Set("fmt", "3")

	log.Printf("Trying to get status of message: '%s'", id)

	res, body, err := c.makeRequest("/sys/status.php", params)
	if err != nil {
		return Response{}, err
	}

	var f interface{}

	err = json.Unmarshal(body, &f)

	if err != nil {
		return Response{}, err
	}

	m := f.(map[string]interface{})

	if val, ok := m["status"]; ok {
		res.Status = SMSCcodeStatus[int(val.(float64))]
	}

	if val, ok := m["error"]; ok {
		return Response{}, errors.New(val.(string))
	}

	res.Id = id
	res.Phone = phone

	return res, nil
}
