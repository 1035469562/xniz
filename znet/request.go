package znet

import "xniz/ziface"

type Request struct {
	conn ziface.IConnection
	data []byte
}

//获取请求连接信息
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.data
}
