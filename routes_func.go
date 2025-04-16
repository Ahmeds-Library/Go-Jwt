package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if u.Username == "Ahmed" || u.Password == "Jwt" {
		tokenstring, err := createToken(u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenstring})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
}
func login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if u.Username == "Ahmed" || u.Password == "Jwt" {
		tokenstring, err := createToken(u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create token"})
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
