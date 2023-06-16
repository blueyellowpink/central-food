package state

import (
    "gorm.io/gorm"

    "db"
)

type AppState struct {
	Db db.Database[*gorm.DB]
}
