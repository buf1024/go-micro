// Package transport is an interface for synchronous connection based communication
package transport

import (
	"time"
)

// Transport is an interface which is used for communication between
// services. It uses connection based socket send/recv semantics and
// has various implementations; http, grpc, quic.
// 传输层提供的接口，重要的是两个
// Dial连接服务端
// Listen服务端监听
type Transport interface {
	Init(...Option) error
	Options() Options
	Dial(addr string, opts ...DialOption) (Client, error)
	Listen(addr string, opts ...ListenOption) (Listener, error)
	String() string
}

// 交互协议
type Message struct {
	Header map[string]string
	Body   []byte
}

// Socket提供接口
// 重要的是两个
// Recv用于接受协议信息
// Send用于发送协议信息
type Socket interface {
	Recv(*Message) error
	Send(*Message) error
	Close() error
	Local() string
	Remote() string
}

type Client interface {
	Socket
}

type Listener interface {
	Addr() string
	Close() error
	Accept(func(Socket)) error
}

type Option func(*Options)

type DialOption func(*DialOptions)

type ListenOption func(*ListenOptions)

var (
	DefaultTransport Transport = newHTTPTransport()

	DefaultDialTimeout = time.Second * 5
)

func NewTransport(opts ...Option) Transport {
	return newHTTPTransport(opts...)
}
