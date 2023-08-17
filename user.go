package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// create user
func Create(c *gin.Context) {
	input := User{}
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", input.Name, input.Email)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// get user
func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var user User
	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

	err = row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

// update user
func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	input := User{}
	err = c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var user User
	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	err = row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	result, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", input.Name, input.Email, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// delete user
func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var user User
	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	err = row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	db.Exec("DELETE FROM users WHERE id = ?", id)

	c.JSON(http.StatusOK, gin.H{})
}
