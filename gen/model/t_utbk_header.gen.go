// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTUtbkHeader = "t_utbk_header"

// TUtbkHeader mapped from table <t_utbk_header>
type TUtbkHeader struct {
	CNis        string    `gorm:"column:c_nis;default:NULL" json:"c_nis"`
	CNoUtbk     string    `gorm:"column:c_no_utbk;primaryKey" json:"c_no_utbk"`
	CKodeGedung int32     `gorm:"column:c_kode_gedung" json:"c_kode_gedung"`
	CNamaSiswa  string    `gorm:"column:c_nama_siswa;default:NULL" json:"c_nama_siswa"`
	CNilai      float64   `gorm:"column:c_nilai" json:"c_nilai"`
	CKelompok   string    `gorm:"column:c_kelompok;default:NULL" json:"c_kelompok"`
	CSesi       string    `gorm:"column:c_sesi;primaryKey" json:"c_sesi"`
	CTglUpload  time.Time `gorm:"column:c_tgl_upload" json:"c_tgl_upload"`
	CTglTampil  time.Time `gorm:"column:c_tgl_tampil" json:"c_tgl_tampil"`
	CTglTutup   time.Time `gorm:"column:c_tgl_tutup" json:"c_tgl_tutup"`
	CEditNilai  string    `gorm:"column:c_edit_nilai;default:0" json:"c_edit_nilai"`
	CUpdater    string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
}

// TableName TUtbkHeader's table name
func (*TUtbkHeader) TableName() string {
	return TableNameTUtbkHeader
}
