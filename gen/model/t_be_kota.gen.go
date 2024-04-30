// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTBeKotum = "t_be_kota"

// TBeKotum mapped from table <t_be_kota>
type TBeKotum struct {
	CIDBeKota      int32     `gorm:"column:c_id_be_kota;primaryKey;autoIncrement:true" json:"c_id_be_kota"`
	CIDPenanda     int32     `gorm:"column:c_id_penanda;not null" json:"c_id_penanda"`
	CIDIsiBe       int32     `gorm:"column:c_id_isi_be;not null" json:"c_id_isi_be"`
	CKondisiBisnis string    `gorm:"column:c_kondisi_bisnis;not null" json:"c_kondisi_bisnis"`
	CIsFollowUp    string    `gorm:"column:c_is_follow_up;not null;default:N" json:"c_is_follow_up"`
	CTahunAjaran   string    `gorm:"column:c_tahun_ajaran;not null;default:2018/2019" json:"c_tahun_ajaran"`
	CStatus        string    `gorm:"column:c_status;not null;default:Aktif" json:"c_status"`
	CUpdater       string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate    time.Time `gorm:"column:c_last_update;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TBeKotum's table name
func (*TBeKotum) TableName() string {
	return TableNameTBeKotum
}