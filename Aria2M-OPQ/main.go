package main

import (
	"encoding/json"
	"fmt"
	"github.com/mcoo/OPQBot"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func main() {
	fmt.Println("Aria2M_for_OPQ_ver.0.1b")
	fmt.Println("By Liumik")
	if !Exists("./config.json") {
		tmp := make(map[string]interface{})
		var url string
		var token string
		var qq int64
		var site string
		fmt.Println("\n请输入OPQ的Web地址: ")
		fmt.Scan(&site)
		fmt.Println("\n请输入Bot账号: ")
		fmt.Scan(&qq)
		fmt.Println("\n请输入url")
		fmt.Scan(&url)
		fmt.Println("请输入token")
		fmt.Scan(&token)
		tmp["Site"] = site
		tmp["Qq"] = qq
		tmp["Url"] = url
		tmp["Token"] = token
		tmp1, _ := json.Marshal(tmp)
		c1, err := os.Create("config.json")
		defer c1.Close()
		if err != nil {
			log.Println("cerr:", err)
			os.Exit(1)
		}
		c1.Write(tmp1)
	}
	c1, err := os.OpenFile("./config.json", os.O_RDONLY, 0600)
	defer c1.Close()
	if err != nil {
		log.Println("openerr:", err)
		os.Exit(1)
	}
	cb, _ := ioutil.ReadAll(c1)
	conf1 := struct {
		Site  string `Site`
		Qq    int64  `Qq`
		Url   string `Url`
		Token string `Token`
	}{}
	json.Unmarshal(cb, &conf1)
	opqBot := OPQBot.NewBotManager(conf1.Qq, conf1.Site)
	err1 := opqBot.Start()
	if err1 != nil {
		log.Println(err.Error())
	}
	defer opqBot.Stop()
	err = opqBot.AddEvent(OPQBot.EventNameOnFriendMessage, func(botQQ int64, packet OPQBot.FriendMsgPack) {
	})

	err = opqBot.AddEvent(OPQBot.EventNameOnGroupShut, func(botQQ int64, packet OPQBot.GroupShutPack) {
		//log.Println(botQQ, packet)
		//log.Println("群聊消息: ", mess.FromNickName+"<"+strconv.FormatInt(mess.FromUserID, 10)+">: "+mess.Content)
		if strings.HasPrefix(packet.EventMsg.Content, "addurl") {
			gid, err := Addurl(strings.Trim(packet.EventMsg.Content, "addurl"), conf1.Url, conf1.Token)
			if err != nil {
				sent2g(&opqBot, packet.EventData.GroupID, "error:"+err.Error())
			} else {
				sent2g(&opqBot, packet.EventData.GroupID, "Successful,gid:"+gid)
			}
		}
		if strings.HasPrefix(packet.EventMsg.Content, "status") {
			rsp, err := Filestatus(strings.Trim(packet.EventMsg.Content, "status"), conf1.Url, conf1.Token)
			if err != nil {
				sent2g(&opqBot, packet.EventData.GroupID, "error:"+err.Error())
			} else {
				sent2g(&opqBot, packet.EventData.GroupID, rsp)
			}
		}
		/*if strings.HasPrefix(fileinfo.FileName, "addbt") {
			urlt := iotqq.Getfile(mess.FromGroupID, fileinfo.FileID)
			gid, err := Addbt(urlt, aria2)
			if err != nil {
				iotqq.Send(mess.FromGroupID, 2, "error:"+err.Error())
			} else {
				iotqq.Send(mess.FromGroupID, 2, "Successful,gid:"+gid)
			}
		}*/
		if strings.HasPrefix(packet.EventMsg.Content, "stop") {
			err := Stop(strings.TrimPrefix(packet.EventMsg.Content, "stop"), conf1.Url, conf1.Token)
			if err != nil {
				sent2g(&opqBot, packet.EventData.GroupID, "error:"+err.Error())
			} else {
				sent2g(&opqBot, packet.EventData.GroupID, "Successful")
			}
		}
		if strings.HasPrefix(packet.EventMsg.Content, "start") {
			err := Start(strings.TrimPrefix(packet.EventMsg.Content, "start"), conf1.Url, conf1.Token)
			if err != nil {
				sent2g(&opqBot, packet.EventData.GroupID, "error:"+err.Error())
			} else {
				sent2g(&opqBot, packet.EventData.GroupID, "Successful")
			}

		}
		if strings.HasPrefix(packet.EventMsg.Content, "del") {
			err := Del(strings.TrimPrefix(packet.EventMsg.Content, "del"), conf1.Url, conf1.Token)
			if err != nil {
				sent2g(&opqBot, packet.EventData.GroupID, "error:"+err.Error())
			} else {
				sent2g(&opqBot, packet.EventData.GroupID, "Successful")
			}
		}
	})
	if err != nil {
		log.Println(err.Error())
	}
	err = opqBot.AddEvent(OPQBot.EventNameOnConnected, func() {
		log.Println("连接成功！！！")
	})
	if err != nil {
		log.Println(err.Error())
	}
	err = opqBot.AddEvent(OPQBot.EventNameOnDisconnected, func() {
		log.Println("连接断开！！")
	})
	if err != nil {
		log.Println(err.Error())
	}
	err = opqBot.AddEvent(OPQBot.EventNameOnOther, func(botQQ int64, e interface{}) {
		log.Println(e)
	})
	if err != nil {
		log.Println(err.Error())
	}
	time.Sleep(1 * time.Hour)
}
