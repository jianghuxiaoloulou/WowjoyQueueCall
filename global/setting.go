package global

import (
	"WowjoyProject/WowjoyQueueCall/pkg/logger"
	"WowjoyProject/WowjoyQueueCall/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	GeneralSetting  *setting.GeneralSettingS
	DatabaseSetting *setting.DatabaseSettingS
	ObjectSetting   *setting.ObjectSettingS
	Logger          *logger.Logger
	WebLogger       *logger.Logger
)
