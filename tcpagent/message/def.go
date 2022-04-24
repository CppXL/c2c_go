package message

// magic number for agent
const AgentMagic uint32 = 0xAFBEBEAF

// magic number for client
const ClientMagic uint32 = 0x6F9A9A6F

// agent to client message head format
type TcpAgentMsgHead struct {
	AgentMagic  uint32 // 4 bytes Agent sign word
	ClientToken uint32 // 4 bytes Token
	Crc         uint32 // 4 bytes Crc Checksum
	Cmd         uint16 // 2 bytes command
	CmdLen      uint16 // 2 bytes command len if exists
}

//  client to agent message head format
type TcpClientMsgHead struct {
	ClientMagic uint32 // 4 bytes Client sign word
	ClientToken uint32 // 4 bytes Token
	Crc         uint32 // 4 bytes Crc checksum
	Cmd         uint16 // 2 bytes command from Agent
	CmdLen      uint16 // 2 bytes command len form Agent
	ResultLen   uint16 // 2 bytes Client execute command result len
}

type msg interface {
	TcpClientMsgHead | TcpAgentMsgHead
}
