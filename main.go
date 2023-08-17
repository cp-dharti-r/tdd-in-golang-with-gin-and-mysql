package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	db, _ = sqlx.Open("mysql", "root:password@/db-name")
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
