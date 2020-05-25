package internal

import (
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"redpacket-sweep/common"
	"redpacket-sweep/conf"
	"redpacket-sweep/msg"
	"time"

)

func (user *User)tokenAuthorize(m *msg.C2S_TokenAuthorize) {
	userData := new(UserData)
	skeleton.Go(func() {
		db := mongoDB.Ref()
		defer mongoDB.UnRef(db)

		err := db.DB(DB).C("users").Find(bson.M{"token": m.Token, "expiredat":bson.M{"$gt": time.Now().Unix()}, "online": true}).One(userData)
		if err != nil {
			log.Debug("find token %v error: %v", m.Token, err)
			userData = nil
			user.WriteMsg(&msg.S2C_Close{})
			user.Close()
		}
	}, func() {
		if userData == nil || user.status == userLogout {
			return
		}
		if oldUser, ok := userIDUsers[userData.UserID]; ok {
			log.Debug("重连进入：ID:%v", userData.AccountID)
			oldUser.resetLogin()
			oldUser.Close()
			oldUser.userData().Token = user.userData().Token
			user.baseData = oldUser.baseData
			userData = oldUser.userData()
		}

		userIDUsers[userData.UserID] = user
		user.baseData.userData = userData
		userID := user.userID()
		user.UserData().(*AgentInfo).userID = userID
		user.autoHeartbeat()
		user.WriteMsg(&msg.S2C_Authorize{})

		//todo: 重连进入房间
		if r, ok := userIDRooms[userData.UserID]; ok {
			_ = r
			//r.Enter(userData.UserID)
		}

		log.Debug("account: %v Token验证登录, 游戏在线人数: %v", userData.AccountID, len(userIDUsers))
	})
}

func (user *User) logout() {
	user.heartbeatTimer = common.StopTimer(user.heartbeatTimer)
	user.exitRoom()
	user.onLogout()
}

func (user *User) onLogout() {
	userID := user.userID()
	//todo:玩家在房间中的话，不删除数据。设置离线,return
	log.Debug("【登出】")

	delete(userIDUsers,userID)
	user.saveUserData()
	log.Debug("account: %v退出游戏", user.baseData.userData.AccountID)
}

//重置登录
func (user *User) resetLogin() {
	if user.UserData() == nil {
		return
	}
	user.UserData().(*AgentInfo).userID = 0
	user.status = userLogout
	user.heartbeatTimer = common.StopTimer(user.heartbeatTimer)
	//todo:设置玩家离线
}

func (user *User)autoHeartbeat() {
	if user.heartbeatStop {
		log.Debug("userID: %v 心跳停止", user.userID())
		user.Close()
		return
	}
	user.heartbeatStop = true
	user.WriteMsg(&msg.S2C_Heartbeat{})
	user.heartbeatTimer = skeleton.AfterFunc(time.Duration(conf.GetCfgTimeout().HeartbeatTimeout)* time.Second, func() {
		user.autoHeartbeat()
	})
}