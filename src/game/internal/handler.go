package internal

import (
	"github.com/name5566/leaf/gate"
	"math/rand"
	"redpacket-sweep/msg"
	"reflect"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	handler(&msg.C2S_Heartbeat{}, handleHeartbeat)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHeartbeat(args []interface{}) {
	a := args[1].(gate.Agent) // 消息的发送者

	if user, ok := userIDUsers[a.UserData().(*AgentInfo).userID]; ok {
		user.heartbeatStop = false
	}
}