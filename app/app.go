package app

import (
	"github.com/aaronbickhaus/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var  (
	router = gin.Default()
)

func StartApplication() {
mapUrls()
logger.Info("application starting")
router.Run(":8080")

}
