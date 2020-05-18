package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/timer"
	"gopkg.in/mgo.v2/bson"
	"redpacket-sweep/msg"
)

var (
	systemOn      = true  // 系统开关
	ipAntiCheatOn = false // IP 防作弊开关
)

// 用户状态
const (
	_          = iota
	userLogin  // 1
	userLogout // 2
)

// 用户角色
const (
	roleRobot  = -2 // 机器人
	roleBlack  = -1 // 拉黑
	rolePlayer = 1  // 玩家
	roleAgent  = 2  // 代理
	roleAdmin  = 3  // 管理员
	roleRoot   = 4  // 超级管理员
)

var (
	userIDUsers = make(map[int]*User)
)

type (
	User struct {
		gate.Agent
		status          int
		baseData       *BaseData
		heartbeatTimer *timer.Timer
		heartbeatStop  bool
	}

	BaseData struct {
		userData *UserData
	}
)

func newUser(a gate.Agent) *User {
	user := new(User)
	user.Agent = a
	user.baseData = new(BaseData)
	user.baseData.userData = new(UserData)
	return user
}

func (user *User) userData() *UserData {
	return user.baseData.userData
}

func (user *User) userID() int {
	return user.userData().UserID
}

func (user *User) accID() int {
	return user.userData().AccountID
}

func (user *User) unionID() string {
	return user.userData().UnionID
}

func (user *User)isRoot() bool {
	return user.userData().Role == roleRoot
}

func (user *User)isRobot() bool {
	return user.userData().Role == roleRobot
}

func (user *User)isPlayer() bool {
	return user.userData().Role == rolePlayer
}

func (user *User) isOffline() bool {
	return user.status == userLogout
}

func (user *User)inRoom() bool {
	return userIDRooms[user.userID()] != nil
}

func (user *User)systemOff() bool {
	if !systemOn {
		log.Debug("系统升级维护中")
		user.WriteMsg(&msg.SL2C_Close{Error: msg.SL2C_Close_SystemOff})
		user.WriteMsg(&msg.S2C_Close{})
		user.Close()
	}
	return !systemOn
}

func (user *User)saveUserData() {
	userdata := user.userData()
	value := bson.M{"$set": bson.M{"winchips": userdata.WinChips, "win" : userdata.Win, "chips" : userdata.Chips}}
	updateUserData(user.userID(), value)
}