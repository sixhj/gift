package boot

import (
	"gift/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	VIPER  *viper.Viper
	SUGAR  *zap.SugaredLogger
	LOG    *zap.Logger
	CONFIG config.Config
)

func Run() error {
	if err := LoadConfig(); err != nil {
		return err
	}
	return nil
}
