package model

import (
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	MinerAddr   string `gorm:"column:miner_addr" json:"miner_addr" form:"miner_addr"`       // miner addr
	Domain      string `gorm:"column:domain" json:"domain" form:"domain"`                   // domain
	PaymentSecs int    `gorm:"column:payment_secs" json:"payment_secs" form:"payment_secs"` // seconds to be paid
}

func (Node) TableName() string {
	return "node"
}
