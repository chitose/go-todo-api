// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "database/sql"

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID          sql.NullString `gorm:"column:id;primaryKey" json:"id"`
	DisplayName sql.NullString `gorm:"column:displayName;not null" json:"displayName"`
	Photo       sql.NullString `gorm:"column:photo;not null" json:"photo"`
	Email       sql.NullString `gorm:"column:email;not null" json:"email"`
	CreatedAt   sql.NullTime   `gorm:"column:createdAt;not null" json:"createdAt"`
	UpdatedAt   sql.NullTime   `gorm:"column:updatedAt;not null" json:"updatedAt"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
