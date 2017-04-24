package sms

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
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

func NewSmsCRuClient(login string, password string, sender string) *SmsCRuClient {

	c := &SmsCRuClient{
		ApiLogin:    login,
		ApiPassword: password,
		Http:        &http.Client{},
		Sender:      sender,
	}
	return c
}

// NewSms creates a new message

func (c *SmsCRuClient) NewSms(to string, text string) *SmsC {
	return &SmsC{
		Phone:  to,
		Mes:    text,
		Sender: c.Sender,
		Format: "3",
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

// SmsSend will send a Sms item to Service
func (c *SmsCRuClient) SmsSend(p *SmsC) (Response, error) {
	var params = url.Values{}

	params.Set("phones", p.Phone)
	params.Set("fmt", p.Format)

	params.Set("mes", p.Mes)

	if len(p.Sender) > 0 {
		params.Set("sender", p.Sender)
	}

	res, body, err := c.makeRequest("/sys/send.php", params)
	if err != nil {
		return Response{}, err
	}

	var f struct {
		Id        int          `json:"id"`
		Cnt       int          `json:"cnt"`
		Error     string          `json:"error"`
		ErrorCode int      `json:"error_code"`
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
		res.Status = f.Error
		return res, nil
	} else {
		res.Status = "OK"
	}

	return res, nil
}

// SmsStatus will get a status of message
func (c *SmsCRuClient) SmsStatus(id string, phone string) (Response, error) {
	params := url.Values{}
	params.Set("id", id)
	params.Set("phone", phone)
	params.Set("fmt", "3")

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
		res.Status = val.(string)
	}

	res.Id = id
	res.Phone = phone

	return res, nil
}
