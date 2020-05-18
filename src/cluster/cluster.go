package cluster

import (
	"encoding/json"
	"math"
	"sync"
	"time"

	"redpacket-sweep/conf"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	lJson "github.com/name5566/leaf/network/json"
)

var (
	server           *network.TCPServer
	clients          []*network.TCPClient
	ClusterProcessor = lJson.NewProcessor()
	agents           sync.Map
)

func Init() {
	if conf.Server.ListenAddr != "" {
		server = new(network.TCPServer)
		server.Addr = conf.Server.ListenAddr
		server.MaxConnNum = int(math.MaxInt32)
		server.PendingWriteNum = conf.PendingWriteNum
		server.LenMsgLen = 4
		server.MaxMsgLen = math.MaxUint32
		server.NewAgent = newAgent

		server.Start()
	}

	for _, addr := range conf.Server.ConnAddrs {
		client := new(network.TCPClient)
		client.Addr = addr
		client.ConnNum = 1
		client.ConnectInterval = 3 * time.Second
		client.PendingWriteNum = conf.PendingWriteNum
		client.LenMsgLen = 4
		client.MaxMsgLen = math.MaxUint32
		client.NewAgent = newAgent
		client.AutoReconnect = true

		client.Start()
		clients = append(clients, client)
	}
}

func Destroy() {
	if server != nil {
		server.Close()
	}

	for _, client := range clients {
		if client != nil {
			client.Close()
		}
	}
}

type Agent struct {
	// WriteMsg  func(a *Agent) (id string, data interface{})
	processor network.Processor
	conn      *network.TCPConn
}

func newAgent(conn *network.TCPConn) network.Agent {
	a := new(Agent)
	a.conn = conn
	a.processor = ClusterProcessor
	if _, ok := agents.LoadOrStore(a, true); ok {
		log.Fatal("duplicate agent %v", a)
	}
	return a
}

func (a *Agent) Run() {
	// 向集群服务器注册服务
	a.WriteMsg("RegistServer", map[string]interface{}{"ServerAddr": conf.Server.WSAddr, "ServerName": conf.Server.ServerName})
	for {
		msg, err := a.conn.ReadMsg()
		if err != nil {
			log.Debug("read message12: %v %v", err, string(msg))
			break
		}
		log.Debug("收到得消息:%v", string(msg))
		// 验证签名
		data := checkSignature(msg)
		if data == nil {
			log.Debug("checkSignature message %v fail", string(msg))
			break
		}
		if a.processor != nil {
			msg, err := a.processor.Unmarshal(data)
			if err != nil {
				log.Debug("unmarshal message error: %v", err)
				break
			}
			err = a.processor.Route(msg, a)
			if err != nil {
				log.Debug("route message error: %v", err)
				break
			}
		}
	}
}

func (a *Agent) OnClose() {
	if _, ok := agents.Load(a); ok {
		agents.Delete(a)
	}
	Destroy()
}

func (a *Agent) WriteMsg(id string, data interface{}) error {
	jsonData := map[string]interface{}{id: data}
	str, _ := json.Marshal(jsonData)
	sign := signature(string(str))
	finalData := map[string]interface{}{"Data": string(str), "Sign": sign}
	pack, err := json.Marshal(finalData)
	if err != nil {
		log.Error("marshal data fail %v", err)
		return err
	}
	err = a.conn.WriteMsg(pack)
	if err != nil {
		log.Error("write fail %v", err)
		return err
	}
	return nil
}

func Broadcast(id string, data interface{}) {
	agents.Range(func(key interface{}, value interface{}) bool {
		if a, ok := key.(*Agent); ok {
			a.WriteMsg(id, data)
		}
		return true
	})
}
