// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTKeunggulan = "t_keunggulan"

// TKeunggulan mapped from table <t_keunggulan>
type TKeunggulan struct {
	CIDKeunggulan       int32     `gorm:"column:c_id_keunggulan;primaryKey;autoIncrement:true" json:"c_id_keunggulan"`
	CIDPenanda          int32     `gorm:"column:c_id_penanda" json:"c_id_penanda"`
	CJenisKeunggulan    string    `gorm:"column:c_jenis_keunggulan" json:"c_jenis_keunggulan"`
	CIsNational         string    `gorm:"column:c_is_national;not null;default:N" json:"c_is_national"`
	CCounterPesaing     string    `gorm:"column:c_counter_pesaing" json:"c_counter_pesaing"`
	CMethodePenyampaian string    `gorm:"column:c_methode_penyampaian" json:"c_methode_penyampaian"`
	CIsFollowUp         string    `gorm:"column:c_is_follow_up;not null;default:N" json:"c_is_follow_up"`
	CStatus             string    `gorm:"column:c_status;not null;default:Aktif" json:"c_status"`
	CUpdater            string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate         time.Time `gorm:"column:c_last_update;default:CURRENT_TIMESTAMP" json:"c_last_update"`
	CTahunAjaran        string    `gorm:"column:c_tahun_ajaran;not null;default:2018/2019" json:"c_tahun_ajaran"`
}

// TableName TKeunggulan's table name
func (*TKeunggulan) TableName() string {
	return TableNameTKeunggulan
}
