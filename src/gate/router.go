package gate

import (
	"redpacket-sweep/game"
	"redpacket-sweep/login"
	"redpacket-sweep/msg"
)

func init() {
	//authorize
	msg.Processor.SetRouter(&msg.C2S_TokenAuthorize{}, login.ChanRPC)
	// game
	msg.Processor.SetRouter(&msg.C2S_Heartbeat{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2SL_StartMatch{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2SL_SendRedPacket{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2SL_TakenRedPacket{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2SL_ExitRoom{}, game.ChanRPC)
}
