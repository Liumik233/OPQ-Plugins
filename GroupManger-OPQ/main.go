package main

import (
	"encoding/json"
	"fmt"
	"github.com/mcoo/OPQBot"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var ver = "GroupManger_for_OPQ_ver.0.1a"

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
	fmt.Println(ver)
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
	}{}
	json.Unmarshal(cb, &conf1)
	opqBot := OPQBot.NewBotManager(conf1.Qq, conf1.Site)
	err1 := opqBot.Start()
	if err1 != nil {
		log.Println(err.Error())
	}
	defer opqBot.Stop()
	err = opqBot.AddEvent(OPQBot.EventNameOnGroupMessage, func(botQQ int64, packet OPQBot.GroupMsgPack) {
		//log.Println(botQQ, packet.Content)
			if strings.HasPrefix(packet.Content, "禁言") {
				rt1, _ :=opqBot.GetUserInfo(packet.FromUserID)
				if rt1.Sex==1{
					c := strings.Fields(strings.TrimPrefix(packet.Content, "禁言"))
					if setjy(packet.FromGroupID, strconv.FormatInt(conf1.Qq, 10), conf1.Site, c[0], c[1]) {
						send2g(&opqBot, packet.FromUserID, "已禁言该成员")
					} else {
						send2g(&opqBot, packet.FromUserID, "禁言失败")
					}
				}else{
					send2g(&opqBot, packet.FromUserID, "你没有该权限！！！")
				}

			}
		if strings.HasPrefix(packet.Content,"设置公告"){
			c:=strings.Fields(strings.TrimPrefix(packet.Content,"设置公告"))
			if setgg(packet.FromGroupID,strconv.FormatInt(conf1.Qq, 10),conf1.Site,c[0],c[1]){
				send2g(&opqBot,packet.FromUserID,"设置公告成功")
			}else{
				send2g(&opqBot,packet.FromUserID,"设置公告失败")
			}
		}
	})
	err = opqBot.AddEvent(OPQBot.EventNameOnFriendMessage, func(botQQ int64, packet OPQBot.FriendMsgPack) {
		log.Println(botQQ, packet.Content)
	})

	err = opqBot.AddEvent(OPQBot.EventNameOnGroupShut, func(botQQ int64, packet OPQBot.GroupShutPack) {
		log.Println(botQQ, packet)
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
	for true {
		time.Sleep(1 * time.Hour)
	}
}
