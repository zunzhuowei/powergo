package main

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"net"
	"server/msg"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3653")
	if err != nil {
		panic(err)
	}
	/*

		// Hello 消息（JSON 格式）
		// 对应游戏服务器 Hello 消息结构体
		data := []byte(`{
			"Hello": {
				"Name": "leaf"
			}
		}`)


		// len + data
			m := make([]byte, 2+len(data))

			// 默认使用大端序
			binary.BigEndian.PutUint16(m, uint16(len(data)))

			copy(m[2:], data)

			// 发送消息 json 消息
			conn.Write(m)

			fmt.Println("00000000000000000000000000000 ----::")
	*/

	protoBuf(conn)

}

func protoBuf(conn net.Conn) {
	// 发送 protoBuf 消息

	// 创建一个消息 Test
	test := &msg.Hello{}
	test.Name = proto.String("Jeff")

	// 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	//-------------------------------
	//| len | id | protobuf message |
	//-------------------------------

	protoBuf := make([]byte, 4+len(data))
	binary.BigEndian.PutUint16(protoBuf, uint16(len(data)+2)) // len = (id len) + (data len)
	binary.BigEndian.PutUint16(protoBuf[2:], 0)               // id
	copy(protoBuf[4:], data)                                  // data
	conn.Write(protoBuf)
}

/*

在 Leaf 中，默认的 Protobuf Processor 将一个完整的 Protobuf 消息定义为如下格式：

-------------------------
| id | protobuf message |
-------------------------
其中 id 为 2 个字节。如果你选择使用 TCP 协议时，在网络中传输的消息格式如下：

-------------------------------
| len | id | protobuf message |
-------------------------------
如果你选择使用 WebSocket 协议时，发送的消息格式如下：

-------------------------
| id | protobuf message |
-------------------------

*/

/*

 现在我们编写一个小程序：

package main

import (
    "log"
    // 辅助库
    "github.com/golang/protobuf/proto"
    // hello.pb.go 的路径
    "example"
)

func main() {
    // 创建一个消息 Test
    test := &example.Test{
        // 使用辅助函数设置域的值
        Label: proto.String("hello"),
        Type:  proto.Int32(17),
        Optionalgroup: &example.Test_OptionalGroup{
            RequiredField: proto.String("good bye"),
        },
    }

    // 进行编码
    data, err := proto.Marshal(test)
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }

    // 进行解码
    newTest := &example.Test{}
    err = proto.Unmarshal(data, newTest)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }

    // 测试结果
    if test.GetLabel() != newTest.GetLabel() {
        log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
    }
}
*/
