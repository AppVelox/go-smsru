package sms_test

import (
	"os"
	"testing"
	"encoding/json"
	"log"
	"github.com/AppVelox/go-smsru"
)

func getPhone() string {
	return os.Getenv("phone")
}

func getSMSRUClient(t *testing.T) (sms.SMSClient, error) {
	credentials_json := []byte(os.Getenv("SMSRUJSON"))
	var m map[string]interface{}
	err := json.Unmarshal(credentials_json, &m)

	if err != nil {
		log.Printf("Check your SMSRUJSON var")
		return nil, err
	}
	m["backend"] = "smsru"

	client, err := sms.NewSmsClient(m)

	if err != nil {
		return nil, err
	}

	return client, nil
}

/* Test Sms
---------------------------------------------*/
func TestSMSRUSmsSend(t *testing.T) {
	c, err := getSMSRUClient(t)

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


func TestSMSRUSmsStatus(t *testing.T) {
	c, _ := getSMSRUClient(t)
	id := "201600-1000000"

	res, err := c.SmsStatus(id, os.Getenv("TEST_PHONE"))

	if err != nil {
		log.Printf("%s", err)
		t.Fail()
	}
	log.Printf("Result is %s %s %s",res.Phone,res.Status,res.Id)
}
