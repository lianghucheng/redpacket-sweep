package internal

//房间类型
const (
	_			= iota
	roomTenMatching		// 1、十个红包匹配场
	roomFiveMacthing	// 2、五个红包匹配场
)

func roomTypeString(roomType int) string {
	switch roomType {
	case roomTenMatching:
		return "十个红包匹配场"
	case roomFiveMacthing:
		return "五个红包匹配场"
	default:
		return "N/A"
	}
}