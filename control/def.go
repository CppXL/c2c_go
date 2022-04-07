package control

type TcpAgentHead struct {
	Sign   uint16 // Sign word 2 bytes
	CToken uint32 // 4 bytes
	Crc    uint32 // 4 bytes
	Cmd    uint16 // 2 bytes command
	CmdLen uint16 // 2 bytes command len if exists
}
type TcpClientHead struct {
	CToken    uint32
	Cmd       uint16
	CmdLen    uint16
	ResultLen uint16
}
