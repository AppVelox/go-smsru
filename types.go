package sms

import (
	"net/http"
	"errors"
)

var error_no_backend = errors.New("No such backend")
var error_no_apiid = errors.New("No ApiId")
var error_no_sender = errors.New("No parameter Sender (it should be at least empty)")
var error_no_login= errors.New("No parameter Login (it should be at least empty)")
var error_no_password= errors.New("No parameter Password (it should be at least empty)")

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
	NewSms(string, string) (*CommonSms)
	NewTestSms(string, string) (*CommonSms)
	SmsStatus(string, string) (Response, error)
}

type CommonSms struct {
	Phone   string
	Message string
	Sender  string
	Test    bool
}

func NewSmsClient(backendInfo map[string]interface{}) (SMSClient, error) {

	if val, ok := backendInfo["backend"]; ok {
		switch val {

		case "smsru":

			var c SMSClient = &SmsRuClient{}
			smsruClient := c.(*SmsRuClient)
			smsruClient.Http = &http.Client{}

			if val, ok := backendInfo["api_key"]; ok {
				smsruClient.ApiId = val.(string)
			} else {
				return nil, error_no_apiid
			}

			if val, ok := backendInfo["sender"]; ok {
				smsruClient.Sender = val.(string)
			} else {
				return nil, error_no_sender
			}

			return smsruClient, nil

		case "iqsmsru":

			var c SMSClient = &IQSMSRuClient{}
			iqsmsruClient := c.(*IQSMSRuClient)
			iqsmsruClient.Http = &http.Client{}

			if val, ok := backendInfo["login"]; ok {
				iqsmsruClient.ApiLogin = val.(string)
			} else {
				return nil, error_no_login
			}

			if val, ok := backendInfo["password"]; ok {
				iqsmsruClient.ApiPassword = val.(string)
			} else {
				return nil, error_no_password
			}

			if val, ok := backendInfo["sender"]; ok {
				iqsmsruClient.Sender = val.(string)
			} else {
				return nil, error_no_sender
			}

			return iqsmsruClient, nil
		case "smscru":

			var c SMSClient = &SmsCRuClient{}
			smscruClient := c.(*SmsCRuClient)
			smscruClient.Http = &http.Client{}

			if val, ok := backendInfo["login"]; ok {
				smscruClient.ApiLogin = val.(string)
			} else {
				return nil, error_no_login
			}

			if val, ok := backendInfo["password"]; ok {
				smscruClient.ApiPassword = val.(string)
			} else {
				return nil, error_no_password
			}

			if val, ok := backendInfo["sender"]; ok {
				smscruClient.Sender = val.(string)
			} else {
				return nil, error_no_sender
			}
			return smscruClient, nil
		}
	}

	return nil, error_no_backend

}
