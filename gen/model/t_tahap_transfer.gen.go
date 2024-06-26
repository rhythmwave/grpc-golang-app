// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTTahapTransfer = "t_tahap_transfer"

// TTahapTransfer mapped from table <t_tahap_transfer>
type TTahapTransfer struct {
	CKodePengajuanAnggaran string    `gorm:"column:c_kode_pengajuan_anggaran;primaryKey" json:"c_kode_pengajuan_anggaran"`
	CTahapKe               int16     `gorm:"column:c_tahap_ke;primaryKey" json:"c_tahap_ke"`
	CJumlahTransfer        int32     `gorm:"column:c_jumlah_transfer;not null" json:"c_jumlah_transfer"`
	CTanggalTransfer       time.Time `gorm:"column:c_tanggal_transfer;not null" json:"c_tanggal_transfer"`
	CIsTransfer            string    `gorm:"column:c_is_transfer;not null;default:n" json:"c_is_transfer"`
	CUpdater               string    `gorm:"column:c_updater;not null" json:"c_updater"`
	CLastUpdate            time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TTahapTransfer's table name
func (*TTahapTransfer) TableName() string {
	return TableNameTTahapTransfer
}
