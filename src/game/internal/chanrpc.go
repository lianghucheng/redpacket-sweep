package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"redpacket-sweep/conf"
	"redpacket-sweep/msg"
	"strings"
	"time"
)

type AgentInfo struct {
	userID int
}

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("TokenAuthorize", rpcTokenAuthorize)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	a.SetUserData(new(AgentInfo))
	skeleton.AfterFunc(time.Duration(conf.GetCfgTimeout().ConnectTimeout)*time.Second, func() { // 3秒后未登录则断开连接
		closeAgent(a)
	})
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	userID := a.UserData().(*AgentInfo).userID
	// a.SetUserData(nil)

	if user, ok := userIDUsers[userID]; ok {
		user.status = userLogout
		user.logout()
	}
}

func closeAgent(a gate.Agent) {
	if a.UserData() != nil && a.UserData().(*AgentInfo).userID < 1 {
		log.Debug("关闭无效的连接")
		a.Close()
	}
}

func rpcTokenAuthorize(args []interface{}) {
	a := args[0].(gate.Agent)
	m := args[1].(*msg.C2S_TokenAuthorize)
	if !systemOn {
		a.WriteMsg(&msg.SL2C_Close{Error: msg.SL2C_Close_SystemOff})
		a.WriteMsg(&msg.S2C_Close{})
		a.Close()
	}
	if a.UserData() == nil {
		a.Close()
		return
	}
	agentInfo := a.UserData().(*AgentInfo)
	if agentInfo.userID > 0 || strings.TrimSpace(m.Token) == "" {
		a.Close()
		return
	}

	newUser(a).tokenAuthorize(m)
}
