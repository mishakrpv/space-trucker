package cmd

import (
	pconfig "github.com/space-trucker/iam/pkg/config"
)

type ZzIamCmdConfiguration struct {
	pconfig.Cfg
}

func NewZzIamCmdConfiguration() *ZzIamCmdConfiguration {
	return &ZzIamCmdConfiguration{
		Cfg: pconfig.Cfg{
			Host: "localhost",
			Port: "8080",
			Log: &pconfig.Log{
				Level: "INFO",
			},
		},
	}
}
