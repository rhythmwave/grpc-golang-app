// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTKebutuhanBarangSplit = "t_kebutuhan_barang_split"

// TKebutuhanBarangSplit mapped from table <t_kebutuhan_barang_split>
type TKebutuhanBarangSplit struct {
	CKodeSplitBarang string    `gorm:"column:c_kode_split_barang;primaryKey" json:"c_kode_split_barang"`
	CKodeKegiatan    string    `gorm:"column:c_kode_kegiatan;primaryKey" json:"c_kode_kegiatan"`
	CKodeBarang      string    `gorm:"column:c_kode_barang;primaryKey" json:"c_kode_barang"`
	CTanggalPakai    time.Time `gorm:"column:c_tanggal_pakai;primaryKey" json:"c_tanggal_pakai"`
	CTanggalPakaiAcc time.Time `gorm:"column:c_tanggal_pakai_acc" json:"c_tanggal_pakai_acc"`
	CPengajuan       float64   `gorm:"column:c_pengajuan;not null" json:"c_pengajuan"`
	CSetujuKacab     float64   `gorm:"column:c_setuju_kacab" json:"c_setuju_kacab"`
	CSetujuBidang    float64   `gorm:"column:c_setuju_bidang" json:"c_setuju_bidang"`
	CSetujuLog       float64   `gorm:"column:c_setuju_log" json:"c_setuju_log"`
	CSetujuAkuntansi float64   `gorm:"column:c_setuju_akuntansi" json:"c_setuju_akuntansi"`
	CStatus          string    `gorm:"column:c_status;default:Create" json:"c_status"`
	CDokumen         string    `gorm:"column:c_dokumen;default:NULL" json:"c_dokumen"`
	CIDVendor        int32     `gorm:"column:c_id_vendor;primaryKey" json:"c_id_vendor"`
	CIsPusat         int16     `gorm:"column:c_is_pusat" json:"c_is_pusat"`
	CFlagSpta        int16     `gorm:"column:c_flag_spta" json:"c_flag_spta"`
	CFlagDrt         int16     `gorm:"column:c_flag_drt" json:"c_flag_drt"`
	CUpdater         string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate      time.Time `gorm:"column:c_last_update;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TKebutuhanBarangSplit's table name
func (*TKebutuhanBarangSplit) TableName() string {
	return TableNameTKebutuhanBarangSplit
}