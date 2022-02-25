package znet

type Message struct {
	Head     uint16
	Sign     uint16
	SendTime uint32
	DataLen  uint16 //消息的长度
	Id       uint16 //消息的ID
	Data     []byte //消息的内容
}

//创建一个Message消息包
func NewMsgPackage(id uint16, data []byte) *Message {
	return &Message{
		DataLen: uint16(len(data)),
		Id:      id,
		Data:    data,
	}
}

//获取消息数据段长度
func (msg *Message) GetDataLen() uint16 {
	return msg.DataLen
}

//获取消息ID
func (msg *Message) GetMsgId() uint16 {
	return msg.Id
}

//获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}
func (msg *Message) GetSign() uint16 {
	return msg.Sign
}
func (msg *Message) GetSendTime() uint32 {
	return msg.SendTime
}

//设置消息数据段长度
func (msg *Message) SetDataLen(len uint16) {
	msg.DataLen = len
}

//设计消息ID
func (msg *Message) SetMsgId(msgId uint16) {
	msg.Id = msgId
}

//设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
