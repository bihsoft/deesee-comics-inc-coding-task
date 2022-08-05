package main

import (
	"flag"
	"task/deeSeeComics/data"
	"task/deeSeeComics/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	data.Migrate()
	port := flag.String("port", "8080", "port number provided in startup arguments")
	flag.Parse()
	router := gin.Default()
	handlers.SetupRequestHandlers(*router)
	router.Run("localhost:" + *port)
}
