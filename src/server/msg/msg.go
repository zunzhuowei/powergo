package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
//var Processor = json.NewProcessor()
//
//func init() {
//	// 这里我们注册了一个 JSON 消息 Hello
//	Processor.Register(&Hello{})
//}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
//type Hello struct {
//	Name string
//}

//////////////////////////////////////

// 使用 Protobuf 消息处理器
var Processor = protobuf.NewProcessor()

func init() {
	// 这里我们注册了消息 Hello
	Processor.Register(&Hello{})
}
