package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"redpacket-sweep/conf"
	"redpacket-sweep/msg"
)

func init() {
	handler(&msg.C2SL_StartMatch{}, handleStartMatch)
	handler(&msg.C2SL_ExitRoom{}, handleExitRoom)
}

func handleStartMatch(args []interface{}) {
	if len(args) != 2 {
		log.Error("args invalid")
		return
	}

	m := args[0].(*msg.C2SL_StartMatch)
	a := args[1].(gate.Agent)

	userID := a.UserData().(*AgentInfo).userID
	user := userIDUsers[userID]
	if user == nil {
		a.WriteMsg(&msg.S2C_Close{})
		a.Close()
		return
	}

	roomRule := roomRule{}
	roomRule.RoomMetaData = (*conf.GetCfgMatchRoomMateData())[m.ItemType].RoomMetaData
	user.createOrMatchingRoom(&roomRule)

	log.Debug("【房间数】：%v， 房间指针：%v", len(roomNumberRooms), roomNumberRooms)
}

func handleExitRoom(args []interface{}) {
	if len(args) != 2{
		return
	}

	_ = args[0].(*msg.C2SL_ExitRoom)
	a := args[1].(gate.Agent)

	user := userIDUsers[a.UserData().(*AgentInfo).userID]
	if user == nil {
		return
	}

	user.exitRoom()
}