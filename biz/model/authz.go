package model

import (
	"gorm.io/gorm"
)

type Authz struct {
	gorm.Model
	AccessKey string `gorm:"column:access_key" json:"access_key" form:"access_key"` // access key
	SecretKey string `gorm:"column:secret_key" json:"secret_key" form:"secret_key"` // secret key
	Role      string `gorm:"column:role" json:"role" form:"role"`                   // role: admin/miner/developer
}

func (Authz) TableName() string {
	return "authz"
}
