package internal

import (
	"github.com/name5566/leaf/gate"
	"redpacket-sweep/msg"
)

func init() {
	handler(&msg.C2SL_SendRedPacket{}, handleSendRedpacket)
}

func handleSendRedpacket(args []interface{}) {
	if len(args) != 2 {
		return
	}

	m := args[0].(*msg.C2SL_SendRedPacket)
	a := args[1].(gate.Agent)

	if user, ok := userIDUsers[a.UserData().(*AgentInfo).userID]; !ok {
		a.Close()
		return
	} else {
		user.sendRedPacket(&m.RedPacketMetaData)
	}

}