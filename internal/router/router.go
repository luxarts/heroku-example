package router
import "github.com/gin-gonic/gin"

func Init() *gin.Engine{
	router := gin.Default()

	mapEndpoints(router)

	return router
}
