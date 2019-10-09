go-micro 传输层
=============
`go-micro/transfer`目录下是go-micro的传输层，微服务的交互都需要经过这一层。go语言，无论是对TCP还是HTTP等网络工具的支持都是比较完整和完善，那么，第一个问题，为什么不直接用go语言提供的网络库，为什么要提供者样一层的封装？

从事过网络编程的人都应该知道，交互需要制定协议，所以直接使用go语言提供的网络库，最终导致的结果也是封装成这样一个传输层，`transport`层的就是封装go-micro服务交互的协议。这个协议超级简单:
```go
// go-micro/transport/transport.go
// 交互协议
type Message struct {
	Header map[string]string
	Body   []byte
}
```
第二个问题是如何封装的问题，go-micro是按照人们比较熟悉的socket的概念去封装，`Client`对应有`DialOption`, `Listener`对应有`ListenOption`，当**连接**上来时，就获取到一个`Socket`，`Socket`处理的是传输层的**协议**。
```go
// go-micro/transport/transport.go
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

type Listener interface {
	Addr() string
	Close() error
	Accept(func(Socket)) error
}
```
go-micro目前提供的默认传输层的网络协议是http协议。

### http_transport http传输层
http传输层中Dial连接服务端的，Dial里面有个比较有意思的函数:
```go
// 将原来的函数封装以下，增加代理有功能的连接
func newConn(dial func(string) (net.Conn, error)) func(string) (net.Conn, error)
```
此处，增加代理功能，在http_proxy里面，首先，用`getURL`判断是否有代理，如果有代理，则使用`HTTP CONNECT`的方法连接到目标，判断代理是否可用，同时获取得连接`Conn`，重点，作者不使用原来的`Conn`, 二手使用了一个叫`pbuffer`的结构体重写Read方法