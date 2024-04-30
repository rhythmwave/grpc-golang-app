// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTDokumenTransfer = "t_dokumen_transfer"

// TDokumenTransfer mapped from table <t_dokumen_transfer>
type TDokumenTransfer struct {
	CIDDokumenTransfer string    `gorm:"column:c_id_dokumen_transfer;primaryKey" json:"c_id_dokumen_transfer"`
	CMingguKe          int16     `gorm:"column:c_minggu_ke;not null" json:"c_minggu_ke"`
	CDeskripsi         string    `gorm:"column:c_deskripsi" json:"c_deskripsi"`
	CStatus            string    `gorm:"column:c_status;default:Create" json:"c_status"`
	CJumlahTransfer    float64   `gorm:"column:c_jumlah_transfer" json:"c_jumlah_transfer"`
	CPetugasTransfer   string    `gorm:"column:c_petugas_transfer;default:NULL" json:"c_petugas_transfer"`
	CTanggalTransfer   time.Time `gorm:"column:c_tanggal_transfer" json:"c_tanggal_transfer"`
	CBuktiTransfer     string    `gorm:"column:c_bukti_transfer;default:NULL" json:"c_bukti_transfer"`
	CUpdater           string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CInserter          string    `gorm:"column:c_inserter;default:NULL" json:"c_inserter"`
	CLastUpdate        time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TDokumenTransfer's table name
func (*TDokumenTransfer) TableName() string {
	return TableNameTDokumenTransfer
}
