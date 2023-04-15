package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

func main() {
	router := gin.Default()

	router.GET("/bankBranchReviews/:bankBranchId", ListBankBranchReviews)
	router.POST("/bankBranchReviews/:bankBranchId", CreateBankBranchReview)

	router.Run(":5001")
}

type BankBranchesReviews struct {
	BankBranchesReviews []BankBranchReview `json:"bankBranchesReviews"`
}

type BankBranchReview struct {
	Id           string `json:"id"`
	BankBranchId string `json:"bankBranchId"`
	Comment      string `json:"comment"`
	Rating       int32  `json:"rating"`
}

func readReviewFile() (BankBranchesReviews, error) {
	var reviews BankBranchesReviews

	content, err := os.ReadFile("data.json")
	if err != nil {
		return reviews, err
	}

	json.Unmarshal(content, &reviews)

	fmt.Println(reviews)
	return reviews, nil
}

func CreateBankBranchReview(c *gin.Context) {
	id := c.Param("bankBranchId")

	var review BankBranchReview
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data, err := readReviewFile()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	review.BankBranchId = id
	review.Id = guuid.New().String()
	data.BankBranchesReviews = append(data.BankBranchesReviews, review)

	file, _ := json.MarshalIndent(data, "", " ")
	err = os.WriteFile("data.json", file, 0644)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created successfully", "Data": map[string]interface{}{"data": review}})
}

func ListBankBranchReviews(c *gin.Context) {
	id := c.Param("bankBranchId")

	data, err := readReviewFile()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var filteredData []BankBranchReview = []BankBranchReview{}
	for _, v := range data.BankBranchesReviews {
		if v.BankBranchId == id {
			filteredData = append(filteredData, v)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": filteredData})
}
