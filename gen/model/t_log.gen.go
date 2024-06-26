// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTLog = "t_log"

// TLog mapped from table <t_log>
type TLog struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CIDLogin     int32     `gorm:"column:c_id_login" json:"c_id_login"`
	CNik         string    `gorm:"column:c_nik;default:NULL" json:"c_nik"`
	CNamaLengkap string    `gorm:"column:c_nama_lengkap;default:NULL" json:"c_nama_lengkap"`
	CUserName    string    `gorm:"column:c_user_name;default:NULL" json:"c_user_name"`
	CTglLogin    time.Time `gorm:"column:c_tgl_login;not null;default:CURRENT_TIMESTAMP" json:"c_tgl_login"`
	CTglLogout   time.Time `gorm:"column:c_tgl_logout" json:"c_tgl_logout"`
}

// TableName TLog's table name
func (*TLog) TableName() string {
	return TableNameTLog
}
