package global

import (
	"sre-dashboard/pkg/config"
)

var (
	appConfig    = config.GetGlobleConfig()
	LogPath      = &appConfig.LogPath
	TapdConfig   = &appConfig.TapdConfig
	DBConfig     = &appConfig.DBConfig
	GitlabConfig = &appConfig.GitlabConfig
)
