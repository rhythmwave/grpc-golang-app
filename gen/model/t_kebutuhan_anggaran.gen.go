// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTKebutuhanAnggaran = "t_kebutuhan_anggaran"

// TKebutuhanAnggaran mapped from table <t_kebutuhan_anggaran>
type TKebutuhanAnggaran struct {
	CKodeKegiatan     string    `gorm:"column:c_kode_kegiatan;primaryKey" json:"c_kode_kegiatan"`
	CKodeItemAnggaran string    `gorm:"column:c_kode_item_anggaran;primaryKey" json:"c_kode_item_anggaran"`
	CJumlah           float64   `gorm:"column:c_jumlah" json:"c_jumlah"`
	CNilaiSatuan      float64   `gorm:"column:c_nilai_satuan" json:"c_nilai_satuan"`
	CTotal            float64   `gorm:"column:c_total" json:"c_total"`
	CIDVendor         int32     `gorm:"column:c_id_vendor" json:"c_id_vendor"`
	CDokumen          string    `gorm:"column:c_dokumen;default:NULL" json:"c_dokumen"`
	CStatus           string    `gorm:"column:c_status;default:Create" json:"c_status"`
	CAcc              string    `gorm:"column:c_acc" json:"c_acc"`
	CUpdater          string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate       time.Time `gorm:"column:c_last_update;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TKebutuhanAnggaran's table name
func (*TKebutuhanAnggaran) TableName() string {
	return TableNameTKebutuhanAnggaran
}
