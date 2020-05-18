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
}
