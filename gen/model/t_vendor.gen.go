// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTVendor = "t_vendor"

// TVendor mapped from table <t_vendor>
type TVendor struct {
	CIDVendor      int32     `gorm:"column:c_id_vendor;primaryKey;autoIncrement:true" json:"c_id_vendor"`
	CNamaVendor    string    `gorm:"column:c_nama_vendor;default:NULL" json:"c_nama_vendor"`
	CAlamatVendor  string    `gorm:"column:c_alamat_vendor;default:NULL" json:"c_alamat_vendor"`
	CJenisVendor   string    `gorm:"column:c_jenis_vendor;default:NULL" json:"c_jenis_vendor"`
	CContactPerson string    `gorm:"column:c_contact_person;default:NULL" json:"c_contact_person"`
	CNoTelepon     string    `gorm:"column:c_no_telepon;default:0" json:"c_no_telepon"`
	CIsPpkp        int32     `gorm:"column:c_is_ppkp" json:"c_is_ppkp"`
	CNoNpwp        string    `gorm:"column:c_no_npwp;default:NULL" json:"c_no_npwp"`
	CNamaNpwp      string    `gorm:"column:c_nama_npwp;default:NULL" json:"c_nama_npwp"`
	CStatus        string    `gorm:"column:c_status;default:NULL" json:"c_status"`
	CUpdater       string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate    time.Time `gorm:"column:c_last_update;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TVendor's table name
func (*TVendor) TableName() string {
	return TableNameTVendor
}