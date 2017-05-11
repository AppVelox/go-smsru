package sms_test

import (
	"os"
	"testing"
	"encoding/json"
	"log"
	"github.com/AppVelox/go-smsru"
)


func getIQSMSRUClient(t *testing.T) (sms.SMSClient, error) {
	credentials_json := []byte(os.Getenv("IQSMSJSON"))
	var m map[string]interface{}
	err := json.Unmarshal(credentials_json, &m)

	if err != nil {
		log.Printf("Check your IQSMSJSON var")
		return nil, err
	}
	m["backend"] = "iqsmsru"

	client, err := sms.NewSmsClient(m)

	if err != nil {
		return nil, err
	}

	return client, nil
}

/* Test Sms
---------------------------------------------*/
func TestIQSMSRUSmsSend(t *testing.T) {
	c, err := getIQSMSRUClient(t)

	if err != nil {
		log.Fatal(err)
	}

	msg := c.NewTestSms(os.Getenv("TEST_PHONE"),"test")

	res, err := c.SmsSend(msg)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Result is %s %s %s",res.Phone,res.Status,res.Id)
}


func TestIQSMSRUSmsStatus(t *testing.T) {
	c, _ := getIQSMSRUClient(t)
	id := "1"

	res, err := c.SmsStatus(id, os.Getenv("TEST_PHONE"))

	if err != nil {
		log.Printf("%s", err)
		t.Fail()
	}
	log.Printf("Result is %s %s %s",res.Phone,res.Status,res.Id)
}
