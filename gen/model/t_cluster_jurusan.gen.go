// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTClusterJurusan = "t_cluster_jurusan"

// TClusterJurusan mapped from table <t_cluster_jurusan>
type TClusterJurusan struct {
	CKodeClusterJurusan int32     `gorm:"column:c_kode_cluster_jurusan;primaryKey;autoIncrement:true" json:"c_kode_cluster_jurusan"`
	CNamaCluster        string    `gorm:"column:c_nama_cluster" json:"c_nama_cluster"`
	CStatus             string    `gorm:"column:c_status;default:Aktif" json:"c_status"`
	CUpdater            string    `gorm:"column:c_updater" json:"c_updater"`
	CCreatedAt          time.Time `gorm:"column:c_created_at;default:CURRENT_TIMESTAMP" json:"c_created_at"`
	CLastUpdate         time.Time `gorm:"column:c_last_update;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TClusterJurusan's table name
func (*TClusterJurusan) TableName() string {
	return TableNameTClusterJurusan
}