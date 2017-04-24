package sms_test

import (
	"os"
	"testing"
	"github.com/AppVelox/go-smsru"
)

func getPhone() string {
	return os.Getenv("phone")
}

func getClient(t *testing.T) *sms.SmsClient {
	apiId := os.Getenv("api_id")
	return sms.NewClient(apiId)
}

/* Test Sms
---------------------------------------------*/
func TestSmsSend(t *testing.T) {
	c := getClient(t)

	msg := sms.NewSms(getPhone(), "Sample")
	msg.Test = true

	_, err := c.SmsSend(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSmsMultiSend(t *testing.T) {
	c := getClient(t)

	msg := sms.NewSms(getPhone(), "Sample")
	multi := sms.NewMulti(msg)
	multi.Test = true

	_, err := c.SmsSend(multi)

	if err != nil {
		t.Fail()
	}
}

func TestSmsStatus(t *testing.T) {
	c := getClient(t)
	id := "201600-1000000"

	_, err := c.SmsStatus(id)

	if err != nil {
		t.Fail()
	}
}