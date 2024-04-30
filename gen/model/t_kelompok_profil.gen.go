// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTKelompokProfil = "t_kelompok_profil"

// TKelompokProfil mapped from table <t_kelompok_profil>
type TKelompokProfil struct {
	CIDKelompok   int32     `gorm:"column:c_id_kelompok;primaryKey;autoIncrement:true" json:"c_id_kelompok"`
	CNamaKelompok string    `gorm:"column:c_nama_kelompok;default:NULL" json:"c_nama_kelompok"`
	CStatus       string    `gorm:"column:c_status;default:NULL" json:"c_status"`
	CUpdater      string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate   time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TKelompokProfil's table name
func (*TKelompokProfil) TableName() string {
	return TableNameTKelompokProfil
}