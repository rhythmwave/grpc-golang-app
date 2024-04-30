// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTRencanaKegiatan = "t_rencana_kegiatan"

// TRencanaKegiatan mapped from table <t_rencana_kegiatan>
type TRencanaKegiatan struct {
	CKodeKegiatan       string    `gorm:"column:c_kode_kegiatan;primaryKey" json:"c_kode_kegiatan"`
	CKodeAcuan          string    `gorm:"column:c_kode_acuan;primaryKey" json:"c_kode_acuan"`
	CIDGedung           int32     `gorm:"column:c_id_gedung;not null" json:"c_id_gedung"`
	CNamaKegiatan       string    `gorm:"column:c_nama_kegiatan;not null" json:"c_nama_kegiatan"`
	CUpline             string    `gorm:"column:c_upline;default:NULL" json:"c_upline"`
	CBulanKegiatanAwal  string    `gorm:"column:c_bulan_kegiatan_awal;not null" json:"c_bulan_kegiatan_awal"`
	CBulanKegiatanAkhir string    `gorm:"column:c_bulan_kegiatan_akhir;not null" json:"c_bulan_kegiatan_akhir"`
	CIDTingkat          string    `gorm:"column:c_id_tingkat;default:NULL" json:"c_id_tingkat"`
	CIDTujuanTingkat    string    `gorm:"column:c_id_tujuan_tingkat;default:NULL" json:"c_id_tujuan_tingkat"`
	CTargetKehadiran    int32     `gorm:"column:c_target_kehadiran" json:"c_target_kehadiran"`
	CTargetCapaian      int32     `gorm:"column:c_target_capaian" json:"c_target_capaian"`
	CTahunAjaran        string    `gorm:"column:c_tahun_ajaran;not null" json:"c_tahun_ajaran"`
	CStatus             string    `gorm:"column:c_status;default:Create" json:"c_status"`
	CUpdater            string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate         time.Time `gorm:"column:c_last_update;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TRencanaKegiatan's table name
func (*TRencanaKegiatan) TableName() string {
	return TableNameTRencanaKegiatan
}