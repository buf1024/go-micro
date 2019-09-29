// Package pool is a connection pool
package pool

import (
	"time"

	"github.com/micro/go-micro/transport"
)

// 连接池
// Pool is an interface for connection pooling
type Pool interface {
	// Close the pool
	Close() error
	// Get a connection 获取连接，获取后删除
	Get(addr string, opts ...transport.DialOption) (Conn, error)
	// Release the connection 释放连接，放回连接池中
	Release(c Conn, status error) error
}

// 封装连接
type Conn interface {
	// unique id of connection
	Id() string
	// time it was created
	Created() time.Time
	// embedded connection
	transport.Client
}

func NewPool(opts ...Option) Pool {
	var options Options
	for _, o := range opts {
		o(&options)
	}
	return newPool(options)
}
