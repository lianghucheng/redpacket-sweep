package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/name5566/leaf/log"
	"redpacket-sweep/metadata"
)

// Config 配置类型
type PublicConfig struct {
	CfgLeaf            		CfgLeaf
	CfgTimeOut         		CfgTimeOut
	CfgMongoDB				CfgMongoDB
	CfgMatchRoomMateData	[]CfgMatchRoomMateData
	CfgCountDown			CfgCountDown
}

type CfgLeaf struct {
	LogLevel    	string
	LogPath     	string
	WSAddr      	string
	CertFile    	string
	KeyFile     	string
	TCPAddr     	string
	MaxConnNum  	int
	ConsolePort 	int
	ProfilePath 	string
	//集群配置信息
	ListenAddr  	string
	ConnAddrs   	[]string
	ServerName		string
	ChipGactRate	int64
	StayRobotNum    int
	SegmentRobotRate int
	SegmentRoomRate  int
}

type CfgTimeOut struct {
	HeartbeatTimeout 		int
	ConnectTimeout   		int
	SystemSendRedPacket		int
}

type CfgMongoDB struct {
	DBName 			string
	DBUrl        	string
	DBMaxConnNum 	int
}

type CfgMatchRoomMateData struct {
	metadata.RoomMetaData
}

type CfgCountDown struct {
	DelayCountDown 			int
	PlayingCountDown        int
	IdleCountDown           int
	SetBoomCountDown        int
}

var opts *PublicConfig
var Server CfgLeaf

func init() {
	opts = &PublicConfig{}
	_, err := toml.DecodeFile("conf/gametoken-redsweep.toml", opts)
	if err != nil {
		log.Fatal("配置文件解析错误:%s", err)
	}
	Server = opts.CfgLeaf
	log.Debug("%v:", Server)
}

func GetCfgTimeout() *CfgTimeOut {
	return &opts.CfgTimeOut
}

func GetCfgMongoDB() *CfgMongoDB {
	return &opts.CfgMongoDB
}

func GetCfgMatchRoomMateData() *[]CfgMatchRoomMateData {
	return &opts.CfgMatchRoomMateData
}

func GetCfgCountDown() *CfgCountDown {
	return &opts.CfgCountDown
}