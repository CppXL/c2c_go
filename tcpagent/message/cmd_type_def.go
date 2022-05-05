package message

type cmd uint16

// command define
const (
	SYSTEMCMD    cmd = 0x0010 + iota // exec system command
	DLFILE                           // download file form somewhere
	GTFILE                           // get file from client
	UPDATESELF                       // update trojan self
	GETSYSINFO                       // get system information such as cpu users
	GETPROCINFO                      // get process information
	SCAN                             // scan live host
	SLEEP                            // sleep sometime
	SETHBREQ                         // set heart beat frequence
	REVERSESHELL                     // reverse shell to controller
	DESTROYSELF                      // destroy self
	RESERVE                          // reserve bit
	HEARTBEAT    cmd = 0xFFFF        // heartbeat
)

// 这里定义client连接到agent时的交互步骤
const (
	CLIENTHELLO cmd = 0x0100 + iota
	CLIENTHELLOACK
)
