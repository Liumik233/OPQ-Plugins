package main

import (
	"bytes"
	"encoding/json"
	"github.com/mcoo/OPQBot"
	"io/ioutil"
	"log"
	"net/http"
)

func sent2f(m *OPQBot.BotManager, uid int64, content string) {
	m.Send(OPQBot.SendMsgPack{
		ToUserUid:  uid,
		Content:    OPQBot.SendTypeTextMsgContent{Content: content},
		SendToType: OPQBot.SendToTypeFriend,
		SendType:   OPQBot.SendTypeTextMsg,
	})
}
func sent2g(m *OPQBot.BotManager, uid int64, content string) {
	m.Send(OPQBot.SendMsgPack{
		ToUserUid:  uid,
		Content:    OPQBot.SendTypeTextMsgContent{Content: content},
		SendToType: OPQBot.SendToTypeGroup,
		SendType:   OPQBot.SendTypeTextMsg,
	})
}
func Getfile(groupid int, fileid string, qq string, url1 string) string {
	url := struct {
		Url string `Url`
	}{}
	tmp := make(map[string]interface{})
	tmp["GroupID"] = groupid
	tmp["FileID"] = fileid
	tmp1, _ := json.Marshal(tmp)
	resp, err := (http.Post(url1+"/v1/LuaApiCaller?funcname=OidbSvc.0x6d6_2&timeout=10&qq="+qq, "application/json", bytes.NewBuffer(tmp1)))
	if err != nil {
		log.Fatal(err)
		return "err"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(body))
	json.Unmarshal(body, &url)
	return url.Url
}