// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTKegiatanAcuan = "t_kegiatan_acuan"

// TKegiatanAcuan mapped from table <t_kegiatan_acuan>
type TKegiatanAcuan struct {
	CIDKegiatanAcuan       int32     `gorm:"column:c_id_kegiatan_acuan;primaryKey;autoIncrement:true" json:"c_id_kegiatan_acuan"`
	CKodeAcuan             string    `gorm:"column:c_kode_acuan;default:NULL" json:"c_kode_acuan"`
	CIDKelompokKegiatan    int32     `gorm:"column:c_id_kelompok_kegiatan" json:"c_id_kelompok_kegiatan"`
	CBidangPenanggungJawab int32     `gorm:"column:c_bidang_penanggung_jawab" json:"c_bidang_penanggung_jawab"`
	CBidangTerkait         string    `gorm:"column:c_bidang_terkait" json:"c_bidang_terkait"`
	CUpline                int32     `gorm:"column:c_upline" json:"c_upline"`
	CNamaKegiatan          string    `gorm:"column:c_nama_kegiatan;default:NULL" json:"c_nama_kegiatan"`
	CTanggalAwal           time.Time `gorm:"column:c_tanggal_awal;not null;default:1970-01-01" json:"c_tanggal_awal"`
	CTanggalAkhir          time.Time `gorm:"column:c_tanggal_akhir;default:1970-01-01" json:"c_tanggal_akhir"`
	CIsPromosi             string    `gorm:"column:c_is_promosi;default:NULL" json:"c_is_promosi"`
	CStatus                string    `gorm:"column:c_status;default:NULL" json:"c_status"`
	CIsKontrol             int16     `gorm:"column:c_is_kontrol;not null" json:"c_is_kontrol"`
	CUpdater               string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate            time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
	CTahunAjaran           string    `gorm:"column:c_tahun_ajaran;not null;default:2018/2019" json:"c_tahun_ajaran"`
}

// TableName TKegiatanAcuan's table name
func (*TKegiatanAcuan) TableName() string {
	return TableNameTKegiatanAcuan
}
