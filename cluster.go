package internal

import (
	"dz/cluster"

	"github.com/name5566/leaf/console"
	"github.com/name5566/leaf/log"
)

func init() {
	// cluster.ClusterProcessor.Register(&msg.UpdateGameConfig{})
	// cluster.ClusterProcessor.SetRouter(&msg.UpdateGameConfig{}, ChanRPC)

	// handler(&msg.UpdateGameConfig{}, updateGameConfig)
	// console.Register("updateGameConfig", "update config from db", updateGameConfig, ChanRPC)
	console.Register("broadcast", "broadcast to all hall", broadcastHall, ChanRPC)
}

// func updateGameConfig(args []interface{}) interface{} {
// 	if err := initGameConfig(); err != nil {
// 		return err.Error()
// 	}
// 	log.Release("update config:%v", *configData)
// 	return "update success"
// }

func broadcastHall(args []interface{}) interface{} {
	if len(args) == 0 {
		log.Error("broadcast args invalid:%v", args)
		return "broadcast args invalid"
	}
	s, ok := args[0].(string)
	if !ok {
		log.Error("broadcast args invalid:%v", args)
		return "broadcast args invalid"
	}
	var data interface{}
	if len(args) > 1 {
		data = args[1]
	}
	cluster.Broadcast(s, data)
	return "broadcast success"
}
