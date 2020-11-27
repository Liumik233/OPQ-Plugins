package main

import (
	"github.com/everpcpc/pixiv"
	iotqq "iotqq/model"
)

func checkpic(id uint64, app *pixiv.AppPixivAPI, mess iotqq.Data) {
	app.Download(id, "tmp/")
}
