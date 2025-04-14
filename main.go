package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func main() {

	start := time.Now()
	r := gin.Default()

	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	r.POST("/login", func(c *gin.Context) {
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
	})

	r.POST("/upload", func(c *gin.Context) {

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
			"message":  "File uploaded and analyzed successfully",
			"file":     file.Filename,
			"result":   result,
			"duration": time.Since(start).String(),
		})
	})

	r.Run(":8080")
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
