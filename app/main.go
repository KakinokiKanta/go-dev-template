package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		"db-host", "db-user", "db-password", "db-name", "5432")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connection is successful")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
