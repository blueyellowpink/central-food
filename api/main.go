package main

import (
	"github.com/gin-gonic/gin"

	"api/internal/router"
	"api/internal/state"
	"db"
)

func main() {
	gin.SetMode(gin.DebugMode)

	dbConfig := db.GormDatabaseConfig{
		Host:                   "localhost",
		User:                   "postgres",
		Password:               "example",
		Port:                   5432,
		DbName:                 "test",
		SslMode:                false,
		DisableImplicitPrepare: true,
	}
	app := state.AppState{Db: db.NewGormDatabase(dbConfig)}

	router := router.New(&app)
	router.Run("0.0.0.0:5000")
}
