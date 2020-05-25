package internal

import (
	"github.com/name5566/leaf/gate"
	"redpacket-sweep/msg"
)

func init() {
	handler(&msg.C2SL_SendRedPacket{}, handleSendRedpacket)
	handler(&msg.C2SL_TakenRedPacket{}, handleTakenRedPacket)
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
		if room, ok := userIDRooms[user.userID()]; ok {
			user.sendRedPacket(room ,&m.RedPacketMetaData)
		}
	}
}

func handleTakenRedPacket(args []interface{}) {
	if len(args) != 2 {
		return
	}

	_ = args[0].(*msg.C2SL_TakenRedPacket)
	a := args[1].(gate.Agent)

	user := userIDUsers[a.UserData().(*AgentInfo).userID]

	if user == nil {
		a.Close()
		return
	}

	if room, ok := userIDRooms[user.userID()]; ok {
		user.takenRedPacket(room)
	}
}