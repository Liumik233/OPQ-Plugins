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
	fmt.Println("Aria2M_for_OPQ_ver.0.2a")
	fmt.Println("By Liumik")
	if !Exists("./config.json") {
		tmp := make(map[string]interface{})
		var apikey string
		var qq int64
		var site string
		fmt.Println("\n请输入OPQ的Web地址: ")
		fmt.Scan(&site)
		fmt.Println("\n请输入Bot账号: ")
		fmt.Scan(&qq)
		fmt.Println("\n请输入url")
		fmt.Scan(&apikey)
		tmp["Site"] = site
		tmp["Qq"] = qq
		tmp["Apikey"] = apikey
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
		Site   string `Site`
		Qq     int64  `Qq`
		Apikey string `Apikey`
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
		type pict struct {
			Url string `json:"Url"`
		}
		type pic struct {
			Content  string `json:"Content"`
			GroupPic []pict `json:"GroupPic"`
		}
		pic1 := pic{}
		json.Unmarshal([]byte(packet.Content), &pic1)
		if strings.HasPrefix(pic1.Content, "bigpic") {
			send2gp(&opqBot, packet.FromGroupID, "放大成功[PICFLAG]", post1(pic1.GroupPic[0].Url, conf1.Apikey))
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
