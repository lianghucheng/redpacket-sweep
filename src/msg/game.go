package msg

import "redpacket-sweep/metadata"

func init() {
	Processor.Register(&C2SL_SendRedPacket{})
	Processor.Register(&SL2C_SendRedPacket{})
	Processor.Register(&C2SL_TakenRedPacket{})
	Processor.Register(&SL2C_TakenRedPacket{})
}

type C2SL_SendRedPacket struct {
	metadata.RedPacketMetaData
}

type SL2C_SendRedPacket struct {
	Error 		int
}

type C2SL_TakenRedPacket struct {

}

type SL2C_TakenRedPacket struct {
	Error 		int
}

