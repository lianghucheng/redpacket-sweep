package msg

import "redpacket-sweep/metadata"

func init() {
	Processor.Register(&C2SL_SendRedPacket{})
	Processor.Register(&SL2C_SendRedPacket{})
	Processor.Register(&C2SL_TakenRedPacket{})
	Processor.Register(&SL2C_TakenRedPacket{})
	Processor.Register(&SL2C_StartGame{})
	Processor.Register(&SL2C_EndGame{})
	Processor.Register(&SL2C_RoundResult{})
}

type C2SL_SendRedPacket struct {
	metadata.RedPacketMetaData
}

type SL2C_SendRedPacket struct {
	Error 		int
}

type C2SL_TakenRedPacket struct {

}

type SL2C_TakenRedPacket struct {
	Error 		int
	TakenCoin	float64
}

type SL2C_StartGame struct {
}

type SL2C_EndGame struct {

}

type RoundResult struct {
	AccountID		int
	NickName		string
	Headimgurl		string
	WinCoin			float64
	IsWin			bool
}

type SL2C_RoundResult struct {
	RoundResults	[]*RoundResult
}