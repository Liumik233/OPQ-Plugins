package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func post1(url1 string, key string) string {
	type rejsont struct {
		Id   string `json:"id"`
		Ourl string `json:"output_url"`
	}
	var rejson1 rejsont
	client := &http.Client{}
	data := url.Values{"image": {url1}} //字符串表单
	req, _ := http.NewRequest("POST", "https://api.deepai.org/api/waifu2x", strings.NewReader(data.Encode()))
	//fmt.Println(strings.NewReader(data.Encode()))
	req.Header.Add("api-key", key)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") //指定Content-type
	resp, _ := client.Do(req)                                           //发送post请求
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body) //读取body
	json.Unmarshal(body, &rejson1)       //json to struct
	return rejson1.Ourl                  //返回处理后的图片url
}

func post2(file string, key string) (string, int) {
	type rejsont struct {
		Id   string `json:"id"`
		Ourl string `json:"output_url"`
		Err  string `json:error`
	}
	var rejson1 rejsont
	client := &http.Client{}
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, _ := bodyWriter.CreateFormFile("image", file) //文件表单
	fh, _ := os.Open(file)                                    //读取文件流
	_, _ = io.Copy(fileWriter, fh)                            //复制文件流
	contentType := bodyWriter.FormDataContentType()           //获取Content-type
	bodyWriter.Close()
	req, err := http.NewRequest("POST", "https://api.deepai.org/api/waifu2x", bodyBuf)
	if err != nil {
		// handle error
	}
	req.Header.Add("api-key", key)
	req.Header.Add("Content-Type", contentType) //指定Content-type
	resp, _ := client.Do(req)                   //发送post请求
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body) //读取body
	json.Unmarshal(body, &rejson1)       //json to struct
	if len(rejson1.Err) == 0 {
		return rejson1.Ourl, 0
	} else {
		return rejson1.Err, 1
	}
}
