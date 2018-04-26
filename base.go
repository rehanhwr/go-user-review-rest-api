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
		Rating 		float
		Review 		string
		CreatedAt 	time.Time
		UpdatedAt 	time.Time
	}
	router := gin.Default()
	router.Run(":3000")
}