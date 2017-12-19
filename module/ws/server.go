package ws

var WsServer *wsServer

func init() {
	WsServer = new(wsServer)
}

type Message struct {
	ToOpenId string
	Mess     string
}
type wsServer struct {
}

//发送ws消息
func (self *wsServer) Send(message interface{}) {

}
