// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTKebutuhanAnggaranSplit = "t_kebutuhan_anggaran_split"

// TKebutuhanAnggaranSplit mapped from table <t_kebutuhan_anggaran_split>
type TKebutuhanAnggaranSplit struct {
	CKodeSplitAnggaran string    `gorm:"column:c_kode_split_anggaran;primaryKey" json:"c_kode_split_anggaran"`
	CKodeKegiatan      string    `gorm:"column:c_kode_kegiatan;primaryKey" json:"c_kode_kegiatan"`
	CKodeItemAnggaran  string    `gorm:"column:c_kode_item_anggaran;primaryKey" json:"c_kode_item_anggaran"`
	CTanggalPakai      time.Time `gorm:"column:c_tanggal_pakai;primaryKey" json:"c_tanggal_pakai"`
	CTanggalPakaiAcc   time.Time `gorm:"column:c_tanggal_pakai_acc" json:"c_tanggal_pakai_acc"`
	CPengajuan         float64   `gorm:"column:c_pengajuan;not null" json:"c_pengajuan"`
	CSetujuKacab       float64   `gorm:"column:c_setuju_kacab" json:"c_setuju_kacab"`
	CSetujuBidang      float64   `gorm:"column:c_setuju_bidang" json:"c_setuju_bidang"`
	CSetujuLog         float64   `gorm:"column:c_setuju_log" json:"c_setuju_log"`
	CSetujuAkuntansi   float64   `gorm:"column:c_setuju_akuntansi" json:"c_setuju_akuntansi"`
	CStatus            string    `gorm:"column:c_status;default:Create" json:"c_status"`
	CDokumen           string    `gorm:"column:c_dokumen;default:NULL" json:"c_dokumen"`
	CIDVendor          int32     `gorm:"column:c_id_vendor;primaryKey" json:"c_id_vendor"`
	CFlagSpta          int16     `gorm:"column:c_flag_spta" json:"c_flag_spta"`
	CFlagDrt           int16     `gorm:"column:c_flag_drt" json:"c_flag_drt"`
	CUpdater           string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate        time.Time `gorm:"column:c_last_update;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TKebutuhanAnggaranSplit's table name
func (*TKebutuhanAnggaranSplit) TableName() string {
	return TableNameTKebutuhanAnggaranSplit
}
