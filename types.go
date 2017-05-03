package sms

import (
	"net/http"
	"errors"
)

var error_no_backend = errors.New("No such backend")

type SmsRuClient struct {
	ApiId  string       `json:"api_id"`
	Http   *http.Client `json:"-"`
	Sender string         `json:"from"`
}

type IQSMSRuClient struct {
	ApiLogin    string       `json:"login"`
	ApiPassword string       `json:"password"`
	Http        *http.Client `json:"-"`
	Sender      string       `json:"sender"`
}

type SmsCRuClient struct {
	ApiLogin    string       `json:"login"`
	ApiPassword string       `json:"psw"`
	Http        *http.Client `json:"-"`
	Sender      string       `json:"sender"`
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



