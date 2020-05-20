package internal

import "redpacket-sweep/metadata"

type RedSweepRoom struct {
	room
	userIDPlayerDatas	map[int]*PlayerData
	metaData 			*roomRule
	redPacketQueue		[]*RedPacket
	allPacketQueue		[]*RedPacket
}

type RedPacket struct {
	userID 							int
	isOpen							bool
	metadata.RedPacketMetaData
}

func newRedSweepRoom(metaData *roomRule) *RedSweepRoom {
	room := new(RedSweepRoom)
	room.roomType = metaData.RoomType
	room.roomIndex = metaData.RoomIndex
	room.userIPAddrs = make(map[string]bool)
	room.userIDPlayerDatas = make(map[int]*PlayerData)

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