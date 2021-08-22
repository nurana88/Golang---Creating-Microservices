package app

import (
	controllers "github.com/nurana/microservices/controllers/ping"
	"github.com/nurana/microservices/controllers/users"
)

func MapUrls() {
	router.GET("/ping", controllers.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login",users.Login)
}
