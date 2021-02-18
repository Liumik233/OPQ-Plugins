package main

import (
	"context"
	"github.com/zyxar/argo/rpc"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type aria2c struct {
	url, token *string
	a          rpc.Client
}

func (a *aria2c) Connaria2() error {
	ctx := context.Background()
	var no rpc.Notifier
	rsp, err := rpc.New(ctx, *a.url, *a.token, time.Second*5, no)
	if err != nil {
		return err
		os.Exit(1)
	}
	a.a = rsp
	ver, err := rsp.GetVersion()
	log.Println("connect successful,ver:", ver.Version)
	return nil
}

func (a *aria2c) Addurl(url1 string) (string, error) {
	url := make([]string, 1)
	url[0] = url1
	gid, err := a.a.AddURI(url)
	if err != nil {
		return gid, err
		log.Println(err)
	}
	return gid, nil
}

func (a *aria2c) Addbt(url string) (string, error) {
	cmd := exec.Command("wget", url, "-O", "./tmp/tmp.torrent")
	cmd.Run()
	gid, err := a.a.AddTorrent("./tmp/tmp.torrent")
	if err != nil {
		return "err", err
	}
	os.Remove("./tmp/tmp.torrent")
	return gid, nil
}

func (a *aria2c) Filestatus(gid string) (string, error) {
	rsp, err := a.a.TellStatus(gid)
	if err != nil {
		return "err", err
		log.Println(err)
	}
	spi, err := strconv.ParseInt(rsp.DownloadSpeed, 10, 64)
	toi, err := strconv.ParseFloat(rsp.TotalLength, 64)
	cpi, err := strconv.ParseFloat(rsp.CompletedLength, 64)

	if rsp.Status == "active" {
		return "文件名：" + strings.Trim(rsp.Files[0].Path, rsp.Dir) + "\n下载状态：" + rsp.Status + "\n下载速度：" + strconv.FormatInt(spi/1024, 10) + "KB/s\n下载进度：" + strconv.FormatInt(int64(cpi/toi*100), 10) + "%", err
	} else {
		return "文件名：" + strings.Trim(rsp.Files[0].Path, rsp.Dir) + "\n下载状态：" + rsp.Status, err
	}
}

func (a *aria2c) Stop(gid string) error {
	_, err := a.a.Pause(gid)
	return err
}

func (a *aria2c) Start(gid string) error {
	_, err := a.a.Unpause(gid)
	return err
}

func (a *aria2c) Del(gid string) error {
	_, err := a.a.Remove(gid)
	return err
}
func (a *aria2c) Closearia2() {
	a.a.Close()
}
