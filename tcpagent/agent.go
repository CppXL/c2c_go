package tcpagent

import (
	"c2c/common/logger"
	"crypto"
	"crypto/ecdsa"
	"net"
	"sync"
)

// 描述 tcp agent自身的连接
type Agent struct {
	FrontListener net.Listener
	ClientList    sync.Map
	S             Secret
}

// 存储ECDH密钥交换算法密钥对和ECDSA签名算法密钥对
type Secret struct {
	ECDHPubkey   crypto.PrivateKey
	ECDHPrivkey  crypto.PublicKey
	ECDSAPubkey  ecdsa.PublicKey
	ECDSAPrivkey ecdsa.PrivateKey
}

func Listen(network, address string) (*Agent, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		logger.Errorf(err.Error())
	}

	return &Agent{FrontListener: listener}, nil
}

func (s *Agent) Run() {

}
