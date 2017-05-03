## Purpose
This is the simplest multi-wrapper for russian sms services.

##### Supported backends
- sms.ru
- iqsms.ru
- smsc.ru

## How to use
```go
package main

import (
	"github.com/AppVelox/go-smsru"
	"log"
)

const (
	API_ID = "your sms.ru api_id"
	IQSMSRU_Login = "your iqsms.ru login"
	IQSMSRU_Pass = "your iqsms.ru password"
	SMSCRU_Login = "your smsc.ru login"
	SMSCRU_Pass  = "your smsc.ru password"
)

var sms_ru_credentials = make(map[string]interface{}) //choose
var smsc_ru_credentials = make(map[string]interface{}) // any
var iqsms_ru_credentials = make(map[string]interface{}) // you want

func main() {
    // sms.ru credentials and/or sender
	sms_ru_credentials["backend"] = "smsru"
	sms_ru_credentials["api_key"] = API_ID
	sms_ru_credentials["sender"] = ""
	
	//iqsms.ru credentials and/or sender
	iqsms_ru_credentials["backend"] = "iqsmsru"
	iqsms_ru_credentials["login"] = IQSMSRU_Login
	iqsms_ru_credentials["password"] = IQSMSRU_Pass
	iqsms_ru_credentials["sender"] = ""
	
	//smsc.ru credentials and/or sender
	smsc_ru_credentials["backend"] = "smscru"
	smsc_ru_credentials["login"] = SMSCRU_Login
	smsc_ru_credentials["password"] = SMSCRU_Pass
	smsc_ru_credentials["sender"] = ""

	client, err := sms.NewSmsClient(sms_ru_credentials)
	if err != nil {
		log.Printf("%s", err) 
	}
	//Sending sms
	msg := client.NewSms("79991234567", "Sample Text")
	res, err := client.SmsSend(msg)

	if err != nil {
		log.Printf("%s", err)
	} else {
		log.Printf("Phone: %s Status: %s Id: %s", res.Phone, res.Status, res.Id)
	}

	//check status
	res, err = client.SmsStatus("sms_string_id", "phone")
	if err != nil {
		log.Printf("%s", err)
	} else {
		log.Printf("Phone: %s Status: %s Id: %s", res.Phone, res.Status, res.Id)
	}

}

```
