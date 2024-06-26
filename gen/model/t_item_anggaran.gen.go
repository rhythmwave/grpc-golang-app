// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTItemAnggaran = "t_item_anggaran"

// TItemAnggaran mapped from table <t_item_anggaran>
type TItemAnggaran struct {
	CKodeItemAnggaran string    `gorm:"column:c_kode_item_anggaran;primaryKey" json:"c_kode_item_anggaran"`
	CNamaItemAnggaran string    `gorm:"column:c_nama_item_anggaran;default:NULL" json:"c_nama_item_anggaran"`
	CIsDokumen        int16     `gorm:"column:c_is_dokumen" json:"c_is_dokumen"`
	CIDKelompokPajak  string    `gorm:"column:c_id_kelompok_pajak;default:0" json:"c_id_kelompok_pajak"`
	CKodeCOA          string    `gorm:"column:c_kode_c_o_a;default:NULL" json:"c_kode_c_o_a"`
	CKeterangan       string    `gorm:"column:c_keterangan" json:"c_keterangan"`
	CStatus           string    `gorm:"column:c_status;default:Aktif" json:"c_status"`
	CUpdater          string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate       time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TItemAnggaran's table name
func (*TItemAnggaran) TableName() string {
	return TableNameTItemAnggaran
}
