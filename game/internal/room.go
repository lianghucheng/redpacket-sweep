package internal

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"redpacket-sweep/common"
)

// 房间状态
const (
	roomIdle = iota // 0
	roomGame        // 1
)

var (
	userIDRooms     = make(map[int]*RedSweepRoom)
	roomNumberRooms = make(map[string]*RedSweepRoom)
	roomNumbers     = []int{}
	roomCounter     = 0
)

type Room interface {
	Enter(int)
	Exit(int)
	SitDown(int) (bool, int)
	StandUp(int)
	SendAllPlayer(int)
	StartGame()
	EndGame()
}

type room struct {
	state         int
	roomType      int
	roomIndex     int
	number        string // 房号
	userIPAddrs   map[string]bool
	creatorUserID int    // 创建者 userID
	ownerUserID   int    // 房主 userID
	desc          string // 描述
}

func init() {
	for i := 0; i < 1000000; i++ {
		roomNumbers = append(roomNumbers, i)
	}
	roomNumbers = common.Shuffle(roomNumbers)
}

func getRoomNumber() string {
	log.Debug("房间计数器: %v", roomCounter)
	roomNumber := fmt.Sprintf("%06d", roomNumbers[roomCounter])
	roomCounter = (roomCounter + 1) % 1000000
	return roomNumber
}

func upsertRobotData(id string, update interface{}) {
	skeleton.Go(func() {
		db := mongoDB.Ref()
		defer mongoDB.UnRef(db)
		_, err := db.DB(DB).C("dz_robot_profit").UpsertId(id, update)
		if err != nil {
			log.Error("upsert %v error: %v", id, err)
		}
	}, nil)
}