// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTDppbCakupanKerja = "t_dppb_cakupan_kerja"

// TDppbCakupanKerja mapped from table <t_dppb_cakupan_kerja>
type TDppbCakupanKerja struct {
	CNik        string    `gorm:"column:c_nik;primaryKey" json:"c_nik"`
	CIDGedung   int32     `gorm:"column:c_id_gedung;primaryKey" json:"c_id_gedung"`
	CUpdater    string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate time.Time `gorm:"column:c_last_update;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TDppbCakupanKerja's table name
func (*TDppbCakupanKerja) TableName() string {
	return TableNameTDppbCakupanKerja
}