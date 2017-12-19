package iwechat

import "github.com/chanxuehong/wechat.v2/mp/core"

var mux *core.ServeMux

func initMux() {

	mux = core.NewServeMux()
	mux.DefaultEventHandleFunc(eventHandler)
	mux.DefaultMsgHandleFunc(messageHandler)

}
