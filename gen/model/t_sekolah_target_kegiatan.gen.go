// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTSekolahTargetKegiatan = "t_sekolah_target_kegiatan"

// TSekolahTargetKegiatan mapped from table <t_sekolah_target_kegiatan>
type TSekolahTargetKegiatan struct {
	CIDSekolah       int32 `gorm:"column:c_id_sekolah;primaryKey" json:"c_id_sekolah"`
	CIDTingkat       int16 `gorm:"column:c_id_tingkat;primaryKey" json:"c_id_tingkat"`
	CIDKegiatan      int32 `gorm:"column:c_id_kegiatan;primaryKey" json:"c_id_kegiatan"`
	CTargetEvent     int16 `gorm:"column:c_target_event;not null" json:"c_target_event"`
	CTargetKehadiran int16 `gorm:"column:c_target_kehadiran;not null" json:"c_target_kehadiran"`
}

// TableName TSekolahTargetKegiatan's table name
func (*TSekolahTargetKegiatan) TableName() string {
	return TableNameTSekolahTargetKegiatan
}