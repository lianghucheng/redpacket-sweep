package internal

import (
	"github.com/name5566/leaf/log"
	"redpacket-sweep/metadata"
)

func (user *User)sendRedPacket(room *RedSweepRoom, redPacketMetadata *metadata.RedPacketMetaData) {
	room.stopSystemRedPacket()
	log.Debug("【发红包】userid:%v", user.userID())
	room.sendRedPacket(user, room.newRedPacketQueue(user.userID(), redPacketMetadata))
}

func (user *User)takenRedPacket(room *RedSweepRoom) {
	if room.status != playing {
		log.Debug("【还没有人发红包】")
		return
	}

	room.takenRedPacket(user)
}