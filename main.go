package main

import (
	"GoCat/configs"
	"GoCat/databases/connection"
	"GoCat/databases/migration"
	"GoCat/modules/categories"
	"GoCat/modules/menu"
	"GoCat/modules/payment"
	"GoCat/modules/role"
	"GoCat/modules/transaction0"
	"GoCat/modules/transaction1"
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

	categories.Initiator(router)
	menu.Initiator(router)
	payment.Initiator(router)
	role.Initiator(router)
	transaction0.Initiator(router)
	transaction1.Initiator(router)
	user.Initiator(router)

	router.Run(":8080")
}
