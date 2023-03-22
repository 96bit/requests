package requests

import (
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const (
	ConfigFile = "./conf/config.yaml"
)

var (
	COURSE *UsersCourse
	OA     *officialaccount.OfficialAccount
	DB     *gorm.DB
	CONFIG *Config
	VP     *viper.Viper
)
