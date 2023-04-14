package routes

import (
	getcollection "bankbranches/Collection"
	model "bankbranches/Model"
	database "bankbranches/databases"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateBankBranch(c *gin.Context) {
	var DB = database.ConnectDB()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var collection = getcollection.GetCollection(DB, "BankBranches")

	id := c.Param("bankBranchId")
	var bankBranch model.BankBranch

	if err := c.ShouldBindJSON(&bankBranch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	edited := bson.M{
		"number":   bankBranch.Number,
		"district": bankBranch.District,
		"phone":    bankBranch.Phone,
		"address":  bankBranch.Address,
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": edited})

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": res})
}
