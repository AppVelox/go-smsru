package sms

import (
	"net/http"
	"time"
	"errors"
)

var error_no_backend = errors.New("No such backend")

//types for sms.ru
type SmsRuClient struct {
	ApiId  string       `json:"api_id"`
	Http   *http.Client `json:"-"`
	Sender string         `json:"from"`
}

type Sms struct {
	To        string            `json:"to"`
	Text      string            `json:"text"`
	Translit  bool              `json:"translit"`
	Multi     map[string]string `json:"multi"`
	From      string            `json:"from"`
	Time      time.Time         `json:"time"`
	Test      bool              `json:"test"`
	PartnerId int               `json:"partner_id"`
}

//types for iqsms.ru

type IQSMSRuClient struct {
	ApiLogin    string       `json:"login"`
	ApiPassword string       `json:"password"`
	Http        *http.Client `json:"-"`
	Sender      string       `json:"sender"`
}

type IQSms struct {
	Phone           string            `json:"phone"`
	Text            string            `json:"text"`
	WapUrl          string            `json:"wapurl"`
	Sender          string            `json:"sender"`
	Flash           string            `json:"flash"`
	ScheduleTime    time.Time      `json:"scheduleTime"`
	StatusQueueName string      `json:"statusQueueName"`
}

//types for smsc.ru

type SmsCRuClient struct {
	ApiLogin    string       `json:"login"`
	ApiPassword string       `json:"psw"`
	Http        *http.Client `json:"-"`
	Sender      string       `json:"sender"`
}

type SmsC struct {
	Phone  string            `json:"phones"`
	Mes    string             `json:"mes"`
	Sender string            `json:"sender"`
	Format string            `json:"fmt"`
}

type Response struct {
	Status string          `json:"status"`
	Id     string          `json:"id"`
	Phone  string          `json:"phone"`
}

type SMSClient interface {
	SmsSend(*CommonSms) (Response, error)
	NewSms(string,string) (*CommonSms)
	SmsStatus(string,string) (Response,error)
}

type CommonSms struct {
	Phone      string
	Message    string
	Sender     string
}

func NewSmsClient(backendInfo map[string]interface{}) (SMSClient, error) {


	if val, ok := backendInfo["backend"]; ok {
		switch val {

		case "smsru":
			var c SMSClient = &SmsRuClient{}
			smsruClient := c.(*SmsRuClient)
			smsruClient.Http = &http.Client{}
			smsruClient.ApiId = backendInfo["api_key"].(string)
			smsruClient.Sender = backendInfo["sender"].(string)

			return smsruClient, nil

		case "iqsmsru":
			var c SMSClient = &IQSMSRuClient{}
			iqsmsruClient := c.(*IQSMSRuClient)
			iqsmsruClient.Http = &http.Client{}
			iqsmsruClient.ApiLogin = backendInfo["login"].(string)
			iqsmsruClient.ApiPassword = backendInfo["password"].(string)
			iqsmsruClient.Sender = backendInfo["sender"].(string)
			return iqsmsruClient, nil
		case "smscru":
			var c SMSClient = &SmsCRuClient{}
			smscruClient := c.(*SmsCRuClient)
			smscruClient.Http = &http.Client{}
			smscruClient.ApiLogin = backendInfo["login"].(string)
			smscruClient.ApiPassword = backendInfo["password"].(string)
			smscruClient.Sender = backendInfo["sender"].(string)
			return smscruClient, nil
		}

	}

	return nil, error_no_backend

}


func (c *SmsRuClient) NewSms(to string, text string) *CommonSms {
	return &CommonSms{
		Phone:   to,
		Message: text,
		Sender: c.Sender,
	}
}

func (c *IQSMSRuClient) NewSms(to string, text string) *CommonSms {
	return &CommonSms{
		Phone:   to,
		Message: text,
		Sender: c.Sender,
	}
}

func (c *SmsCRuClient) NewSms(to string, text string) *CommonSms {
	return &CommonSms{
		Phone:   to,
		Message: text,
		Sender: c.Sender,
	}
}

