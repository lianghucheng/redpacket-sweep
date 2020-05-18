package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/name5566/leaf/log"
)

// Config 配置类型
type PublicConfig struct {
	CfgLeaf            	CfgLeaf
	CfgTimeOut         	CfgTimeOut
	CfgMongoDB			CfgMongoDB
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
}

type CfgTimeOut struct {
	HeartbeatTimeout int
	ConnectTimeout   int
}

type CfgMongoDB struct {
	DBName 			string
	DBUrl        	string
	DBMaxConnNum 	int
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