package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api/internal/state"
)

type AuthController struct {
	svc AuthService
}

func AddAuthRoutes(rg *gin.RouterGroup, app *state.AppState) {
	group := rg.Group("/auth")

	repo := AuthRepoImpl{db: app.Db}
	service := AuthServiceImpl{repo: &repo}
	controller := AuthController{svc: &service}

	group.POST("/sign-up", controller.signUp)
	group.POST("/sign-in", controller.signIn)
}

type SignUp struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Email string `json:"email" binding:"required,isEmail"`
}

func (c *AuthController) signUp(ctx *gin.Context) {
    var params SignUp

    if err := ctx.BindJSON(&params); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }

    if _, err := c.svc.signUp(&params); err != nil {
        ctx.JSON(http.StatusBadRequest, "Error")
        return
    }

	ctx.JSON(http.StatusOK, params)
}

type SignIn struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func (c *AuthController) signIn(ctx *gin.Context) {
    var params SignIn

    if err := ctx.BindJSON(&params); err != nil {
        ctx.JSON(http.StatusBadRequest, "Error Bind JSON")
        return
    }

	// c.svc.signIn()
	ctx.JSON(http.StatusOK, params)
}
