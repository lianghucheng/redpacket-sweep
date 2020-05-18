package internal

import (
	"redpacket-sweep/game"
	"redpacket-sweep/msg"
	"reflect"
)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.C2S_TokenAuthorize{}, handleTokenAuthorize)
}

func handleTokenAuthorize(args []interface{}) {
	game.ChanRPC.Go("TokenAuthorize", args[1], args[0])
}