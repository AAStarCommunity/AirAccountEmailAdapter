package repository

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Mail struct {
	Model
	// `json:"realName" gorm:"column:real_name;type:nvarchar(16);not null;default('')"`
	Sender      string `json:"sender" gorm:"column:sender;type:nvarchar(512);not null;"`
	Receiver    string `json:"receiver" gorm:"column:receiver;type:nvarchar(512);not null;"`
	Subject     string `json:"subject" gorm:"column:subject;type:nvarchar(512);not null;"`
	Unread      bool   `json:"unread" gorm:"column:unread;type:boolean;not null;"`
	Fingerprint string `json:"fingerprint" gorm:"column:fingerprint;type:nvarchar(128);not null;"`
}
