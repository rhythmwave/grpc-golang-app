// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTIsiBah = "t_isi_bah"

// TIsiBah mapped from table <t_isi_bah>
type TIsiBah struct {
	CIDBah      int32     `gorm:"column:c_id_bah" json:"c_id_bah"`
	CKodeBab    string    `gorm:"column:c_kode_bab" json:"c_kode_bab"`
	CPertemuan  int16     `gorm:"column:c_pertemuan" json:"c_pertemuan"`
	CUpdater    string    `gorm:"column:c_updater" json:"c_updater"`
	CLastUpdate time.Time `gorm:"column:c_last_update" json:"c_last_update"`
	CID         int32     `gorm:"column:c_id;primaryKey;autoIncrement:true" json:"c_id"`
}

// TableName TIsiBah's table name
func (*TIsiBah) TableName() string {
	return TableNameTIsiBah
}
