// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTTahapTransferAnggaran = "t_tahap_transfer_anggaran"

// TTahapTransferAnggaran mapped from table <t_tahap_transfer_anggaran>
type TTahapTransferAnggaran struct {
	CIDKegiatan      int32     `gorm:"column:c_id_kegiatan;primaryKey" json:"c_id_kegiatan"`
	CIDJenisBiaya    string    `gorm:"column:c_id_jenis_biaya;primaryKey" json:"c_id_jenis_biaya"`
	CKelompok        string    `gorm:"column:c_kelompok;primaryKey;default:ANGGARAN" json:"c_kelompok"`
	CIDDppbIsi       int32     `gorm:"column:c_id_dppb_isi;primaryKey" json:"c_id_dppb_isi"`
	CTahapKe         int32     `gorm:"column:c_tahap_ke;primaryKey" json:"c_tahap_ke"`
	CJumlahTransfer  int32     `gorm:"column:c_jumlah_transfer;not null" json:"c_jumlah_transfer"`
	CTanggalTransfer time.Time `gorm:"column:c_tanggal_transfer" json:"c_tanggal_transfer"`
	CUpdater         string    `gorm:"column:c_updater;not null" json:"c_updater"`
	CLastUpdate      time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
	CSumberDana      string    `gorm:"column:c_sumber_dana;default:NULL" json:"c_sumber_dana"`
	CStatus          string    `gorm:"column:c_status;not null;default:Belum Transfer" json:"c_status"`
	CIsTransfer      int16     `gorm:"column:c_is_transfer_;not null" json:"c_is_transfer_"`
}

// TableName TTahapTransferAnggaran's table name
func (*TTahapTransferAnggaran) TableName() string {
	return TableNameTTahapTransferAnggaran
}
