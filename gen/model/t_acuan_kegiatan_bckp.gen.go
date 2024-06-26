// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTAcuanKegiatanBckp = "t_acuan_kegiatan_bckp"

// TAcuanKegiatanBckp mapped from table <t_acuan_kegiatan_bckp>
type TAcuanKegiatanBckp struct {
	CKodeAcuan             string    `gorm:"column:c_kode_acuan" json:"c_kode_acuan"`
	CKodeKelBarang         string    `gorm:"column:c_kode_kel_barang" json:"c_kode_kel_barang"`
	CKodeKelAnggaran       string    `gorm:"column:c_kode_kel_anggaran" json:"c_kode_kel_anggaran"`
	CKodeKepanitiaan       string    `gorm:"column:c_kode_kepanitiaan" json:"c_kode_kepanitiaan"`
	CBidangPenanggungJawab int32     `gorm:"column:c_bidang_penanggung_jawab" json:"c_bidang_penanggung_jawab"`
	CBidangTerkait         string    `gorm:"column:c_bidang_terkait" json:"c_bidang_terkait"`
	CNamaKegiatan          string    `gorm:"column:c_nama_kegiatan" json:"c_nama_kegiatan"`
	CUpline                string    `gorm:"column:c_upline" json:"c_upline"`
	CIsPromosi             int16     `gorm:"column:c_is_promosi" json:"c_is_promosi"`
	CIsPusat               int16     `gorm:"column:c_is_pusat" json:"c_is_pusat"`
	CStatus                string    `gorm:"column:c_status" json:"c_status"`
	CUpdater               string    `gorm:"column:c_updater" json:"c_updater"`
	CLastUpdate            time.Time `gorm:"column:c_last_update" json:"c_last_update"`
}

// TableName TAcuanKegiatanBckp's table name
func (*TAcuanKegiatanBckp) TableName() string {
	return TableNameTAcuanKegiatanBckp
}
