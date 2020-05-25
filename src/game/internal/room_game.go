package internal

import (
	"github.com/name5566/leaf/log"
	"math/rand"
	"redpacket-sweep/common"
	"redpacket-sweep/conf"
	"redpacket-sweep/metadata"
	"redpacket-sweep/msg"
	"time"
)

func (room *RedSweepRoom)sendRedPacket(user *User,redPacket []*RedPacket) {
	room.allPacketQueue = append(room.allPacketQueue,redPacket...)
	room.redPacketQueue = append(room.redPacketQueue, redPacket...)

	//当user为nil时代表系统发红包
	if user != nil {
		player := room.userIDPlayerDatas[user.userID()]

		player.status = playing
		player.redPacketNum += len(redPacket)
	}

	room.startRedPacket()
}

func (room *RedSweepRoom)startRedPacket() {
	if room.status != roomIdle {
		return
	}
	room.status = roomGame

	room.startGame()
}

func (room *RedSweepRoom)takenRedPacket(user *User) {
	curr := room.currentRedpacket
	if curr == nil {
		log.Debug("【还没有装红包】")
		return
	}

	if user.chips() < curr.Quota * int64(room.rule.Multiple) {
		log.Debug("【你的钱不够了】")
		return
	}

	if user.userID() == curr.userID {
		log.Debug("【不能抢自己的】")
		return
	}
	player := room.userIDPlayerDatas[user.userID()]
	taken := int64(0)
	if curr.TakenNum >= room.rule.RedPacketNum{
		log.Debug("【手速慢了哟】")
		return
	} else if curr.TakenNum + 1 == room.rule.RedPacketNum {
		taken = curr.Total
		curr.isOpen = true
	}else {
		//不允许摇出100%和0%的概率
		base := 98
		n := rand.Intn(base) + 1

		taken = curr.Total * int64(n) / int64(100)
	}
	player.takenCoin = taken
	curr.Total -= taken
	curr.TakenNum++
	room.takenPlayer = append(room.takenPlayer, player)

	player.writeMsg(&msg.SL2C_TakenRedPacket{
		TakenCoin:common.TranferChipRate(player.winCoin),
	})
}

func (room *RedSweepRoom)startGame() {
	redpacket := room.redPacketQueue[0]
	room.redPacketQueue = room.redPacketQueue[1:]

	if redpacket.userID != -1 {
		player := room.userIDPlayerDatas[redpacket.userID]
		player.redPacketNum--
		player.takenCoin = -redpacket.Total
	}

	room.currentRedpacket = redpacket
	room.broadcast(&msg.SL2C_StartGame{})

	skeleton.AfterFunc(time.Duration(conf.GetCfgCountDown().SetBoomCountDown) *time.Second, func() {
		skeleton.AfterFunc(time.Duration(conf.GetCfgCountDown().PlayingCountDown) *time.Second, func() {
			room.endGame()
		})
	})
}

func (room *RedSweepRoom)endGame() {
	room.broadcast(&msg.SL2C_EndGame{})
	if len(room.redPacketQueue) < 1 {
		log.Debug("【红包发完了】")
		room.status = roomIdle
		room.systemSendRedPacket()
	}
	curr := room.currentRedpacket
	if curr.userID != -1 {
		player := room.userIDPlayerDatas[curr.userID]
		if player.redPacketNum <= 0 {
			player.status = waiting
		}
	}

	if len(room.takenPlayer) == 0 {
		log.Debug("【无人抢】")
	}

	room.settlement()
	room.cleanRedPacket()

	if room.status != roomIdle {
		skeleton.AfterFunc(2 * time.Second, func() {
			room.startGame()
		})
	}
}

