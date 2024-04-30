// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTDppbIsi = "t_dppb_isi"

// TDppbIsi mapped from table <t_dppb_isi>
type TDppbIsi struct {
	CIDDppbIsi          int32     `gorm:"column:c_id_dppb_isi;primaryKey;autoIncrement:true" json:"c_id_dppb_isi"`
	CIDDppb             string    `gorm:"column:c_id_dppb;default:NULL" json:"c_id_dppb"`
	CIDKegiatan         int32     `gorm:"column:c_id_kegiatan" json:"c_id_kegiatan"`
	CIDVendor           int32     `gorm:"column:c_id_vendor" json:"c_id_vendor"`
	CKodeBarang         string    `gorm:"column:c_kode_barang;default:NULL" json:"c_kode_barang"`
	CNamaBarang         string    `gorm:"column:c_nama_barang;default:NULL" json:"c_nama_barang"`
	CSpesifikasi        string    `gorm:"column:c_spesifikasi;default:NULL" json:"c_spesifikasi"`
	CSatuan             string    `gorm:"column:c_satuan;default:NULL" json:"c_satuan"`
	CJumlahUsul         float64   `gorm:"column:c_jumlah_usul" json:"c_jumlah_usul"`
	CJumlahSetujuBidang float64   `gorm:"column:c_jumlah_setuju_bidang" json:"c_jumlah_setuju_bidang"`
	CJumlahSetujuLog    float64   `gorm:"column:c_jumlah_setuju_log" json:"c_jumlah_setuju_log"`
	CHargaVendor        float64   `gorm:"column:c_harga_vendor" json:"c_harga_vendor"`
	CHargaDisetujui     float64   `gorm:"column:c_harga_disetujui" json:"c_harga_disetujui"`
	CSumberPengadaan    string    `gorm:"column:c_sumber_pengadaan;default:NULL" json:"c_sumber_pengadaan"`
	CStatus             string    `gorm:"column:c_status;default:NULL" json:"c_status"`
	CUpdater            string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate         time.Time `gorm:"column:c_last_update" json:"c_last_update"`
	CStatusPengajuan    string    `gorm:"column:c_status_pengajuan;not null;default:RENCANA" json:"c_status_pengajuan"`
}

// TableName TDppbIsi's table name
func (*TDppbIsi) TableName() string {
	return TableNameTDppbIsi
}
