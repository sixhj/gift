package boot

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Mysql struct {
	Path        string `mapstructure:"path" json:"path" yaml:"path"`                               // 服务器地址
	Port        string `mapstructure:"port" json:"port" yaml:"port"`                               // 端口
	Config      string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname      string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username    string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password    string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	MaxIdleCons int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenCons int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode     string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap      bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}

func InitMysql() error {
	m := DefaultConfig.Mysql
	if m.Dbname == "" {
		return errors.New("mysql config err")
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                      dsn,  // DSN data source name
		DefaultStringSize:        255,  // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(m.MaxIdleCons)
	sqlDB.SetMaxOpenConns(m.MaxOpenCons)
	DefaultMysql = db

	return nil
}
