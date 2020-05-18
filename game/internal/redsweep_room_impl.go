package internal

import "github.com/name5566/leaf/log"

func (room *RedSweepRoom)Exit(userID int) {
	player := room.userIDPlayerDatas[userID]

	//todo:在游戏中...
	//todo:站起...
	log.Debug("userID: %v 退出 %v 房: %v, 类型: %v, 索引: %v", userID, roomTypeString(room.roomType), room.number, room.roomType, room.roomIndex)
	_ = player
	//player.writeMsg(&msg.SL2C_StandUp{})
}