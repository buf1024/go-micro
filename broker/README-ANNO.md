http_broker:
```go
	// 开始真正的监听
	Connect() error
	// 发布可以理解为客户端发送请求给服务端
	Publish(topic string, m *Message, opts ...PublishOption) error
	// 订阅可理解为服务端监听接受信息
	Subscribe(topic string, h Handler, opts ...SubscribeOption) (Subscriber, error)
```
以上三个为broker的主要接口，用于实现http发布和订阅的功能。

订阅依赖于两个接口：`Connect`和`Subscribe`。
`Connect`实现服务的监听，默认的监听http2监听路径为`/_sub`,当然也可以像传输层一样，通过context的`http_handlers`增加处理路径。
`Connect`实现服务的监听还有一个功能是，Subscribe定时向注册中心注册。
`/_sub`的处理函数里面，是`ServeHTTP`,判断是否是合法的POST， 消息有`:topic`字段，路径有具体subscribe的id，然后调用具体的处理器。

`Subscribe`主要是像注册中心注册服务，格式为：
```go
	node := &registry.Node{
		Id:      id, // broker- + uuid 
		Address: mnet.HostPort(addr, port), // 地址端口
		Metadata: map[string]string{ // 连接属性
			"secure": fmt.Sprintf("%t", secure),
		},
	}

	// check for queue group or broadcast queue
	version := options.Queue
	if len(version) == 0 {
		version = broadcastVersion // ff.http.broadcast是否广播的
	}

	// 注册中心注册的名称
	service := &registry.Service{
		Name:    "topic:" + topic, // 服务名称
		Version: version, // 是否广播
		Nodes:   []*registry.Node{node}, // 服务端口
	}
```

以上，监听需要两个步骤调用，先Connect，然后Subscribe

Publish发布信息
发布消息的消息go routine方式放到inbox里面，避免阻塞。
想根据 `"topic:" + topic`获取服务，然后更加`Version`广播发送还是单个发送
发送Header:
:topic
发送地址:
%s://%s%s?%s 如： https://192.168.1.100:9901/_sub?id=broker-uuid.uuid