package config

import (
	"fmt"
	"syscall"
)

const (
	PORT = "PORT"
)

func GetEnvVar(varName string) (string, error) {
	val, found := syscall.Getenv(varName)
	if !found || val == "" {
		err := fmt.Errorf("env var %s not found", varName)
		return "", err
	} else {
		return val, nil
	}
}
