package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// InitGorm gorm初始化
func InitGorm(MysqlDataSource string) *gorm.DB {
	ormLogger := logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       MysqlDataSource,
		DefaultStringSize:         256,  // string类型默认长度
		DisableDatetimePrecision:  true, // 禁止datetime精度
		DontSupportRenameIndex:    true, // 重命名索引
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("数据库连接失败")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)  // 最大连接池
	sqlDB.SetMaxOpenConns(100) // 打开连接数

	return db
}
