package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Ahmeds-Library/Go-Jwt/database"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if u.Username == "Ahmed" && u.Password == "Jwt" {
		tokenstring, err := createToken(u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create token"})
			return
		}

    var dbPassword string
	err = database.Db.QueryRow("SELECT password FROM users WHERE username = $1", u.Username).Scan(&dbPassword)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    if dbPassword != u.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
        return
    }

		c.JSON(http.StatusOK, gin.H{"token": tokenstring})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
}

func upload(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "File not found"})
		return
	}
	savePath := filepath.Join("uploads", file.Filename)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}

	content, err := os.ReadFile("./uploads/File.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := string(content)
	result := Analyze(data)

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded and analyzed successfully",
		"file":    file.Filename,
		"result":  result,
	})
}

func signup(c *gin.Context) {
    var u User

    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }	

	_, err := database.Db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", u.Username, u.Password)
    if err != nil {
		fmt.Println(err, "Line 86")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Username already exists or error in DB"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
