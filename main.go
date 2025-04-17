package main

import (
	"github.com/Ahmeds-Library/Go-Jwt/database"
	"github.com/Ahmeds-Library/Go-Jwt/middleware"
	"github.com/Ahmeds-Library/Go-Jwt/route_func"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	r := gin.Default()
	r.POST("/signup", route_func.Signup)
	r.POST("/login", route_func.Login)
	r.POST("/upload", middleware.AuthMiddleware(), route_func.Upload)
	r.Run(":8000")
}
