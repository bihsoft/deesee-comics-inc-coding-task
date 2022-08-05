package handlers

import (
	"task/deeSeeComics/routes"
	"task/deeSeeComics/services"

	"github.com/gin-gonic/gin"
)

func SetupRequestHandlers(router gin.Engine) {
	router.GET(routes.GetSuperheroes, services.GetSuperheroes)
	router.GET(routes.GetEncryptedSuperheroes, services.GetEncryptedSuperheroes)
}
