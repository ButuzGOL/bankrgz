package routes

import (
	getcollection "bankbranches/Collection"
	model "bankbranches/Model"
	database "bankbranches/databases"
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

	var collection = getcollection.GetCollection(DB, "BankBranches")

	var bankBranch model.BankBranch
	if err := c.ShouldBindJSON(&bankBranch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	payload := model.BankBranch{
		ID:       primitive.NewObjectID(),
		Number:   bankBranch.Number,
		District: bankBranch.District,
		Phone:    bankBranch.Phone,
		Address:  bankBranch.Address,
	}

	result, err := collection.InsertOne(ctx, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created successfully", "Data": map[string]interface{}{"data": result}})
}
