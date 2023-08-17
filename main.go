package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:password@/db-name")
}

func main() {
	router := gin.Default()

	router.POST("/api/users", Create)
	router.GET("/api/users/:id", Get)
	router.PUT("/api/users/:id", Update)
	router.DELETE("/api/users/:id", Delete)

	router.Run(":8000")

	defer db.Close()
}
