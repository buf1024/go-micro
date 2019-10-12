`newRpcHandler`生成`registry.Endpoint`规则：
如proto:
```proto
service Test {
	rpc Call(Request) returns (Response) {}
	rpc PingPong(stream Request) returns (stream Response) {}
}

message Request {
	string name = 1;
}

message Response {
	string msg = 1;
}
```
Handler:
```go
type Test struct{}
func (e *Test) Call(ctx context.Context, req *test.Request, rsp *test.Response) error {

	return nil
}
func (e *Test) PingPong(ctx context.Context, stream test.Test_PingPongStream) error {
}
```

调用:
`newRcpHandler(&Test{})` 函数里面通过反射机制获取函数名称，参数，4个参数时是普通的rpc调用，3个参数时是stream rcp调用。
生成的`registry.Endpoint`：
```go
	[registry.Endpoint{
		Name: "Test.Call",
		Request: &registry.Value{
			Name: "Request",
			Type: "Request",
			Values: [] // ...具体的Request类型和名称
		},
		Response: &registry.Value{
			Name: "Response",
			Type: "Response",
			Values: [] // ...具体的Response类型和名称
		},
		Metadata: map[string]string{
			"stream": "false",
			// ...其他传过来的参数
		},
	},
	// ...其他entry
	]
```
这样，即可通过后获取go函数装换成字符串的表现形式，如此即可对外暴露出来调用。

`func (router *router) Handle(h Handler) error {`
Handle里面添加处理函数，`prepareMethod`检查合理，符合参数规范。

`func newSubscriber(topic string, sub interface{}, opts ...SubscriberOption) Subscriber {`
为订阅函数，第二个参数`sub`有两种形式，一种是函数另外一种是结构体。
函数又有两种参数形式，一种是只有一个参数`reqType`，另外一种是两个参数，第一个参数是`context`,另一个参数是`reqType`
结构体里面的函数也是有两种形式，不同的是，结构体里面所有符合函数签名的都会注册。
生成的registry.Endpoint
```go
registry.Endpoint{
			Name:    "Func", // 结构体的为: 结构体名.函数名
			Request: registry.Value{}// 请求参数的签名,
			Metadata: map[string]string{
				"topic":      topic,
				"subscriber": "true",
			},
		}
```

Server注册的形式:
```go
service := &registry.Service{
		Name:      config.Name,
		Version:   config.Version,
		Nodes:     []*registry.Node{
            Id:       config.Name + "-" + config.Id,
            Address:  addr,
            Metadata: {
                node.Metadata["transport"] = config.Transport.String()
                node.Metadata["broker"] = config.Broker.String()
                node.Metadata["server"] = s.String()
                node.Metadata["registry"] = config.Registry.String()
                node.Metadata["protocol"] = "mucp"
                // ... 其他自定义的键值数
            },
		Endpoints: [
            // subscribe 和 handler的的endpoint
        ],
	}
```

`func (s *rpcServer) ServeConn(sock transport.Socket) `接受新的服务连接


