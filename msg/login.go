package msg

func init() {
	Processor.Register(&SL2C_Close{})
	Processor.Register(&C2S_Heartbeat{})
	Processor.Register(&S2C_Heartbeat{})
	Processor.Register(&SL2C_UpdateChips{})
	Processor.Register(&C2S_TokenAuthorize{})
	Processor.Register(&S2C_Authorize{})
	Processor.Register(&S2C_Close{})
}

const (
	SL2C_Close_LoginRepeated = 1 // 您的账号在其他设备上线，非本人操作请注意修改密码
	SL2C_Close_IDZerError    = 2 // 登录出错，请重新登录
	SL2C_Close_TokenInvalid  = 3 // 登录状态失效，请重新登录
	SL2C_Close_SystemOff     = 4 // 系统升级维护中，请稍后重试

)

type SL2C_Close struct {
	Error        int
	WeChatNumber string
}

type (
	C2S_Heartbeat struct{}

	S2C_Heartbeat struct{}
)

type SL2C_UpdateChips struct {
	Chips float64
}

type C2S_TokenAuthorize struct {
	Token   string
	Connect bool
}

type S2C_Authorize struct {
}

type S2C_Close struct{}