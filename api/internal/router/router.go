package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api/internal/modules/auth"
	"api/internal/state"
	"api/internal/validator"
)

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func New(app *state.AppState) *gin.Engine {
	r := gin.Default()

	if err := validator.BindValidators(); err != nil {
		panic("can not bind validators")
	}

	r.GET("/ping", ping)

	v1 := r.Group("/v1")
	auth.AddAuthRoutes(v1, app)

	return r
}
