package route_func

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Ahmeds-Library/Go-Jwt/analyze"
	"github.com/Ahmeds-Library/Go-Jwt/create_token"
	"github.com/Ahmeds-Library/Go-Jwt/database"
	db_results "github.com/Ahmeds-Library/Go-Jwt/save_results"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {
	var u User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := database.Db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", u.Username, u.Password)
	if err != nil {
		fmt.Println(err, "Line30")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Username already exists or error in DB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var dbPassword string
	err := database.Db.QueryRow("SELECT password FROM users WHERE username = $1", u.Username).Scan(&dbPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if u.Password == dbPassword {
		tokenstring, err := create_token.CreateToken(u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create token"})
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

func Upload(c *gin.Context) {

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
	result := analyze.Analyze(data)

	db_results.SaveResult(database.Db, result)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "File uploaded and analyzed successfully and also all Results are saved successfully in Database	",
		"file":    file.Filename,
		"result":  result,
	})
}
