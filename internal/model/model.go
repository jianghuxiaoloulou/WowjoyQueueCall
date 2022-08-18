package model

import (
	"WowjoyProject/WowjoyQueueCall/pkg/setting"
	"database/sql"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type KeyData struct {
	instance_key                                                                         sql.NullInt64
	dcmfile, imgfile, dcmremotekey, jpgremotekey, ip, virpath                            sql.NullString
	jpgstatus, dcmstatus, jpglocalstatus, dcmlocalstatus, jpgcloudstatus, dcmcloudstatus sql.NullInt16
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.DBType, databaseSetting.DBConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(databaseSetting.MaxLifetime) * time.Minute)
	db.SetMaxOpenConns(databaseSetting.MaxIdleConns)
	db.SetMaxIdleConns(databaseSetting.MaxIdleConns)

	return db, nil
}

func NewOtherDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.DBType, databaseSetting.OtherDBConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(databaseSetting.MaxLifetime) * time.Minute)
	db.SetMaxOpenConns(databaseSetting.MaxIdleConns)
	db.SetMaxIdleConns(databaseSetting.MaxIdleConns)

	return db, nil
}
