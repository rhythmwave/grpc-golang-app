// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTIsiPermintaanTransfer = "t_isi_permintaan_transfer"

// TIsiPermintaanTransfer mapped from table <t_isi_permintaan_transfer>
type TIsiPermintaanTransfer struct {
	CIDPermintaanTransfer string    `gorm:"column:c_id_permintaan_transfer;primaryKey" json:"c_id_permintaan_transfer"`
	CKodeSplit            string    `gorm:"column:c_kode_split;primaryKey" json:"c_kode_split"`
	CStatus               string    `gorm:"column:c_status;default:Create" json:"c_status"`
	CUpdater              string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate           time.Time `gorm:"column:c_last_update;not null;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TIsiPermintaanTransfer's table name
func (*TIsiPermintaanTransfer) TableName() string {
	return TableNameTIsiPermintaanTransfer
}
