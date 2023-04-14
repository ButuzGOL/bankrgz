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

func ReadBankBranch(c *gin.Context) {
	var DB = database.ConnectDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var collection = getcollection.GetCollection(DB, "BankBranches")

	id := c.Param("bankBranchId")

	var result model.BankBranch

	objId, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&result)

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": res})
}
