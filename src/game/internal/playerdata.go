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
	carryCoin	int64
	winCoin		int64
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

func (data *PlayerData)isRobot() bool {
	return data.user().isRobot()
}

func (data *PlayerData)accID() int {
	return data.user().accID()
}

func (data *PlayerData)nickname() string  {
	return data.user().nickname()
}

func newPlayerData(userID int) *PlayerData {
	user := userIDUsers[userID]
	data := new(PlayerData)
	data.userID = userID
	data.carryCoin = user.chips()
	return data
}