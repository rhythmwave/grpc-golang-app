// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTLevelVerif = "t_level_verif"

// TLevelVerif mapped from table <t_level_verif>
type TLevelVerif struct {
	CIDBidang   int32     `gorm:"column:c_id_bidang;primaryKey" json:"c_id_bidang"`
	CIDLevel    string    `gorm:"column:c_id_level;primaryKey" json:"c_id_level"`
	CUrut       int16     `gorm:"column:c_urut;primaryKey" json:"c_urut"`
	CNamaLevel  string    `gorm:"column:c_nama_level;default:NULL" json:"c_nama_level"`
	CHargaMin   int32     `gorm:"column:c_harga_min" json:"c_harga_min"`
	CHargaMax   int32     `gorm:"column:c_harga_max" json:"c_harga_max"`
	CUpdater    string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TLevelVerif's table name
func (*TLevelVerif) TableName() string {
	return TableNameTLevelVerif
}