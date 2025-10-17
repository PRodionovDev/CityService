package Middleware

import (
    "net/http"
	"log"
    "os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func AuthMiddleware() gin.HandlerFunc {
    requiredToken := os.Getenv("API_TOKEN")

    if requiredToken == "" {
      log.Fatal("Please set API_TOKEN environment variable")
    }

    return func(c *gin.Context) {
      if c.Request.Header["Authorization"] == nil {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API token required"})
        return
      }

      token := c.Request.Header["Authorization"][0]

      if token != requiredToken {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API token"})
        return
      }

      c.Next()
    }
}