package main

import (
	"encoding/json"
	"fmt"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"io/ioutil"
	"iotqq/model"
	"log"
	"os"
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
	fmt.Println("BigPic_for_OPQ_ver.0.01a")
	fmt.Println("作者:Liumik")
	if !Exists("./config.json") {
		tmp := make(map[string]interface{})
		var url string
		var site string
		var port int
		var token string
		port = 8888
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
		tmp["Site"] = site
		tmp["Port"] = port
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
		Port  int    `Port`
		Qq    string `Qq`
		Url   string `Url`
		Token string `Token`
	}{}
	json.Unmarshal(cb, &conf1)
	qq = conf1.Qq
	aria2 := Connaria2(conf1.Url, conf1.Token)
	defer aria2.Close()
	runtime.GOMAXPROCS(runtime.NumCPU())
	BotUrl = conf1.Site + ":" + strconv.Itoa(conf1.Port)
	iotqq.Set(BotUrl, qq)
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(conf1.Site, conf1.Port, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatal(err)
	}

	err = c.On("OnGroupMsgs", func(h *gosocketio.Channel, args iotqq.Message) {
		fileinfo := struct {
			FileID   string `FileID`
			FileName string `FileName`
		}{}
		var mess iotqq.Data = args.CurrentPacket.Data
		/*
			mess.Content 消息内容 string
			mess.FromGroupID 来源QQ群 int
			mess.FromUserID 来源QQ int64
			mess.iotqqType 消息类型 string
		*/
		json.Unmarshal([]byte(mess.Content), &fileinfo)
		log.Println("群聊消息: ", mess.FromNickName+"<"+strconv.FormatInt(mess.FromUserID, 10)+">: "+mess.Content)
		if strings.HasPrefix(mess.Content, "addurl") {
			gid, err := Addurl(strings.Trim(mess.Content, "addurl"), aria2)
			if err != nil {
				iotqq.Send(mess.FromGroupID, 2, "error:"+err.Error())
			} else {
				iotqq.Send(mess.FromGroupID, 2, "Successful,gid:"+gid)
			}
		}
		if strings.HasPrefix(mess.Content, "status") {
			rsp, err := Filestatus(strings.Trim(mess.Content, "status"), aria2)
			if err != nil {
				iotqq.Send(mess.FromGroupID, 2, "error:"+err.Error())
			} else {
				iotqq.Send(mess.FromGroupID, 2, rsp)
			}
		}
		if strings.HasPrefix(fileinfo.FileName, "addbt") {
			rsp := iotqq.Getfile(mess.FromGroupID, fileinfo.FileID)
			gid, err := Addbt(rsp, aria2)
			if err != nil {
				iotqq.Send(mess.FromGroupID, 2, "error:"+err.Error())
			} else {
				iotqq.Send(mess.FromGroupID, 2, "Successful,gid:"+gid)
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
