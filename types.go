package sms

import (
	"net/http"
	"time"
)

//types for sms.ru
type SmsRuClient struct {
	ApiId string       `json:"api_id"`
	Http  *http.Client `json:"-"`
	From  string         `json:"from"`
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
	Phone     string            `json:"phone"`
	Text      string            `json:"text"`
	WapUrl    string            `json:"wapurl"`
	Sender    string 	    `json:"sender"`
	Flash     string            `json:"flash"`
	ScheduleTime time.Time      `json:"scheduleTime"`
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
	Phone     string            `json:"phones"`
	Mes      string             `json:"mes"`
	Sender    string 	    `json:"sender"`
	Format    string 	    `json:"fmt"`
}


type Response struct {
	Status    string          `json:"status"`
	Id        string          `json:"id"`
	Phone     string          `json:"phone"`
}

