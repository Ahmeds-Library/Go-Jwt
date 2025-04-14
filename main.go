package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"
	"github.com/gin-gonic/gin"
	"os"
)

func main () {

	start := time.Now()

	r := gin.Default()

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
			"message": "File uploaded and analyzed successfully",
			"file":    file.Filename,
			"result":  result,
			"duration": time.Since(start).String(),
		})
	})

	r.Run(":8080")
}

func Analyze(a string) string {
	var words, lines, punctuations, spaces, vowels, digits, paragraphs, specialChars, consonants int

	for i := 0; i < len(a); i++ {
		if a[i] == ' ' {
			words++
			spaces++
		} else if a[i] == '\n' {
			words++
			lines++
			paragraphs++
		} else if a[i] == '\t' {
			words++
		} else if a[i] == '.' || a[i] == ',' || a[i] == '!' || a[i] == '?' || a[i] == ';' || a[i] == ':' || a[i] == ')' || a[i] == '(' || a[i] == '-' || a[i] == '_' || a[i] == '"' || a[i] == '\'' {
			punctuations++
		} else if a[i] == 'a' || a[i] == 'e' || a[i] == 'i' || a[i] == 'o' || a[i] == 'u' || a[i] == 'A' || a[i] == 'E' || a[i] == 'I' || a[i] == 'O' || a[i] == 'U' {
			vowels++
		} else if a[i] >= '0' && a[i] <= '9' {
			digits++
		} else if a[i] == '@' || a[i] == '#' || a[i] == '$' || a[i] == '%' || a[i] == '^' || a[i] == '&' || a[i] == '*' || a[i] == '+' || a[i] == '=' || a[i] == '{' || a[i] == '}' || a[i] == '[' || a[i] == ']' {
			specialChars++
		} else if (a[i] >= 'a' && a[i] <= 'z' || a[i] >= 'A' && a[i] <= 'Z') && !(a[i] == 'a' || a[i] == 'e' || a[i] == 'i' || a[i] == 'o' || a[i] == 'u' || a[i] == 'A' || a[i] == 'E' || a[i] == 'I' || a[i] == 'O' || a[i] == 'U') {
			consonants++
		}
	}

	result := fmt.Sprintf("Words: %d, Lines: %d, Punctuations: %d, Spaces: %d, Vowels: %d, Digits: %d, Paragraphs: %d, Special Characters: %d, Consonants: %d",
		words, lines+1, punctuations, spaces, vowels, digits, paragraphs+1, specialChars, consonants)
	return result
}


