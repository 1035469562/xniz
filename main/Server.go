package main

import (
	"fmt"
	"xniz/ziface"
	"xniz/znet"
)

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter //一定要先基础BaseRouter
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("Hello Zinx Router V0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("[max]")
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Serve()
	////创建socket TCP Server
	//listener, err := net.Listen("tcp", "127.0.0.1:7777")
	//if err != nil {
	//	fmt.Println("server listen err:", err)
	//	return
	//}
	//
	////创建服务器gotoutine，负责从客户端goroutine读取粘包的数据，然后进行解析
	//
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		fmt.Println("server accept err:", err)
	//	}
	//	go func(conn net.Conn) {
	//		dp := znet.NewDataPack()
	//		for {
	//			headData := make([]byte, dp.GetHeadLen())
	//			_, err := io.ReadFull(conn, headData)
	//			if err != nil {
	//				fmt.Println("read head error")
	//				break
	//			}
	//			msgHead, err := dp.Unpack(headData)
	//			if err != nil {
	//				fmt.Println("server unpack err:", err)
	//				return
	//			}
	//			if msgHead.GetDataLen() > 0 {
	//				//msg 是有data数据的，需要再次读取data数据
	//				msg := msgHead.(*znet.Message)
	//				msg.Data = make([]byte, msg.GetDataLen())
	//
	//				//根据dataLen从io中读取字节流
	//				_, err := io.ReadFull(conn, msg.Data)
	//				if err != nil {
	//					fmt.Println("server unpack data err:", err)
	//					return
	//				}
	//				fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
	//			}
	//		}
	//	}(conn)
	//}

}
