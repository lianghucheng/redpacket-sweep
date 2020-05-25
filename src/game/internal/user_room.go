package internal

import (
	"github.com/name5566/leaf/log"
	"redpacket-sweep/conf"
	"redpacket-sweep/msg"
)

func (user *User) exitRoom() {
	userID := user.userID()
	if room, ok := userIDRooms[userID]; ok {
		room.Exit(userID)
	}
}

func (user *User)createOrMatchingRoom(rule *roomRule) {


	if user.enterMatchingRoom(rule) {
		return
	}

	user.createRoom(rule)
}

func (user *User)enterMatchingRoom(rule *roomRule) bool {
	for _, room :=range roomNumberRooms {
		player := int64(len(room.userIDPlayerDatas) - room.robotNum())
		rateNum := rule.MaxPlayerNum * int64(conf.Server.SegmentRoomRate) / 100
		if !user.isRobot() && rateNum <= player {
			continue
		}
		if user.isRobot() && int64(room.robotNum()) >= rule.MaxPlayerNum / 3 {
			return false
		}
		log.Debug("【metaData】:%v   【rule】:%v", *room, *rule)
		if *room.rule == *rule {
			room.Enter(user)
			return true
		}
	}
	return false
}

func (user *User)createRoom(rule *roomRule) {
	if user.systemOff() {
		return
	}
	//usdt过滤

	number := getRoomNumber()
	if _, ok := roomNumberRooms[number]; ok {
		user.WriteMsg(&msg.S2C_Close{})
		user.Close()
		return
	}
	userID := user.userID()
	room := newRedSweepRoom(rule)
	room.number = number
	room.ownerUserID = userID
	room.creatorUserID = room.ownerUserID

	roomNumberRooms[number] = room
	room.Enter(user)
}

func (user *User)reconnect(room *RedSweepRoom) {
	user.WriteMsg(&msg.SL2C_EnterRoom{ // 不能使用 player.writeMsg
		Error:         msg.SL2C_EnterRoom_OK,
		Desc:          room.desc,
	})
	log.Debug("userID: %v 重连进入 %v房: %v, 类型: %v", user.userID(), roomTypeString(room.roomType), room.number, room.roomType)
}