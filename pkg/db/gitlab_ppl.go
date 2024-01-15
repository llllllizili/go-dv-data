package db

import (
	"sre-dashboard/pkg/global"
	"sre-dashboard/pkg/logging"
	"time"
)

// 表名
func (merge *PipelineData) TableName() string {
	return "pipeline_data"
}

// 表列
type PipelineData struct {
	ID       int        `gorm:"primarykey"`
	Project  string     `gorm:"column:project"`
	Duration int        `gorm:"column:duration"`
	Status   string     `gorm:"column:status"`
	UpdataAt *time.Time `gorm:"column:updata_at"`
}

// Add 插入一条新的记录，返回主键ID及成功标志
func (ppl *PipelineData) Add() (success bool) {
	db := GetDB()
	defer CloseDB(db)

	if global.DBConfig.Init {
		logging.CLILog.Debug("init PipelineData db")
		r := db.AutoMigrate(&PipelineData{})
		if r != nil {
			logging.CLILog.Error(r)
			return false
		}
	}

	if result := db.Create(ppl); result.RowsAffected > 0 {
		return true
	} else {
		logging.CLILog.Error(result)
		return false
	}
}

func (ppl *PipelineData) IDSearch() (success bool) {
	db := GetDB()
	defer CloseDB(db)

	result := db.First(&ppl, ppl.ID)

	if result.Error != nil {
		logging.CLILog.Info(result.Error)
		return false
	} else {
		return true
	}

}
