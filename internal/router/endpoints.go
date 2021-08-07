package router

import (
	"github.com/gin-gonic/gin"
	"heroku-example/internal/defines"
	"net/http"
)

func mapEndpoints(r *gin.Engine){
	r.GET(defines.PathsPing, healthCheck)
}

func healthCheck(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "pong")
}
