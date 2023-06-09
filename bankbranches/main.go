package main

import (
	routes "bankbranches/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/bankBranches", routes.ListBankBranch)
	router.POST("/bankBranches", routes.CreateBankBranch)
	router.GET("/bankBranches/:bankBranchId", routes.ReadBankBranch)
	router.DELETE("/bankBranches/:bankBranchId", routes.DeleteBankBranch)
	router.PUT("/bankBranches/:bankBranchId", routes.UpdateBankBranch)

	router.Run(":7001")
}
