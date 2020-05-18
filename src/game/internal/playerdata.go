package internal

//玩家状态
const (
	waiting = iota
	playing
)

//玩家游戏状态（当玩家状态为playing时有效
const (
	Waiting = iota
	ActionBet
)

type PlayerData struct {
	userID 		int
	offline		bool
	status 		int
	gameStatus	int

}

func (data *PlayerData)user() *User {
	return userIDUsers[data.userID]
}

func (data *PlayerData)writeMsg(msg interface{}) {
	if !data.offline {
		data.user().WriteMsg(msg)
	}
}

func (data *PlayerData)isRoot() bool {
	return data.user().isRoot()
}

func (data *PlayerData) playing() bool {
	return data.status == playing
}