func (room *RedSweepRoom)settlement() {
	roundResults := make(map[int]*RoundResult)
	allLoseCoin := int64(0)
	curr := room.currentRedpacket
	log.Debug("【领取红包的人】%v", room.takenPlayer)
	for _, player := range room.takenPlayer {
		roundResult := new(RoundResult)
		log.Debug("【结算之前】%v", player.takenCoin)
		boom := int64(common.TranferChipRate(player.takenCoin) * float64(conf.Server.ChipGactRate / 10)) % 10
		if int(boom) == curr.Boom {
			loseCoin := curr.Quota - player.takenCoin
			log.Debug("【爆炸的钱】%v", loseCoin)
			player.takenCoin -= loseCoin
			allLoseCoin += loseCoin
		}
		log.Debug("【倍数】%v", room.rule.Multiple)
		roundResult.WinCoin = player.takenCoin * int64(room.rule.Multiple)
		if player.takenCoin > 0 {
			roundResult.IsWin = true
		}
		roundResult.Headimgurl = player.user().headimgurl()
		roundResult.AccountID = player.user().accID()
		roundResult.NickName = player.user().nickname()
		roundResults[player.accID()] = roundResult

		log.Debug("【结算数据】%v accounid:%v", *roundResults[player.accID()], player.accID())
	}

	if curr.userID != -1 {
		roundResult := new(RoundResult)
		player := room.userIDPlayerDatas[curr.userID]
		log.Debug("【结算之前】%v", player.takenCoin)
		log.Debug("【收到爆炸的钱】%v", allLoseCoin)
		player.takenCoin += allLoseCoin + curr.Total
		roundResult.WinCoin = player.takenCoin * int64(room.rule.Multiple)
		if player.takenCoin > 0 {
			roundResult.IsWin = true
		}
		roundResult.Headimgurl = player.user().headimgurl()
		roundResult.AccountID = player.user().accID()
		roundResult.NickName = player.user().nickname()
		roundResults[player.accID()] = roundResult

		log.Debug("【结算数据】%v accounid:%v", *roundResults[player.accID()], player.accID())
	}

	room.sendRoundResult(roundResults)
	room.calculate()
}

func (room *RedSweepRoom)sendRoundResult(roundResults map[int]*RoundResult) {
	for _, p := range room.userIDPlayerDatas {
		p.writeMsg(&msg.SL2C_RoundResult{
			RoundResults:ToMsgRoundResult(p.accID(), roundResults),
		})
	}
}

func (room *RedSweepRoom)calculate() {
	for _, v := range room.takenPlayer {
		v.calculate(room.rule.Multiple)
	}
	if room.currentRedpacket.userID != -1 {
		player := room.userIDPlayerDatas[room.currentRedpacket.userID]
		player.calculate(room.rule.Multiple)
	}
}

func (room *RedSweepRoom)cleanRedPacket() {
	room.currentRedpacket = nil
	room.takenPlayer = make([]*PlayerData,0)
	for _, p := range room.userIDPlayerDatas {
		p.takenCoin = 0
	}
}

func (room *RedSweepRoom)newRedPacketQueue(userid int, redPacketMetadata *metadata.RedPacketMetaData) []*RedPacket {
	redPacketQueue := make([]*RedPacket, 0)
	for i := 0; i < redPacketMetadata.Num; i++ {
		redPacket := new(RedPacket)
		redPacket.userID = userid
		redPacket.Total = int64(redPacketMetadata.Quota) * conf.Server.ChipGactRate
		redPacket.Quota = redPacket.Total
		redPacket.Boom = redPacketMetadata.Boom
		redPacketQueue = append(redPacketQueue, redPacket)
	}

	return redPacketQueue
}

func (room *RedSweepRoom)systemSendRedPacket() {
	room.systemSendRedPacketTimer = skeleton.AfterFunc(time.Duration(conf.GetCfgTimeout().SystemSendRedPacket)* time.Second, func() {
		redPacketMetadata := new(metadata.RedPacketMetaData)
		redPacketMetadata.Quota = 80
		redPacketMetadata.Boom= rand.Intn(10)
		redPacketMetadata.Num = 1

		redPacketQueue := room.newRedPacketQueue(-1, redPacketMetadata)

		log.Debug("【发红包】userid:%v", "系统")
		room.sendRedPacket(nil, redPacketQueue)
	})
}

func (room *RedSweepRoom)stopSystemRedPacket() {
	if room.systemSendRedPacketTimer != nil {
		room.systemSendRedPacketTimer.Stop()
		room.systemSendRedPacketTimer = nil
	}
}