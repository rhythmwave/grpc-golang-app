// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTInfoProfil = "t_info_profil"

// TInfoProfil mapped from table <t_info_profil>
type TInfoProfil struct {
	CUpline      int32  `gorm:"column:c_upline;primaryKey" json:"c_upline"`
	CIDSekolah   int32  `gorm:"column:c_id_sekolah;primaryKey" json:"c_id_sekolah"`
	CTahunAjaran string `gorm:"column:c_tahun_ajaran;primaryKey" json:"c_tahun_ajaran"`
	CInfo        string `gorm:"column:c_info" json:"c_info"`
}

// TableName TInfoProfil's table name
func (*TInfoProfil) TableName() string {
	return TableNameTInfoProfil
}
