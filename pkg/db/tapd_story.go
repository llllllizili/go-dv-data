package db

import (
	"sre-dashboard/pkg/global"
	"sre-dashboard/pkg/logging"
)

// 表名
func (p *Story) TableName() string {
	return "tapd_story"
}

// 表列
type Story struct {
	ID       string `gorm:"primarykey"`
	Name     string `gorm:"column:name"`
	Owner    string `gorm:"column:owner"`
	Status   string `gorm:"column:status"`
	Creator  string `gorm:"column:creator"`
	Created  string `gorm:"column:create_datetime"`
	Modified string `gorm:"column:update_datetime"`
	Pm       string `gorm:"column:pm"` // 责任人、项目经理
}

// 初始化 db
func (story *Story) Init() (success bool) {
	db := GetDB()
	defer CloseDB(db)
	if global.DBConfig.Init {
		logging.CLILog.Info("init db")
		r := db.AutoMigrate(&Story{})
		if r != nil {
			logging.CLILog.Error(r)
			return false
		}
	}
	return true
}

// Add 插入一条新的记录，返回主键ID及成功标志
func (story *Story) Add() (success bool) {
	db := GetDB()
	defer CloseDB(db)

	if global.DBConfig.Init {
		logging.CLILog.Debug("init TAPD Story db")
		r := db.AutoMigrate(&Story{})
		if r != nil {
			logging.CLILog.Error(r)
			return false
		}
	}

	if result := db.Create(story); result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func (story *Story) IDSearch() (success bool) {
	db := GetDB()
	defer CloseDB(db)

	result := db.First(&story, story.ID)

	if result.Error != nil {
		return false
	} else {
		return true
	}

}
