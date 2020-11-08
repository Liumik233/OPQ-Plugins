package main

import (
	"context"
	"github.com/zyxar/argo/rpc"
	"log"
	"os"
	"time"
)

func Connaria2(url string, token string) rpc.Client {
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

func Addurl(url string, aria2 rpc.Client) (string, error) {
	gid, err := aria2.AddURI(url)
	if err != nil {
		return gid, err
		log.Println(err)
	}
	return gid, nil
}

func Addbt(url string, aria2 rpc.Client) (string, error) {
	/*cmd := exec.Command("wget", url, "-O ./tmp/tmp.torrent")
	cmd.Run()
	gid, err := aria2.AddTorrent("./tmp/tmp.torrent")
	if err != nil {
		return "err", err
	}*/
	log.Println(url)
	return gid, nil
}

func Filestatus(gid string, aria2 rpc.Client) (string, error) {
	rsp, err := aria2.TellStatus(gid)
	if err != nil {
		return "err", err
		log.Println(err)
	}
	if rsp.Status == "active" {
		return "状态：" + rsp.Status + "\n下载速度：" + rsp.DownloadSpeed + "\n下载进度：" + rsp.CompletedLength + "/" + rsp.TotalLength, err
	} else {
		return "状态：" + rsp.Status, err
	}
}

func stop(gid string, aria2 rpc.Client) (string, error) {
	rsp, err := aria2.Pause(gid)
	if err != nil {
		return "err", err
		log.Println(err)
	}
	return rsp, nil
}
func start(gid string, aria2 rpc.Client) (string, error) {
	rsp, err := aria2.Unpause(gid)
	if err != nil {
		return "err", err
		log.Println(err)
	}
	return rsp, nil
}
