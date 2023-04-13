package routes

import (
	getcollection "bankrgz/Collection"
	model "bankrgz/Model"
	database "bankrgz/databases"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBankBranch(c *gin.Context) {
	var DB = database.ConnectDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var bankBankConnection = getcollection.GetCollection(DB, "BankBranches")

	var bankBranch model.BankBranch

	if err := c.ShouldBindJSON(&bankBranch); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	bankBranchPayload := model.BankBranch{
		ID:       primitive.NewObjectID(),
		Number:   bankBranch.Number,
		District: bankBranch.District,
		Phone:    bankBranch.Phone,
		Address:  bankBranch.Address,
	}

	result, err := bankBankConnection.InsertOne(ctx, bankBranchPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}
