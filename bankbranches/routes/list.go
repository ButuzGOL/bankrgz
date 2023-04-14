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
)

// function which returns a list of all bank branches
func ListBankBranch(c *gin.Context) {
	var DB = database.ConnectDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var collection = getcollection.GetCollection(DB, "BankBranches")

	var results []*model.BankBranch

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	for cur.Next(ctx) {
		var result model.BankBranch
		err := cur.Decode(&result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		results = append(results, &result)
	}

	res := map[string]interface{}{"data": results}

	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": res})
}
