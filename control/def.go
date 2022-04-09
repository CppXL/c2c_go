package control

type TcpAgentHead struct {
	ASign  uint16 // 2 bytes Agent sign word
	CToken uint32 // 4 bytes Token
	Crc    uint32 // 4 bytes Crc Checksum
	Cmd    uint16 // 2 bytes command
	CmdLen uint16 // 2 bytes command len if exists
}
type TcpClientHead struct {
	CSign     uint16 // 2 bytes Client sign word
	CToken    uint32 // 4 bytes Token
	Crc       uint32 // 4 bytes Crc checksum
	Cmd       uint16 // 2 bytes command from Agent
	CmdLen    uint16 // 2 bytes command len form Agent
	ResultLen uint16 // 2 bytes Client execute command result len
}

// 1 byte of execute command status

type cmd uint16

// command define
const (
	SYSTEMCMD cmd = 0x10 + iota
	DOWNLOADFILE
	UPDATE
	GETSYSINFO
	GETPROCINFO
	SCAN
	SLEEP
)




