package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	var loggers = log.Default()
	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// 准备入参
	// 用户实体
	type User struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	user := User{
		Id:   1,
		Name: "bob",
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		loggers.Printf("序列化用户失败！err:%+v", err)
		return
	}

	// 创建客户端
	client := &http.Client{}
	// 创建请求
	request, err := http.NewRequest("POST", "http://192.168.100.103:9090/api/v1/query_range?query=2022-10-18&start=17:58:00Z\\&end=%sT%s:15Z\\&step=30d", bytes.NewReader(jsonBytes))
	if err != nil {
		loggers.Printf("创建请求失败！err:%+v", err)
		return
	}
	// 设置请求头
	request.Header.Add("Cookie", "123")
	request.Header.Add("Content-Type", "application/json;charset=utf-8")
	request.Header.Add("Token", "456")

	// 设置1分钟的超时时间
	client.Timeout = 1 * time.Minute

	// 发起请求
	resp, err := client.Do(request)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		loggers.Printf("读取Body失败 error: %+v", err)
		return
	}
	loggers.Println(string(body))

}
