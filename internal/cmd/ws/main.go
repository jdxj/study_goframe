package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"

	v1 "study_goframe/api/pb_hello/v1"
)

var ctx = gctx.New()

func main() {
	s := g.Server()
	s.BindHandler("/ws", func(r *ghttp.Request) {
		ws, err := r.WebSocket()
		if err != nil {
			glog.Error(ctx, err)
			r.Exit()
		}
		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
					return
				}
				glog.Error(ctx, err)
				return
			}
			req := v1.HelloReq{}
			err = proto.Unmarshal(msg, &req)
			if err != nil {
				glog.Error(ctx, err)
				return
			}

			res := v1.HelloRes{Reply: "hello " + req.GetName() + "!"}
			resBytes, err := proto.Marshal(&res)
			if err != nil {
				glog.Error(ctx, err)
				return
			}
			if err = ws.WriteMessage(websocket.BinaryMessage, resBytes); err != nil {
				glog.Error(ctx, err)
				return
			}
		}
	})
	s.SetPort(49152)
	s.Run()
}
