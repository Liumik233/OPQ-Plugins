package main

import (
	"context"
	"github.com/zyxar/argo/rpc"
	"log"
	"os"
	"time"
)

func connaria2(url string, token string) rpc.Client {
	ctx := context.Background()
	var no rpc.Notifier
	rsp, err := rpc.New(ctx, url, token, time.Second*3000, no)
	if err != nil {
		log.Println("err:", err)
		os.Exit(1)
	}
	ver, err := rsp.GetVersion()
	log.Println("connect successful,ver:", ver.Version)
	return rsp
}

func addurl(url string, aria2 rpc.Client) (string, error) {
	gid, err := aria2.AddURI(url)
	if err != nil {
		return gid, err
		log.Println(err)
	}
	return gid, nil
}

/*func addbt(url string, aria2 rpc.Client) (string, error) {

	gid,err:=aria2.AddTorrent("./tmp/tmp.torrent")
	if err!=nil{
		return gid,err
	}
	return gid,nil
}*/
//因api限制，暂时无法实现
func file(gid string, aria2 rpc.Client) ([]rpc.FileInfo, error) {
	rsp, err := aria2.GetFiles(gid)
	if err != nil {
		return rsp, err
		log.Println(err)
	}
	return rsp, nil
}
