package msg

func init() {
	Processor.Register(&C2SL_StartMatch{})
	Processor.Register(&SL2C_EnterRoom{})
}

type C2SL_StartMatch struct {
		RoomType	int
		ItemType	int
}

const (
	SL2C_EnterRoom_OK                = 0
	SL2C_EnterRoom_InvalidRoomNumber = 1 // 房间: +SL2C_EnterRoom.Number+不存在
	SL2C_EnterRoom_Full              = 2 // 房间: +SL2C_EnterRoom.Number+已满
	SL2C_EnterRoom_Unknown           = 3 // 进入房间出错，请稍后重试
	SL2C_EnterRoom_MinChipsLimit     = 4 // 需要+SL2C_EnterRoom.MinChips+金币才能进入，请先购买金币
	SL2C_EnterRoom_MaxChipsLimit     = 5 // 拥有金币超过+SL2C_EnterRoom.MaxChips
	SL2C_EnterRoom_NotRightNow       = 6 // 比赛暂未开始，请到时再来
)

type SL2C_EnterRoom struct {
	Error         	int
	Desc          	string
	Num 			int64
	Headimgurl    	string
	AccountID		int
	CarryCoin		float64
	WinCoin			float64
}

const (
	SL2C_SitDown_OK            = 0
	SL2C_SitDown_MinChipsLimit = 1 // 需要+DZ2C_SitDown.MinChips+金币才能坐下，请先购买金币
	SL2C_SitDown_MaxChipsLimit = 2 // 拥有金币超过+DZ2C_SitDown.MaxChips
	SL2C_SitDown_PositionFull  = 3 // 位置已满
)

type SL2C_SitDown struct {
	Error 		int
	UserID 		int
	AccountID	int
	Nickname	string
	Num			int64
}

type SL2C_StandUp struct {
	Error		int
	UserID 		int64
}

