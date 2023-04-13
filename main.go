package main

import (
	routes "bankrgz/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/bankBranches", routes.CreateBankBranch)

	router.Run(":3000")
}
