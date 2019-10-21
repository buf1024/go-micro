package server

import (
	"github.com/micro/go-micro/codec"
	"github.com/micro/go-micro/transport"
)

type rpcRequest struct {
	service     string // Micro-Service
	method      string // Micro-Method
	endpoint    string // Micro-Endpoint
	contentType string // Content-Type
	socket      transport.Socket // 伪socket
	codec       codec.Codec
	header      map[string]string // transfer层的Header
	body        []byte  // transfer层的Body
	rawBody     interface{}
	stream      bool // stream标识
	first       bool
}

type rpcMessage struct {
	topic       string
	contentType string
	payload     interface{}
}

func (r *rpcRequest) Codec() codec.Reader {
	return r.codec
}

func (r *rpcRequest) ContentType() string {
	return r.contentType
}

func (r *rpcRequest) Service() string {
	return r.service
}

func (r *rpcRequest) Method() string {
	return r.method
}

func (r *rpcRequest) Endpoint() string {
	return r.endpoint
}

func (r *rpcRequest) Header() map[string]string {
	return r.header
}

func (r *rpcRequest) Body() interface{} {
	return r.rawBody
}

func (r *rpcRequest) Read() ([]byte, error) {
	// got a body
	if r.first {
		b := r.body
		r.first = false
		return b, nil
	}

	var msg transport.Message
	err := r.socket.Recv(&msg)
	if err != nil {
		return nil, err
	}
	r.header = msg.Header

	return msg.Body, nil
}

func (r *rpcRequest) Stream() bool {
	return r.stream
}

func (r *rpcMessage) ContentType() string {
	return r.contentType
}

func (r *rpcMessage) Topic() string {
	return r.topic
}

func (r *rpcMessage) Payload() interface{} {
	return r.payload
}
