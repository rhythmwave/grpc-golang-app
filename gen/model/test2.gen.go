// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTest2 = "test2"

// Test2 mapped from table <test2>
type Test2 struct {
	ID      int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	IDTest1 int32 `gorm:"column:id_test1;not null" json:"id_test1"`
}

// TableName Test2's table name
func (*Test2) TableName() string {
	return TableNameTest2
}
