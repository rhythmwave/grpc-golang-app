// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTVendorGedung = "t_vendor_gedung"

// TVendorGedung mapped from table <t_vendor_gedung>
type TVendorGedung struct {
	CIDVendor   int32     `gorm:"column:c_id_vendor;primaryKey" json:"c_id_vendor"`
	CIDGedung   int32     `gorm:"column:c_id_gedung;primaryKey" json:"c_id_gedung"`
	CUpdater    string    `gorm:"column:c_updater;default:NULL" json:"c_updater"`
	CLastUpdate time.Time `gorm:"column:c_last_update;default:CURRENT_TIMESTAMP" json:"c_last_update"`
}

// TableName TVendorGedung's table name
func (*TVendorGedung) TableName() string {
	return TableNameTVendorGedung
}