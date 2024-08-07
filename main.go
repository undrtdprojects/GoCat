package main

import (
	"GoCat/configs"
	"GoCat/databases/connection"
	"GoCat/databases/migration"
	"GoCat/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Initiator()

	connection.Initiator()
	defer connection.DBConnections.Close()

	migration.Initiator(connection.DBConnections)

	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()

	category.Initiator(router)
	book.Initiator(router)
	user.Initiator(router)

	router.Run(":8080")
}
