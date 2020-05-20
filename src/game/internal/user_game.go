package internal

import "redpacket-sweep/metadata"

func (user *User)sendRedPacket(redPacketMetadata *metadata.RedPacketMetaData) {
	if room, ok := userIDRooms[user.userID()]; !ok {
		return
	} else {
		redPacketQueue := make([]*RedPacket, 0)
		for i := 0; i < redPacketMetadata.Num; i++ {
			redPacket := new(RedPacket)
			redPacket.RedPacketMetaData = *redPacketMetadata
			redPacket.userID = user.userID()
			redPacketQueue = append(redPacketQueue, redPacket)
		}

		room.sendRedPacket(redPacketQueue)
	}
}