// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTest1 = "test1"

// Test1 mapped from table <test1>
type Test1 struct {
	ID int32 `gorm:"column:id" json:"id"`
}

// TableName Test1's table name
func (*Test1) TableName() string {
	return TableNameTest1
}
