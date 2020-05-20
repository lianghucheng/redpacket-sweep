package internal

import "redpacket-sweep/msg"

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
		if user.isRobot() && room.robotNum() >= rule.RedPacketNum {
			return false
		}
		if *room.metaData == *rule {
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