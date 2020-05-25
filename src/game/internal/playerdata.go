package internal

import (
	"github.com/name5566/leaf/log"
	"redpacket-sweep/common"
	"redpacket-sweep/msg"
)

//玩家状态
const (
	waiting = iota
	playing
)

//玩家游戏状态（当玩家状态为playing时有效
const (
	Waiting = iota
	ActionBet
)

type PlayerData struct {
	userID 			int
	offline			bool
	status 			int
	gameStatus		int
	carryCoin		int64
	winCoin			int64
	redPacketNum 	int
	takenCoin 		int64
}

func (data *PlayerData)user() *User {
	return userIDUsers[data.userID]
}

func (data *PlayerData)writeMsg(msg interface{}) {
	if !data.offline {
		data.user().WriteMsg(msg)
	}
}

func (data *PlayerData)isRoot() bool {
	return data.user().isRoot()
}

func (data *PlayerData) playing() bool {
	return data.status == playing
}

func (data *PlayerData)isRobot() bool {
	return data.user().isRobot()
}

func (data *PlayerData)accID() int {
	return data.user().accID()
}

func (data *PlayerData)nickname() string  {
	return data.user().nickname()
}

func newPlayerData(userID int) *PlayerData {
	user := userIDUsers[userID]
	data := new(PlayerData)
	data.userID = userID
	data.carryCoin = user.chips()
	return data
}

func (v *PlayerData)calculate(multi int) {
	winCoin := int64(0)
	if v.takenCoin > 0 {
		winCoin = v.takenCoin * 5 / 100
	} else {
		winCoin = v.takenCoin
	}

	v.winCoin += winCoin * int64(multi)
	log.Debug("【赢的钱】%v", v.winCoin)
	v.user().userData().Chips += winCoin
	v.user().saveUserData()
	v.writeMsg(&msg.SL2C_UpdateChips{
		Chips:common.TranferChipRate(v.user().chips()),
	})
}