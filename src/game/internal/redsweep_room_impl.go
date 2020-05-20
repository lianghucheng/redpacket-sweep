package internal

import (
	"github.com/name5566/leaf/log"
	"redpacket-sweep/common"
	"redpacket-sweep/msg"
)

func (room *RedSweepRoom)Exit(userID int) {
	player := room.userIDPlayerDatas[userID]

	//todo:在游戏中...
	//todo:站起...
	log.Debug("userID: %v 退出 %v 房: %v, 类型: %v, 索引: %v", userID, roomTypeString(room.roomType), room.number, room.roomType, room.roomIndex)
	_ = player
	//player.writeMsg(&msg.SL2C_StandUp{})
}

func (room *RedSweepRoom) Enter(user *User) {
	if player, ok := room.userIDPlayerDatas[user.userID()]; ok {
		player.user().WriteMsg(&msg.SL2C_EnterRoom{ // 不能使用 player.writeMsg
			Error:         msg.SL2C_EnterRoom_OK,
			Desc:          room.desc,
		})
		log.Debug("userID: %v 重连进入 %v房: %v, 类型: %v", user.userID(), roomTypeString(room.roomType), room.number, room.roomType)
		return
	}
	if user.systemOff() {
		return
	}

	room.SitDown(user.userID())
	playerData := room.userIDPlayerDatas[user.userID()]
	userIDRooms[user.userID()] = room
	user.WriteMsg(&msg.SL2C_EnterRoom{
		Error:         	msg.SL2C_EnterRoom_OK,
		Desc:          	room.desc,
		Num: 			int64(len(room.userIDPlayerDatas)),
		Headimgurl:    	user.headimgurl(),
		AccountID:		user.accID(),
		CarryCoin:		common.TranferChipRate(playerData.carryCoin),
		WinCoin:		common.TranferChipRate(playerData.winCoin),
	})

	log.Debug("userID: %v 进入 %v房: %v, 房主: %v, 类型: %v, 索引: %v, 位置: %v", user.userID(), roomTypeString(room.roomType), room.number, room.ownerUserID, room.roomType, room.roomIndex)
	return
}

func (room *RedSweepRoom) SitDown(userID int) (bool, int) {
	player := room.userIDPlayerDatas[userID]
	if player == nil {
		player = newPlayerData(userID)
		room.userIDPlayerDatas[userID] = player
	}
	room.broadcastExclude(&msg.SL2C_SitDown{
		Error:      msg.SL2C_SitDown_OK,
		AccountID:  player.accID(),
		Nickname:   player.nickname(),
		UserID:		player.userID,
		Num:		int64(len(room.userIDPlayerDatas)),
	}, player.userID)
	log.Debug("userID: %v 坐下 %v房: %v, 类型: %v, 索引: %v, 位置: %v", userID, roomTypeString(room.roomType), room.number, room.roomType, room.roomIndex)
	return true, player.userID
}