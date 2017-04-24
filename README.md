```
package main

import (
    "log"
    "github.com/AppVelox/go-smsru"
)

const (
	API_ID        = "your sms.ru api_id"
	IQSMSRU_Login = ""
	IQSMSRU_Pass  = ""
)

func main() {
    // client := sms.NewIQSmsRuClient(IQSMSRU_Login,IQSMSRU_Pass,"") // iqsms.ru needs login and password
    client := sms.NewSmsRuClient(API_ID,"")  // sms.ru needs api_id
	msg := client.NewSms("phone", "SUPER")

	res, err := client.SmsSend(msg)
	if err != nil {
		log.Panic(err)
	} else {
		log.Printf("Status = %d, Id = %s", res.Status, res.Id)
	}

	res, _ = client.SmsStatus(res.Id)
	fmt.Println(res)
}
```