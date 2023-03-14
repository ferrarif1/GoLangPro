package models

import (
	"time"

	"gorm.io/gorm"
)

/*
gorm.Model的作用：帮助自动引入四个字段：
   ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
*/

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LogInTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"colum:login_out_time" json:"login_out_time"` //用这个命令创建别名
	IsLoginout    bool
	DeviceInfo    string
}
