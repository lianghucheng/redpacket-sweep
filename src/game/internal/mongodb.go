package internal

import (
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"redpacket-sweep/conf"
)

var (
	DB = conf.GetCfgMongoDB().DBName
	mongoDB *mongodb.DialContext
)

func init() {
	db, err := mongodb.Dial(conf.GetCfgMongoDB().DBUrl, conf.GetCfgMongoDB().DBMaxConnNum)
	if err != nil {
		log.Fatal("dial mongodb error: %v", err)
	}
	mongoDB = db
}

func mongoDBDestroy() {
	mongoDB.Close()
	mongoDB = nil
}

func mongoDBNextSeq(id string) (int, error) {
	return mongoDB.NextSeq(DB, "counters", id)
}