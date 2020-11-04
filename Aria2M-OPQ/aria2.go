package main

import (
	"context"
	"github.com/zyxar/argo/rpc"
	"log"
	"os"
	"strconv"
	"time"
)

func connaria2(url string, token string) rpc.Client {
	ctx := context.Background()
	var no rpc.Notifier
	rsp, err := rpc.New(ctx, url, token, time.Second*30, no)
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

func addbt(url string, aria2 rpc.Client) (string, error) {
	gid, err := aria2.AddTorrent("./tmp/tmp.torrent")
	if err != nil {
		return gid, err
	}
	return gid, nil
}

func filestatus(gid string, aria2 rpc.Client) (string, error) {
	rsp, err := aria2.TellStatus(gid)
	if err != nil {
		return "err", err
		log.Println(err)
	}
	ic, _ := strconv.Atoi(rsp.CompletedLength)
	it, _ := strconv.Atoi(rsp.TotalLength)
	return "下载速度：" + rsp.DownloadSpeed + "\n下载进度：" + strconv.Itoa(ic/it*100) + "%", nil
}
