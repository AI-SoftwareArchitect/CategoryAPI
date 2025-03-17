package middleware

import (
    "github.com/didip/tollbooth"
    "github.com/gin-gonic/gin"
)

func RateLimitMiddleware() gin.HandlerFunc {
    limiter := tollbooth.NewLimiter(100, nil)
    return func(c *gin.Context) {
        httpError := tollbooth.LimitByRequest(limiter, c.Writer, c.Request)
        if httpError != nil {
            c.AbortWithStatusJSON(httpError.StatusCode, gin.H{"error": httpError.Message})
            return
        }
        c.Next()
    }
}