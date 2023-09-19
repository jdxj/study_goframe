package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"

	v1 "study_goframe/api/pb_hello/v1"
)

func main() {
	client := gclient.NewWebSocket()
	client.HandshakeTimeout = time.Second    // 设置超时时间
	client.Proxy = http.ProxyFromEnvironment // 设置代理
	client.TLSClientConfig = &tls.Config{}   // 设置 tls 配置

	conn, _, err := client.Dial("ws://127.0.0.1:49152/ws", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	req := &v1.HelloReq{
		Name: "zhang san",
	}
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	err = conn.WriteMessage(websocket.BinaryMessage, reqBytes)
	if err != nil {
		panic(err)
	}

	_, data, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	res := &v1.HelloRes{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.GetReply())
}
