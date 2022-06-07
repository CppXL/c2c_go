package client

import (
	"crypto"
	"net"
	"time"
)

// 描述agent和client之间的连接，对于每个连接到agent的client来说都会有这样的结构存储相关信息
type Client struct {
	Token        uint32     // client token
	conn         net.Conn   // client connection
	SysInfo      systeminfo // client system version info
	CliStatus    byte       // client status
	CliCryptInfo *CryptInfo // client crypt info
	CliConf      ClientConf // client config info
}

// 描述每个agent和client的连接加密部分信息
type CryptInfo struct {
	AESSecret        []byte
	AESIv            []byte
	ClientECDHPubKey crypto.PublicKey
}

type ClientConf struct {
	HbReq time.Duration
}

type systeminfo struct {
	SysType byte
	Major   []byte
	Minor   []byte
}

const (
	WINDOWS byte = 1 + iota
	LINUX
	MAC
)
