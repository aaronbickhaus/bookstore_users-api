package app

import (
	"github.com/aaronbickhaus/bookstore_users-api/controllers/ping"
	"github.com/aaronbickhaus/bookstore_users-api/controllers/user"
)
func mapUrls() {
router.GET("/ping", ping.Ping)
router.GET("/users/:user_id", user.Get)
router.POST("/users", user.Create)
router.GET("/internal/users/search", user.Search)
router.POST("/users/:user_id", user.Update)
router.DELETE("/users/:user_id", user.Delete)
router.PATCH("/users/:user_id", user.Update)
}
