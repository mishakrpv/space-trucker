package cmd

import (
	pconfig "github.com/space-trucker/iam/pkg/config"
)

type ZzIamCmdConfiguration struct {
	pconfig.Cfg
}

func NewZzIamCmdConfiguration() *ZzIamCmdConfiguration {
	return &ZzIamCmdConfiguration{}
}
