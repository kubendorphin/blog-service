package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

// 数据库连接字符串 DatabaseSettingS
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dialector := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	// debug
	fmt.Println(
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		//db.LogMode(true)
		db.Debug()
	}
	db.SingularTable(true)
	//获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	updateTimeStampForCreateCallback(db)
	updateTimeStampForUpdateCallback(db)
	deleteCallback(db)
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}

// 新增行为的回调 基于旧版本gorm编写
//
//	func updateTimeStampForCreateCallback(scope *gorm.Scope) {
//		if scope.HasError() {
//			nowTime := time.Now().Unix()
//			if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
//				if createTimeField.IsBlank {
//					_ = createTimeField.Set(nowTime)
//				}
//			}
//			if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
//				if modifyTimeField.IsBlank {
//					_ = modifyTimeField.Set(nowTime)
//				}
//			}
//		}
//	}
//
// 基于新版本gorm编写
func updateTimeStampForCreateCallback(db *gorm.DB) {
	db.Callback().Create().Before("gorm:cteate").Register("update_timestamp", func(tx *gorm.DB) {
		nowTime := tx.Statement.Context.Value("nowTime")
		if nowTime == nil {
			nowTime = tx.NowFunc().Unix()
		}
		//tx.Statement.SetColumn("CreatedOn", nowTime)
		if tx.Statement.Schema != nil {
			if createdOnField, ok := tx.Statement.Schema.FieldsByName["CreatedOn"]; ok {
				if createdOnField.Tag.Get("UPDATE") == "false" {
					tx.Statement.SetColumn("CreatedOn", nowTime)
				}
			}
		}
		tx.Statement.SetColumn("ModifiedOn", nowTime)
	})
}

// 新增更新行为的回调
func updateTimeStampForUpdateCallback(db *gorm.DB) {
	db.Callback().Update().Before("gorm:update").Register("update_timestamp", func(tx *gorm.DB) {
		nowTime := tx.Statement.Context.Value("nowTime")
		if nowTime == nil {
			nowTime = tx.NowFunc().Unix()
		}
		tx.Statement.SetColumn("ModifiedOn", nowTime)
	})
}

// 新增删除行为的回调
func deleteCallback(db *gorm.DB) {
	db.Callback().Delete().Before("gorm:delete").Register("delete_timestamp", func(tx *gorm.DB) {
		nowTime := tx.Statement.Context.Value("nowTime")
		if nowTime == nil {
			nowTime = tx.NowFunc().Unix()
		}
		tx.Statement.SetColumn("DeletedOn", nowTime)
	})
}
