// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTDaftarPesaing = "t_daftar_pesaing"

// TDaftarPesaing mapped from table <t_daftar_pesaing>
type TDaftarPesaing struct {
	CIDPesaing   int32     `gorm:"column:c_id_pesaing;primaryKey" json:"c_id_pesaing"`
	CNamaPesaing string    `gorm:"column:c_nama_pesaing;default:NULL" json:"c_nama_pesaing"`
	CAlamat      string    `gorm:"column:c_alamat" json:"c_alamat"`
	CUpdater     string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate  time.Time `gorm:"column:c_last_update;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TDaftarPesaing's table name
func (*TDaftarPesaing) TableName() string {
	return TableNameTDaftarPesaing
}
