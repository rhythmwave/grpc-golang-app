// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTIsiBe = "t_isi_be"

// TIsiBe mapped from table <t_isi_be>
type TIsiBe struct {
	CIDIsiBe    int32     `gorm:"column:c_id_isi_be;primaryKey;autoIncrement:true" json:"c_id_isi_be"`
	CUraianIsi  string    `gorm:"column:c_uraian_isi" json:"c_uraian_isi"`
	CKelompokBe string    `gorm:"column:c_kelompok_be;not null;default:Internal" json:"c_kelompok_be"`
	CStatus     string    `gorm:"column:c_status;default:Aktif" json:"c_status"`
	CUpdater    string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TIsiBe's table name
func (*TIsiBe) TableName() string {
	return TableNameTIsiBe
}
