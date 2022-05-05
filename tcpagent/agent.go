package tcpagent

import (
	"c2c/common/def"
	"c2c/common/logger"
	"c2c/controller"
	"c2c/taskqueue"
	"crypto"
	"crypto/ecdsa"
	"net"
	"sync"
)

var Atq *taskqueue.TasksQueue = new(taskqueue.TasksQueue)

// 注册到controller中
func init() {
	var TcpAgent *Agent = &Agent{}
	params, ok := controller.AgentParamsLists["TCPAGENT"]
	if !ok {
		logger.Errorf("error for \" %s \" not found\n", "TCPAGENT")
	}
	// try to listen port
	// if listen port failed, then add failed mark to agent status
	TcpAgent.Params = &params
	controller.RegisterAgent(TcpAgent, Atq)

	// 注册之后运行
	TcpAgent.run()
}

// 描述 tcp agent的结构
type Agent struct {
	FrontListener      net.Listener     // 提供client的listener
	ClientList         sync.Map         // 连接到tcpagent的client list
	Secrets            *Secret          // 存储agent自身的各个公钥和私钥
	Params             *def.AgentParams // Agent的一些参数
	def.AgentInterface                  // 匿名接口
}

// 存储ECDH密钥交换算法密钥对和ECDSA签名算法密钥对
type Secret struct {
	AESSecret    []byte // 提供给后端的aes加密密钥
	AESIv        []byte
	ECDHPubkey   crypto.PrivateKey
	ECDHPrivkey  crypto.PublicKey
	ECDSAPubkey  ecdsa.PublicKey
	ECDSAPrivkey ecdsa.PrivateKey
}
