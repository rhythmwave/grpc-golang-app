// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTBiayaSetujuBidang = "t_biaya_setuju_bidang"

// TBiayaSetujuBidang mapped from table <t_biaya_setuju_bidang>
type TBiayaSetujuBidang struct {
	CIDKegiatan              int32     `gorm:"column:c_id_kegiatan;not null" json:"c_id_kegiatan"`
	CIDJenis                 int32     `gorm:"column:c_id_jenis;not null" json:"c_id_jenis"`
	CBiayaSetuju             int32     `gorm:"column:c_biaya_setuju;not null" json:"c_biaya_setuju"`
	CTanggalPersetujuanBiaya time.Time `gorm:"column:c_tanggal_persetujuan_biaya;not null" json:"c_tanggal_persetujuan_biaya"`
	CKodePengajuanAnggaran   string    `gorm:"column:c_kode_pengajuan_anggaran;default:NULL" json:"c_kode_pengajuan_anggaran"`
	CUpdater                 string    `gorm:"column:c_updater;not null" json:"c_updater"`
	CLastUpdate              time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TBiayaSetujuBidang's table name
func (*TBiayaSetujuBidang) TableName() string {
	return TableNameTBiayaSetujuBidang
}
