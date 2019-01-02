package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
	"server/msg/resp"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.Login{}, handlerLogin)
}

func handlerLogin(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Login)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("login %v", m)

	a.WriteMsg(&resp.LoginResp{
		Su:  proto.Bool(true),
		Exp: proto.Int32(3600),
	})
}
