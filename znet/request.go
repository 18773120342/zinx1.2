package znet

import (
	"github.com/18773120342/zinx1.2/ziface"
)

type Request struct {
	Conn ziface.IConnection //已经和客户端建立好的 链接
	Msg  ziface.IMessage    //客户端请求的数据
}

//获取请求连接信息
func (r *Request) GetConnection() ziface.IConnection {
	return r.Conn
}

//获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.Msg.GetData()
}

//获取请求的消息的ID
func (r *Request) GetMsgID() uint16 {
	return r.Msg.GetMsgId()
}

func (r *Request) GetSign() uint16 {
	return r.Msg.GetSign()
}

func (r *Request) GetSendTime() uint32 {
	return r.Msg.GetSendTime()
}
