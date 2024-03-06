package main

import (
	db "bankassignment/database"
	"bankassignment/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("New Bank assignment")
	pg_db := db.ConnectToDB()
	fmt.Println(pg_db)

	server := gin.Default()
	routes.BankRoutes(server)
	routes.BranchRoutes(server)
	routes.AccountRoutes(server)
	routes.CustomerRoutes(server)
	routes.TransactionRoutes(server)
	routes.MappingRoutes(server)
	server.Run(":3000")
}
