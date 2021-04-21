package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func xq(qq string, ti string) string {
	ti1 := time.Now().Year() + time.Now().Hour() + time.Now().Minute() + time.Now().Day()
	tm := time.Unix(int64(ti1), 0)
	tm1 := tm.Format("2006-01-02-03:04:05-PM")
	rsp, err := http.Get(surl + "/xq?qq=" + qq + "&time=" + ti + "&key=" + tm1)
	if err != nil {
		log.Println(err)
	}
	rbody, _ := io.ReadAll(rsp.Body)
	return string(rbody)
}
func ck(qq string) (string, string) {
	ti1 := time.Now().Year() + time.Now().Hour() + time.Now().Minute() + time.Now().Day()
	tm := time.Unix(int64(ti1), 0)
	tm1 := tm.Format("2006-01-02-03:04:05-PM")
	rsp, err := http.Get(surl + "/ck?qq=" + qq + "&key=" + tm1)
	if err != nil {
		log.Println(err)
	}
	rbody, _ := io.ReadAll(rsp.Body)
	var tmp2 struct {
		Id   string `Id`
		Time string `Time`
	}
	json.Unmarshal(rbody, &tmp2)
	return tmp2.Id, tmp2.Time
}
func cuser(qq string, ti string) string {
	ti1 := time.Now().Year() + time.Now().Hour() + time.Now().Minute() + time.Now().Day()
	tm := time.Unix(int64(ti1), 0)
	tm1 := tm.Format("2006-01-02-03:04:05-PM")
	fmt.Println(tm1)
	rsp, err := http.Get(surl + "/cuser?qq=" + qq + "&time=" + ti + "&key=" + tm1)
	if err != nil {
		log.Println(err)
	}
	rbody, _ := io.ReadAll(rsp.Body)
	return string(rbody)
}
