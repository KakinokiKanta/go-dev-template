package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// データベース接続
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		"db", "db-user", "db-password", "db-name", "5432")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/get", func(ctx *gin.Context) {
		var (
			id int
			name string
			password string
		)
		row := db.QueryRow("SELECT * FROM accounts")
		err := row.Scan(&id, &name, &password)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err})
			return
		}
		if row.Err() != nil {
			ctx.JSON(200, gin.H{
				"id": id,
				"name": name,
				"password": password,
			})
		}
	})
	r.Run()
}
