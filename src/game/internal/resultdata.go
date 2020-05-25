package internal

import (
	"redpacket-sweep/common"
	"redpacket-sweep/msg"
)

type RoundResult struct {
	AccountID		int
	NickName		string
	Headimgurl		string
	WinCoin			int64
	IsWin			bool
}

func ToMsgRoundResult(accid int, roundResults map[int]*RoundResult) []*msg.RoundResult {
	rt := make([]*msg.RoundResult, 0)
	me := roundResults[accid]
	if me != nil {
		rt = append(rt, me.ToMsgRoundResult())
	}
	for _,v := range roundResults {
		if v.AccountID == accid {
			continue
		}
		rt = append(rt, v.ToMsgRoundResult())
	}

	return rt
}

func (ctx *RoundResult)ToMsgRoundResult() *msg.RoundResult {
	rt := new(msg.RoundResult)
	rt.NickName = ctx.NickName
	rt.AccountID = ctx.AccountID
	rt.Headimgurl = ctx.Headimgurl
	rt.IsWin = ctx.IsWin
	rt.WinCoin = common.TranferChipRate(ctx.WinCoin)

	return rt
}