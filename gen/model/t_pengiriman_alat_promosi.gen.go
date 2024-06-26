// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTPengirimanAlatPromosi = "t_pengiriman_alat_promosi"

// TPengirimanAlatPromosi mapped from table <t_pengiriman_alat_promosi>
type TPengirimanAlatPromosi struct {
	CIDAlatPromosi int32     `gorm:"column:c_id_alat_promosi;not null" json:"c_id_alat_promosi"`
	CIDPenanda     int32     `gorm:"column:c_id_penanda;not null" json:"c_id_penanda"`
	CJumlahKirim   int32     `gorm:"column:c_jumlah_kirim;not null" json:"c_jumlah_kirim"`
	CTanggalKirim  time.Time `gorm:"column:c_tanggal_kirim;not null" json:"c_tanggal_kirim"`
	CUpdater       string    `gorm:"column:c_updater;not null" json:"c_updater"`
	CLastUpdate    time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TPengirimanAlatPromosi's table name
func (*TPengirimanAlatPromosi) TableName() string {
	return TableNameTPengirimanAlatPromosi
}
