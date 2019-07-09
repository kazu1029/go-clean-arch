package external

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kazu1029/go-clean-arch/adapter/controllers"
	"github.com/kazu1029/go-clean-arch/external/mysql"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	logger := &Logger{}
	conn := mysql.Connect()
	userController := controllers.NewUserController(conn, logger)

	router.GET("/hc", func(c *gin.Context) { c.String(http.StatusOK, "hello") })
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })

	Router = router
}
