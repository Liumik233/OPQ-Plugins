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

var ver = "XrayManger_for_OPQ_ver.0.1a"
var surl string

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
	surl = "http://txyhk.liumik.tech:8031"
	fmt.Println(ver)
	fmt.Println("By Liumik")
	if !Exists("./config.json") {
		tmp := make(map[string]interface{})
		var qq int64
		var site []byte
		fmt.Println("\n请输入OPQ的Web地址: ")
		fmt.Scan(&site)
		fmt.Println("\n请输入Bot账号: ")
		fmt.Scan(&qq)
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
		Site string `Site`
		Qq   int64  `Qq`
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
	})
	err = opqBot.AddEvent(OPQBot.EventNameOnFriendMessage, func(botQQ int64, packet OPQBot.FriendMsgPack) {
		//log.Println(botQQ, packet.Content)
		if packet.FromUin == 253145061 {
			if strings.HasPrefix(packet.Content, "查看信息") {
				id, time := ck(strings.TrimPrefix(packet.Content, "查看信息"))
				send2p(&opqBot, packet.FromUin, "Group: A\nIP: txyhk.liumik.tech\nPort: 443\nTLS: 开启\nID:"+id+"\nEncryption: none\nNetwork: WebSocket host: txyhk.liumik.tech, path: /PudI4Hbh/\n"+"Vless链接：vless://"+id+"@txyhk.liumik.tech:443?encryption=none&security=tls&type=ws&&host=txyhk.liumik.tech&path=/PudI4Hbh/\n"+"到期时间："+time)
			}
			if strings.HasPrefix(packet.Content, "续期") {
				rt := strings.Fields(strings.TrimPrefix(packet.Content, "续期"))
				str := xq(rt[0], rt[1])
				send2p(&opqBot, packet.FromUin, str)
			}
			if strings.HasPrefix(packet.Content, "创建") {
				rt := strings.Fields(strings.TrimPrefix(packet.Content, "创建"))
				send2p(&opqBot, packet.FromUin, cuser(rt[0], rt[1]))
			}
			if strings.HasPrefix(packet.Content, "切换") {
				surl = strings.TrimPrefix(packet.Content, "切换")
				send2p(&opqBot, packet.FromUin, "切换成功")
			}
			if strings.HasPrefix(packet.Content, "服务器") {
				strings.TrimPrefix(packet.Content, "服务器")
				send2p(&opqBot, packet.FromUin, "当前服务器为"+surl)
			}
		}
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
