package services

import (
	"net/http"
	"task/deeSeeComics/data"
	"task/deeSeeComics/domain/models"

	"github.com/gin-gonic/gin"
)

func GetSuperheroes(ctx *gin.Context) {
	filter := ctx.Query("filter")
	var entities = data.SuperHeroes(filter)
	var payload []models.Superhero
	for i := range entities {
		payload = append(payload, entities[i].MapToModel(false, "name", "identity", "birthday", "superpowers"))
	}
	ctx.IndentedJSON(http.StatusOK, payload)
}

func GetEncryptedSuperheroes(ctx *gin.Context) {
	filter := ctx.Query("filter")
	var entities = data.SuperHeroes(filter)
	var payload []models.Superhero
	for i := range entities {
		payload = append(payload, entities[i].MapToModel(true, "name", "identity", "birthday", "superpowers"))
	}
	ctx.IndentedJSON(http.StatusOK, payload)
}
