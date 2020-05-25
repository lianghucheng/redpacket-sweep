package internal

import "github.com/name5566/leaf/timer"

type RedSweepRoom struct {
	room
	userIDPlayerDatas			map[int]*PlayerData
	rule 						*roomRule
	redPacketQueue				[]*RedPacket
	allPacketQueue				[]*RedPacket
	currentRedpacket 			*RedPacket
	takenPlayer 				[]*PlayerData

	systemSendRedPacketTimer	*timer.Timer
}

type RedPacket struct {
	userID 							int
	isOpen							bool
	Total							int64
	Quota							int64
	TakenNum						int
	Boom 							int
}

func newRedSweepRoom(rule *roomRule) *RedSweepRoom {
	room := new(RedSweepRoom)
	room.roomType = rule.RoomType
	room.roomIndex = rule.RoomIndex
	room.userIPAddrs = make(map[string]bool)
	room.userIDPlayerDatas = make(map[int]*PlayerData)
	room.rule = rule

	return room
}

//机器人人数
func (room *RedSweepRoom) robotNum() int {
	count := 0
	for _, value := range room.userIDPlayerDatas {
		if value.isRobot() {
			count++
		}
	}
	return count
}

func (room *RedSweepRoom) broadcastExclude(msg interface{}, userID int) {
	for _, player := range room.userIDPlayerDatas {
		if player.userID == userID { // 指定位置的玩家不发消息
			continue
		}
		player.writeMsg(msg)
	}
	for _, player := range room.userIDPlayerDatas {
		if player.userID == -1 { // 给站起围观的玩家发消息
			player.writeMsg(msg)
		}
	}
}

func (room *RedSweepRoom) broadcast(msg interface{}) {
	for _, player := range room.userIDPlayerDatas {
		player.writeMsg(msg)
	}
	for _, player := range room.userIDPlayerDatas {
		if player.userID == -1 { // 给站起围观的玩家发消息
			player.writeMsg(msg)
		}
	}
}