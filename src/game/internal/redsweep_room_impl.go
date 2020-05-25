package internal

import (
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"redpacket-sweep/common"
	"redpacket-sweep/conf"
	"redpacket-sweep/msg"
)

func (room *RedSweepRoom)Exit(userID int) {
	player := room.userIDPlayerDatas[userID]
	if player.status == playing {
		log.Debug("玩家在游戏中")
		player.writeMsg(&msg.SL2C_ExitRoom{
			Error: msg.SL2C_ExitRoomGame,
		})
		return
	}
	player.offline = false
	room.StandUp(userID)
	delete(userIDRooms, userID)
	room.broadcast(&msg.SL2C_ExitRoom{
		Error: msg.SL2C_ExitRoomOK,
		AccountID: player.accID(),
	})
	updateUserData(userID, bson.M{"$set": bson.M{"gameid": 0}})
	if !player.user().isRobot() {
		player.user().WriteMsg(&msg.S2C_Close{})
		player.user().Close()
	}
	log.Debug("userID: %v 退出 %v 房: %v, 类型: %v, 索引: %v", userID, roomTypeString(room.roomType), room.number, room.roomType, room.roomIndex)
	if !room.checkDisband() {
		room.transferOwner(userID)
	}

	log.Debug("【房间数】：%v， 房间指针：%v", len(roomNumberRooms), roomNumberRooms)
}

func (room *RedSweepRoom) Enter(user *User) {
	if player, ok := room.userIDPlayerDatas[user.userID()]; ok {
		player.user().reconnect(room)
		player.offline = true
	}
	if user.systemOff() {
		return
	}

	room.SitDown(user.userID())

	room.segmentRobot()

	playerData := room.userIDPlayerDatas[user.userID()]
	playerData.offline = true
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

	log.Debug("人数:%v  userID: %v 进入 %v房: %v, 房主: %v, 类型: %v, 索引: %v", len(room.userIDPlayerDatas), user.userID(), roomTypeString(room.roomType), room.number, room.ownerUserID, room.roomType, room.roomIndex)
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
	log.Debug("userID: %v 坐下 %v房: %v, 类型: %v, 索引: %v", userID, roomTypeString(room.roomType), room.number, room.roomType, room.roomIndex)
	return true, player.userID
}

func (room *RedSweepRoom) StandUp(userID int) {
	player := room.userIDPlayerDatas[userID]
	delete(room.userIDPlayerDatas, userID)
	room.broadcastExclude(&msg.SL2C_StandUp{
		Error:      msg.SL2C_SitDown_OK,
		AccountID:  player.accID(),
		Num:		int64(len(room.userIDPlayerDatas)),
	}, player.userID)
}

func (room *RedSweepRoom)checkDisband() bool {
	if len(room.userIDPlayerDatas) < 1 {
		room.clean()
		return true
	}
	return false
}

func (room *RedSweepRoom)clean() {
	delete(roomNumberRooms, room.number)
}

// 转移房主
func (room *RedSweepRoom)transferOwner(userID int) {
	if room.ownerUserID == userID {
		for _, other := range room.userIDPlayerDatas {
			room.ownerUserID = other.userID
			break
		}
	}
}

func (room *RedSweepRoom)segmentRobot() {
	playernum := int64(len(room.userIDPlayerDatas)- room.robotNum())
	rateNum := room.rule.MaxPlayerNum*int64(conf.Server.SegmentRobotRate)/100
	if rateNum <= playernum {
		count := 0
		for _,v := range room.userIDPlayerDatas {
			if v.isRobot() {
				room.Exit(v.userID)
				count++
			}
			stay := conf.Server.StayRobotNum
			if count >= room.robotNum() - stay {
				break
			}
		}
	}
}
