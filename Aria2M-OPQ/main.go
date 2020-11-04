package main

import (
	"fmt"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"iotqq/model"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var BotUrl, qq string
var conf iotqq.Conf
var zanok, qd []int64

func periodlycall(d time.Duration, f func()) {
	for x := range time.Tick(d) {
		f()
		log.Println(x)
	}
}
func resetzan() {

	m1 := len(zanok)
	for m := 0; m < m1; m++ {
		i := 0
		zanok = append(zanok[:i], zanok[i+1:]...)
	}
	m2 := len(qd)
	for m := 0; m < m2; m++ {
		i := 0
		qd = append(qd[:i], qd[i+1:]...)
	}
}
func SendJoin(c *gosocketio.Client) {
	log.Println("获取QQ号连接")
	result, err := c.Ack("GetWebConn", qq, time.Second*5)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("emit", result)
	}
}
func main() {
	var url string
	var site string
	var port int
	var token string
	port = 8888
	fmt.Println("BigPic_for_OPQ_ver.0.01a")
	fmt.Println("作者:Liumik")
	fmt.Println("\n请输入OPQ的Web地址(无需http://和端口): ")
	fmt.Scan(&site)
	fmt.Println("\n请输入OPQ的端口号: ")
	fmt.Scan(&port)
	fmt.Println("\n请输入QQ机器人账号: ")
	fmt.Scan(&qq)
	fmt.Println("\n请输入url")
	fmt.Scan(&url)
	fmt.Println("请输入token")
	fmt.Scan(&token)
	aria2 := connaria2(url, token)
	runtime.GOMAXPROCS(runtime.NumCPU())
	BotUrl = site + ":" + strconv.Itoa(port)
	iotqq.Set(BotUrl, qq)
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(site, port, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatal(err)
	}

	err = c.On("OnGroupMsgs", func(h *gosocketio.Channel, args iotqq.Message) {
		var mess iotqq.Data = args.CurrentPacket.Data
		/*
			mess.Content 消息内容 string
			mess.FromGroupID 来源QQ群 int
			mess.FromUserID 来源QQ int64
			mess.iotqqType 消息类型 string
		*/
		log.Println("群聊消息: ", mess.FromNickName+"<"+strconv.FormatInt(mess.FromUserID, 10)+">: "+mess.Content)
		if strings.HasPrefix(mess.Content, "addurl") {
			gid, err := addurl(strings.Trim(mess.Content, "addurl"), aria2)
			if err != nil {
				iotqq.Send(mess.FromGroupID, 2, "gid:"+gid+",error:"+err.Error())
			} else {
				iotqq.Send(mess.FromGroupID, 2, "Successful,gid:"+gid)
			}
		}
		if strings.HasPrefix(mess.Content, "status") {
			rsp, err := filestatus(strings.Trim(mess.Content, "status"), aria2)
			if err != nil {
				iotqq.Send(mess.FromGroupID, 2, "error:"+err.Error())
			} else {
				iotqq.Send(mess.FromGroupID, 2, rsp)
			}
		}

	})
	if err != nil {
		log.Fatal(err)
	}
	err = c.On("OnFriendMsgs", func(h *gosocketio.Channel, args iotqq.Message) {
		log.Println("私聊消息: ", args.CurrentPacket.Data.Content)

	})
	if err != nil {
		log.Fatal(err)
	}
	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}
	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("连接成功")
	})
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	go SendJoin(c)
	periodlycall(24*time.Hour, resetzan)
home:
	time.Sleep(600 * time.Second)
	SendJoin(c)
	goto home
	log.Println(" [x] Complete")
}
