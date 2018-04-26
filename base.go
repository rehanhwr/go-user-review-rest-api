package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// change second parameter root:password to [username]:[password] of yours.
	// parseTime=true is used for changing mysql timestamp to time.Time golang
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/dummy_db?parseTime=true")
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
	// e.g. /user-review/1
	router.GET("/user-review/:id", func(c *gin.Context) {
		var (
			userReview UserReview
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("SELECT * FROM user_review WHERE id = ?;", id)
		err = row.Scan(&userReview.Id, &userReview.OrderId, &userReview.ProductId,
		 			&userReview.UserId, &userReview.Rating, &userReview.Review,
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

	// GET all userReview
	router.GET("/user-reviews", func(c *gin.Context) {
		var (
			userReview  UserReview
			userReviews []UserReview
		)
		rows, err := db.Query("SELECT * FROM user_review;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&userReview.Id, &userReview.OrderId, &userReview.ProductId,
		 			&userReview.UserId, &userReview.Rating, &userReview.Review,
		 			&userReview.CreatedAt, &userReview.UpdatedAt)
		
			userReviews = append(userReviews, userReview)

			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": userReviews,
			"count":  len(userReviews),
		})
	})

	// POST new userReview details
	router.POST("/user-review", func(c *gin.Context) {
		orderId := c.PostForm("orderId")
		productId := c.PostForm("productId");
		userId := c.PostForm("userId");
		rating := c.PostForm("rating");
		review := c.PostForm("review");
		stmt, err := db.Prepare("INSERT INTO user_review (order_id, product_id, user_id, rating, review) VALUES(?,?,?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(orderId, productId, userId, rating, review)

		if err != nil {
			fmt.Print(err.Error())
		}

		defer stmt.Close()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("user-review from user_id: %s successfully created", userId),
		})
	})

	// PUT - update a userReview details
	// e.g. /user-review?id=1
	router.PUT("/user-review", func(c *gin.Context) {
		id := c.Query("id")
		rating := c.PostForm("rating");
		review := c.PostForm("review");
		stmt, err := db.Prepare("UPDATE user_review SET rating= ?, review= ? WHERE id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(rating, review, id)
		if err != nil {
			fmt.Print(err.Error())
		}

		defer stmt.Close()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("user-review with id %s successfully updated", id),
		})
	})

	// Delete resources
	// e.g. /user-review?id=1
	router.DELETE("/user-review", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("DELETE FROM user_review WHERE id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted user-review with id: %s", id),
		})
	})

	router.Run(":3000")
}