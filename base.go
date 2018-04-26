package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/dummy_db")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	type UserReview struct {
		Id 			int
		OrderId 	string
		ProductId 	string
		UserId 		string
		Rating 		float64
		Review 		string
		CreatedAt 	time.Time
		UpdatedAt 	time.Time
	}
	router := gin.Default()

	// GET a userReview detail
	router.GET("/user-review/:id", func(c *gin.Context) {
		var (
			userReview UserReview
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("SELECT * from user_review where id = ?;", id)
		err = row.Scan(&userReview.Id, &userReview.OrderId, &userReview.ProductId,
		 			&userReview.ProductId, &userReview.Rating, &userReview.Review,
		 			&userReview.CreatedAt, &userReview.UpdatedAt)
		
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": userReview,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	router.Run(":3000")
}