package main

import (
	"bytes"
	"encoding/json"
	"github.com/mcoo/OPQBot"
	"io/ioutil"
	"log"
	"net/http"
)

func send2f(m *OPQBot.BotManager, uid int64, content string) {
	m.Send(OPQBot.SendMsgPack{
		ToUserUid:  uid,
		Content:    OPQBot.SendTypeTextMsgContent{Content: content},
		SendToType: OPQBot.SendToTypeFriend,
	}) //发送好友消息
}
func send2p(m *OPQBot.BotManager, uid int64, content string) {
	m.Send(OPQBot.SendMsgPack{
		ToUserUid:  uid,
		Content:    OPQBot.SendTypeTextMsgContent{Content: content},
		SendToType: OPQBot.SendToTypeFriend,
	}) //发送好友消息
}
func send2g(m *OPQBot.BotManager, uid int64, content string) {
	m.Send(OPQBot.SendMsgPack{
		ToUserUid:  uid,
		Content:    OPQBot.SendTypeTextMsgContent{Content: content},
		SendToType: OPQBot.SendToTypeGroup,
	}) //发送群消息
}
func send2gp(m *OPQBot.BotManager, uid int64, content string, picurl string) {
	m.Send(OPQBot.SendMsgPack{
		ToUserUid:  uid,
		Content:    OPQBot.SendTypePicMsgByUrlContent{Content: content, PicUrl: picurl},
		SendToType: OPQBot.SendToTypeGroup,
	}) //发送群消息
}
func Getfile(groupid int64, fileid string, qq string, url1 string) string {
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
func setjy(groupid int64, qq string, url1 string, time string, uid string) bool {
	ret := struct {
		Ret int `Ret`
	}{}
	tmp := make(map[string]interface{})
	tmp["GroupID"] = groupid
	tmp["ShutUpUserID"] = uid
	tmp["ShutTime"] = time
	tmp1, _ := json.Marshal(tmp)
	resp, err := (http.Post(url1+"/v1/LuaApiCaller?funcname=OidbSvc.0x570_8&timeout=10&qq="+qq, "application/json", bytes.NewBuffer(tmp1)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(body))
	json.Unmarshal(body, &ret)
	if ret.Ret == 0 {
		return true
	} else {
		return false
	}
}
func setgg(groupid int64, qq string, url1 string, title string, text string) bool {
	ret := struct {
		Ret int `Ret`
	}{}
	tmp := make(map[string]interface{})
	tmp["GroupID"] = groupid
	tmp["Title"] = title
	tmp["Text"] = text
	tmp["Pinned"] = 0
	tmp["Type"] = 10
	tmp1, _ := json.Marshal(tmp)
	resp, err := (http.Post(url1+"/v1/Group/Announce&timeout=10&qq="+qq, "application/json", bytes.NewBuffer(tmp1)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(body))
	json.Unmarshal(body, &ret)
	if ret.Ret == 0 {
		return true
	} else {
		return false
	}
}
func ch(groupid int64, qq string, url1 string, seq int64, random int64) bool {
	ret := struct {
		Ret int `Ret`
	}{}
	tmp := make(map[string]interface{})
	tmp["GroupID"] = groupid
	tmp["MsgSeq"] = seq
	tmp["MsgRandom"] = random
	tmp1, _ := json.Marshal(tmp)
	resp, err := (http.Post(url1+"/v1/LuaApiCaller?funcname=PbMessageSvc.PbMsgWithDraw&timeout=10&qq="+qq, "application/json", bytes.NewBuffer(tmp1)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(body))
	json.Unmarshal(body, &ret)
	if ret.Ret == 0 {
		return true
	} else {
		return false
	}
}
