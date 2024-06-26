// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTDiskusiPesertaStrategicMeeting = "t_diskusi_peserta_strategic_meeting"

// TDiskusiPesertaStrategicMeeting mapped from table <t_diskusi_peserta_strategic_meeting>
type TDiskusiPesertaStrategicMeeting struct {
	CNik        string    `gorm:"column:c_nik;primaryKey" json:"c_nik"`
	CBidang     string    `gorm:"column:c_bidang;primaryKey" json:"c_bidang"`
	CNikPenilai string    `gorm:"column:c_nik_penilai;primaryKey" json:"c_nik_penilai"`
	CNilai      int32     `gorm:"column:c_nilai;not null;comment:1-10" json:"c_nilai"` // 1-10
	CKeterangan string    `gorm:"column:c_keterangan" json:"c_keterangan"`
	CLastUpdate time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TDiskusiPesertaStrategicMeeting's table name
func (*TDiskusiPesertaStrategicMeeting) TableName() string {
	return TableNameTDiskusiPesertaStrategicMeeting
}
