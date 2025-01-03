package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS //ServerSettingS 结构体指针
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
