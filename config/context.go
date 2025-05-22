package config

import (
	"fmt"
	"syscall"
)

type AppContext struct {
	*ServerConfig
}

func (*AppContext) WithAppConfig() *AppContext {
	serverConfig, err := NewServerConfig()
	if err != nil {
		fmt.Println(err)
		syscall.Exit(1)
	}
	return &AppContext{
		ServerConfig: serverConfig,
	}
}
func (AppContext) IsAppContext() {}
