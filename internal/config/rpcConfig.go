package config

import (
	"strconv"
	"time"
)

type rpcConfig struct {
	scheme   string
	user     string
	password string
	host     string
	port     string
	timeout  string
	path     string
	useMock  string
	mockDir  string
}

type RpcConfigurator interface {
	GetScheme() string
	GetUser() string
	GetPassword() string
	GetHost() string
	GetPort() string
	GetTimeout() time.Duration
	GetPath() string
	UseMock() bool
	GetMockDir() string
}

func (r *rpcConfig) UseMock() bool {
	useMock, _ := strconv.ParseBool(r.useMock)
	return useMock
}
func (r *rpcConfig) GetMockDir() string {
	return r.mockDir
}
func (r *rpcConfig) GetUser() string {
	return r.user
}

func (r *rpcConfig) GetPassword() string {
	return r.password
}

func (r *rpcConfig) GetHost() string {
	return r.host
}

func (r *rpcConfig) GetPort() string {
	return r.port
}

func (r *rpcConfig) GetTimeout() time.Duration {
	t, _ := strconv.ParseInt(r.timeout, 10, 64)
	return time.Duration(t) * time.Second
}

func (r *rpcConfig) GetPath() string {
	return r.path
}

func (r *rpcConfig) GetScheme() string {
	return r.scheme
}
