package db

import (
	"sre-dashboard/pkg/global"
	"sre-dashboard/pkg/logging"
	"time"
)

// 表名
func (merge *MergeRequest) TableName() string {
	return "gitlab_merge"
}

// 表列
type MergeRequest struct {
	ID        int        `gorm:"primarykey"`
	Title     string     `gorm:"column:title"`
	MergeAt   *time.Time `gorm:"column:merge_at"`
	ProjectID int        `gorm:"column:project_id"`
}

// Add 插入一条新的记录，返回主键ID及成功标志
func (merge *MergeRequest) Add() (success bool) {
	db := GetDB()
	defer CloseDB(db)

	if global.DBConfig.Init {
		logging.CLILog.Debug("init MergeRequest db")
		r := db.AutoMigrate(&MergeRequest{})
		if r != nil {
			logging.CLILog.Error(r)
			return false
		}
	}
	if merge.MergeAt == nil {
		return false
	}

	if result := db.Create(merge); result.RowsAffected > 0 {
		return true
	} else {
		logging.CLILog.Error(result)
		return false
	}
}

func (merge *MergeRequest) IDSearch() (success bool) {
	db := GetDB()
	defer CloseDB(db)

	result := db.First(&merge, merge.ID)

	if result.Error != nil {
		logging.CLILog.Error(result.Error)
		return false
	} else {
		return true
	}

}
