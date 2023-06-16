package models

import (
	"gorm.io/gorm"

	database "db"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

var _ database.Internal = (*Person)(nil)

func (m *Person) Internal() {}
