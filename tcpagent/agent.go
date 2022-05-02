package tcpagent

import (
	"c2c/common/def"
	"c2c/common/logger"
	"c2c/controller"
	"crypto"
	"crypto/ecdsa"
	"net"
	"sync"
)

// 注册到controller中
func init() {
	var TcpAgent Agent = Agent{}
	params, ok := controller.AgentParamsLists["TCPAGENT"]
	if !ok {
		logger.Errorf("error for \" %s \" not found\n", "TCPAGENT")
	}
	TcpAgent.Params = params
	controller.RegisterAgent(&TcpAgent)
}

// 描述 tcp agent自身的连接
type Agent struct {
	// 提供client的listener
	FrontListener net.Listener
	// 连接到tcpagent的client list
	ClientList sync.Map
	// 存储agent自身的各个公钥和私钥
	Secrets Secret
	// Agent的一些参数
	Params def.AgentParams
	// 匿名接口
	def.AgentInterface
}

// 存储ECDH密钥交换算法密钥对和ECDSA签名算法密钥对
type Secret struct {
	AESSecret    []byte
	AESIv        []byte
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
