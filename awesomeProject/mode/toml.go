package model

import "time"

type Mail struct {
	Host  string
	User  string
	Port  int
	Pwd   string
	From  string
	To    string
	RTo   string
	IsSsl bool
}
type Job struct {
	JobName []string
}
type CronTime struct {
	CronStart string
}

type Config struct {
	Mail     Mail
	Job      Job
	CronTime CronTime
	Url      Url

	//	FetchTime FetchTime
}
type Url struct {
	ApiUrl string
}

//type FetchTime struct {
//	StartTime int64
//	EndTime   int64
//}

var StartTime = time.Now().Unix()

// var StartTime = time.Now().Add(-time.Hour * time.Duration(7*24)).Unix()
var EndTime = time.Now().Unix()
var C Config
