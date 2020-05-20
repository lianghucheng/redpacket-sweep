package internal

func (room *RedSweepRoom)sendRedPacket(redPacket []*RedPacket) {
	room.allPacketQueue = append(redPacket)
	room.startRedPacket()
}

func (room *RedSweepRoom)startRedPacket() {
	if room.status == roomGame {
		return
	}

	//递归
}