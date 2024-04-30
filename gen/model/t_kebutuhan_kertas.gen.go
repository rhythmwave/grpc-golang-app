// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTKebutuhanKerta = "t_kebutuhan_kertas"

// TKebutuhanKerta mapped from table <t_kebutuhan_kertas>
type TKebutuhanKerta struct {
	CIDKebutuhanKertas   int32     `gorm:"column:c_id_kebutuhan_kertas;primaryKey;autoIncrement:true" json:"c_id_kebutuhan_kertas"`
	CIDPusatRisso        int32     `gorm:"column:c_id_pusat_risso" json:"c_id_pusat_risso"`
	CIDBidang            int32     `gorm:"column:c_id_bidang" json:"c_id_bidang"`
	CIDKegiatanAcuan     int32     `gorm:"column:c_id_kegiatan_acuan" json:"c_id_kegiatan_acuan"`
	CIDKebutuhan         int32     `gorm:"column:c_id_kebutuhan" json:"c_id_kebutuhan"`
	CIDJenisKertas       int32     `gorm:"column:c_id_jenis_kertas" json:"c_id_jenis_kertas"`
	CPembagiKertas       int32     `gorm:"column:c_pembagi_kertas" json:"c_pembagi_kertas"`
	CPengaliKertas       int32     `gorm:"column:c_pengali_kertas" json:"c_pengali_kertas"`
	CTargetKeikutsertaan int32     `gorm:"column:c_target_keikutsertaan" json:"c_target_keikutsertaan"`
	CIsEvent             string    `gorm:"column:c_is_event;default:NULL" json:"c_is_event"`
	CTanggalAwal         time.Time `gorm:"column:c_tanggal_awal" json:"c_tanggal_awal"`
	CTanggalAkhir        time.Time `gorm:"column:c_tanggal_akhir" json:"c_tanggal_akhir"`
	CJumlah              int32     `gorm:"column:c_jumlah" json:"c_jumlah"`
	CJumlahAkhir         int32     `gorm:"column:c_jumlah_akhir" json:"c_jumlah_akhir"`
	CUpdater             string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate          time.Time `gorm:"column:c_last_update" json:"c_last_update"`
}

// TableName TKebutuhanKerta's table name
func (*TKebutuhanKerta) TableName() string {
	return TableNameTKebutuhanKerta
}
