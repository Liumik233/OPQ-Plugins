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
	fmt.Println("Aria2M_for_OPQ_ver.0.1c")
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
	err = opqBot.AddEvent(OPQBot.EventNameOnGroupMessage, func(botQQ int64, packet OPQBot.GroupMsgPack) {
		log.Println(botQQ, packet.Content)
		fileinfo := struct {
			FileID   string `FileID`
			FileName string `FileName`
		}{}
		json.Unmarshal([]byte(packet.Content), &fileinfo)
		if strings.HasPrefix(packet.Content, "addurl_") {
			gid, err := Addurl(strings.Trim(packet.Content, "addurl_"), conf1.Url, conf1.Token)
			if err != nil {
				sent2g(&opqBot, packet.FromGroupID, "error:"+err.Error())
			} else {
				sent2g(&opqBot, packet.FromGroupID, "Successful,gid:"+gid)
			}
		}
		if strings.HasPrefix(packet.Content, "status_") {
			rsp, err := Filestatus(strings.Trim(packet.Content, "status_"), conf1.Url, conf1.Token)
			if err != nil {
				sent2g(&opqBot, packet.FromGroupID, "error:"+err.Error())
			} else {
				sent2g(&opqBot, packet.FromGroupID, rsp)
			}
		}
		if strings.HasPrefix(fileinfo.FileName, "addbt_") {
			urlt := Getfile(packet.FromGroupID, fileinfo.FileID, strconv.FormatInt(conf1.Qq, 10), conf1.Site)
			gid, err := Addbt(urlt, conf1.Url, conf1.Token)
			if err != nil {
				sent2g(&opqBot, packet.FromGroupID, "error:"+err.Error())
			} else {
				sent2g(&opqBot, packet.FromGroupID, "Successful,gid:"+gid)
			}
		}
		if strings.HasPrefix(packet.Content, "stop_") {
			err := Stop(strings.TrimPrefix(packet.Content, "stop_"), conf1.Url, conf1.Token)
			if err != nil {
				sent2g(&opqBot, packet.FromGroupID, "error:"+err.Error())
			} else {
				sent2g(&opqBot, packet.FromGroupID, "Successful")
			}
		}
		if strings.HasPrefix(packet.Content, "start_") {
			err := Start(strings.TrimPrefix(packet.Content, "start_"), conf1.Url, conf1.Token)
			if err != nil {
				sent2g(&opqBot, packet.FromGroupID, "error:"+err.Error())
			} else {
				sent2g(&opqBot, packet.FromGroupID, "Successful")
			}

		}
		if strings.HasPrefix(packet.Content, "del_") {
			err := Del(strings.TrimPrefix(packet.Content, "del_"), conf1.Url, conf1.Token)
			if err != nil {
				sent2g(&opqBot, packet.FromGroupID, "error:"+err.Error())
			} else {
				sent2g(&opqBot, packet.FromGroupID, "Successful")
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
	time.Sleep(1 * time.Hour)
}
