package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"redpacket-sweep/cluster"
	"redpacket-sweep/conf"
	"redpacket-sweep/game"
	"redpacket-sweep/gate"
	"redpacket-sweep/login"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	// 启动集群
	cluster.Init()
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
